package ui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/nuixyz/kanarenshu/internal/data"
	"github.com/nuixyz/kanarenshu/internal/game"
	"github.com/nuixyz/kanarenshu/internal/logger"
	"github.com/nuixyz/kanarenshu/internal/storage"
	"github.com/nuixyz/kanarenshu/internal/theme"
	"github.com/nuixyz/kanarenshu/internal/ui/screens"
)

type Renderer struct {
	current tea.Model
	palette theme.Palette
	lives   int
	cfg     storage.Config

	themeList []string
	themeIdx  int
}

func NewRenderer(palette theme.Palette, cfg storage.Config) Renderer {
	themeList := sortedThemeList()
	themeIdx := indexOfTheme(themeList, cfg.Theme)

	p := palette
	r := Renderer{
		palette:   p,
		lives:     cfg.Lives,
		cfg:       cfg,
		themeList: themeList,
		themeIdx:  themeIdx,
	}
	r.current = r.newMenu()
	return r
}

func (r Renderer) Init() tea.Cmd {
	return r.current.Init()
}

func (r Renderer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// global shortcut
	if key, ok := msg.(tea.KeyMsg); ok && key.String() == "ctrl+t" {
		r.themeIdx = (r.themeIdx + 1) % len(r.themeList)
		name := r.themeList[r.themeIdx]
		r.applyThemeByName(name)
		r.cfg.Theme = name
		// Rebuild the current screen with the new palette colours.
		r.current = r.rebuildCurrent()
		logger.Info("Theme cycled to %s", name)
		return r, r.current.Init()
	}

	// screen scoped
	switch msg := msg.(type) {

	case screens.StartStudyMsg:
		logger.Info("Transitioning to Study Screen, mode=%d", msg.Mode)
		mode := data.Mode(msg.Mode)

		highest, err := storage.HighestLevelFor(mode)
		if err != nil {
			logger.Error("Failed to load progress: %v", err)
			highest = 0
		}

		cfg := game.Config{
			Mode:       mode,
			Lives:      r.lives,
			StartLevel: highest,
		}

		r.current = r.newStudy(cfg)
		return r, r.current.Init()

	case screens.StartKanjiStudyMsg:
		logger.Info("Transitioning to Kanji Study Screen, JLPT=%s", msg.JLPT)

		highest, err := storage.HighestKanjiLevelFor(msg.JLPT)
		if err != nil {
			logger.Error("Failed to load Kanji progress: %v", err)
			highest = 0
		}

		cfg := game.KanjiConfig{
			JLPT:       msg.JLPT,
			Lives:      r.lives,
			StartLevel: highest,
		}

		r.current = r.newKanjiStudy(cfg)
		return r, r.current.Init()

	case screens.SessionEndMsg:
		logger.Info("Transitioning to Results Screen, score=%d", msg.Summary.Score)
		r.current = r.newResults(msg.Summary)
		return r, r.current.Init()

	case screens.KanjiSessionEndMsg:
		logger.Info("Transitioning to Kanji Results Screen, score=%d", msg.Summary.Score)
		r.current = r.newResults(msg.Summary)
		return r, r.current.Init()

	case screens.BackToMenuMsg:
		logger.Info("Transitioning back to Menu")
		r.current = r.newMenu()
		return r, r.current.Init()

	case screens.PlayAgainMsg:
		logger.Info("Play Again, mode=%d", msg.Mode)

		if msg.Mode == 3 { // Kanji
			highest, err := storage.HighestKanjiLevelFor(msg.JLPT)
			if err != nil {
				logger.Error("Failed to load Kanji progress: %v", err)
				highest = 0
			}
			cfg := game.KanjiConfig{
				JLPT:       msg.JLPT,
				Lives:      r.lives,
				StartLevel: highest,
			}
			r.current = r.newKanjiStudy(cfg)
			return r, r.current.Init()
		}

		mode := data.Mode(msg.Mode)
		highest, err := storage.HighestLevelFor(mode)
		if err != nil {
			logger.Error("Failed to load progress: %v", err)
			highest = 0
		}
		cfg := game.Config{
			Mode:       mode,
			Lives:      r.lives,
			StartLevel: highest,
		}
		r.current = r.newStudy(cfg)
		return r, r.current.Init()

	case screens.OpenSettingsMsg:
		logger.Info("Transitionning to Settings Screen")
		r.current = r.newSettings()
		return r, r.current.Init()

	case screens.OpenJLPTSelectMsg:
		logger.Info("Transitioning to JLPT Select Screen")
		r.current = r.newKanjiMenu()
		return r, r.current.Init()

	case screens.ApplyThemeMsg:
		r.applyThemeByName(msg.ThemeName)
		r.cfg.Theme = msg.ThemeName
		r.themeIdx = indexOfTheme(r.themeList, msg.ThemeName)
		r.current = r.newSettings()
		return r, r.current.Init()

	case screens.SaveSettingsMsg:
		logger.Info("Settings saved, theme=%s", msg.Config.Theme)
		r.cfg = msg.Config
		r.lives = msg.Config.Lives
		r.applyThemeByName(msg.Config.Theme)
		r.themeIdx = indexOfTheme(r.themeList, msg.Config.Theme)
		r.current = r.newMenu()
		return r, r.current.Init()
	}

	// delegate to active screen
	var cmd tea.Cmd
	r.current, cmd = r.current.Update(msg)
	return r, cmd
}

func (r Renderer) View() string {
	return r.current.View()
}

// constructors
func (r *Renderer) newMenu() tea.Model {
	p := r.palette
	return screens.NewMenuModel(p.Bg, p.Fg, p.Accent, p.Muted, p.SelBg)
}

func (r *Renderer) newStudy(cfg game.Config) tea.Model {
	p := r.palette
	return screens.NewStudyModel(cfg, p.Bg, p.Fg, p.Accent, p.Muted, p.Correct, p.Wrong, p.Border)
}

func (r *Renderer) newKanjiStudy(cfg game.KanjiConfig) tea.Model {
	p := r.palette
	return screens.NewKanjiStudyModel(cfg, p.Bg, p.Fg, p.Accent, p.Muted, p.Correct, p.Wrong, p.Border)
}

func (r *Renderer) newResults(sum game.Summary) tea.Model {
	p := r.palette
	return screens.NewResultsModel(sum, p.Bg, p.Fg, p.Accent, p.Muted, p.Correct, p.Wrong)
}

func (r *Renderer) newKanjiMenu() tea.Model {
	p := r.palette
	return screens.NewJLPTSelectModel(p.Bg, p.Fg, p.Accent, p.Muted, p.SelBg)
}

func (r *Renderer) newSettings() tea.Model {
	p := r.palette
	return screens.NewSettingsModel(r.cfg, p.Bg, p.Fg, p.Accent, p.Muted, p.SelBg)
}

func (r *Renderer) rebuildCurrent() tea.Model {
	switch r.current.(type) {
	case screens.SettingsModel:
		return r.newSettings()
	case screens.MenuModel:
		return r.newMenu()
	case screens.JLPTSelectModel:
		return r.newKanjiMenu()
	case screens.KanjiStudyModel:
		return r.newMenu()
	default:
		return r.newMenu()
	}
}

func (r *Renderer) applyThemeByName(name string) {
	p, err := theme.Load(name)
	if err != nil {
		logger.Error("Failed to load theme %q: %v", name, err)
		return
	}
	r.palette = p
}

func sortedThemeList() []string {
	all := theme.Available()
	for i := 1; i < len(all); i++ {
		for j := i; j > 0 && all[j] < all[j-1]; j-- {
			all[j], all[j-1] = all[j-1], all[j]
		}
	}
	return all
}

func indexOfTheme(list []string, name string) int {
	for i, s := range list {
		if s == name {
			return i
		}
	}
	return 0
}
