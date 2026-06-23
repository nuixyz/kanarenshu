package ui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/nuixyz/kanarenshu/internal/data"
	"github.com/nuixyz/kanarenshu/internal/game"
	"github.com/nuixyz/kanarenshu/internal/logger"
	"github.com/nuixyz/kanarenshu/internal/ui/screens"
)

type Palette struct {
	Bg      string
	Fg      string
	Accent  string
	Correct string
	Wrong   string
	Muted   string
	Border  string
	SelBg   string
}

func DefaultPalette() Palette {
	return Palette{
		Bg:      "#1a1b26",
		Fg:      "#c0caf5",
		Accent:  "#7aa2f7",
		Correct: "#9ece6a",
		Wrong:   "#f7768e",
		Muted:   "#565f89",
		Border:  "#3b4261",
		SelBg:   "#7aa2f7",
	}
}

type Renderer struct {
	current tea.Model
	palette Palette
	lives   int
}

func NewRenderer(palette Palette, lives int) Renderer {
	p := palette
	return Renderer{
		palette: p,
		lives:   lives,
		current: screens.NewMenuModel(
			p.Bg, p.Fg, p.Accent, p.Muted, p.SelBg,
		),
	}
}

func (r Renderer) Init() tea.Cmd {
	return r.current.Init()
}

func (r Renderer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case screens.StartStudyMsg:
		logger.Info("transitioning to study screen, mode=%d", msg.Mode)
		cfg := game.Config{
			Mode:       data.Mode(msg.Mode),
			Lives:      r.lives,
			StartLevel: 0,
		}
		p := r.palette
		r.current = screens.NewStudyModel(
			cfg,
			p.Bg, p.Fg, p.Accent, p.Muted, p.Correct, p.Wrong, p.Border,
		)
		return r, r.current.Init()

	case screens.SessionEndMsg:
		logger.Info("transitioning to results screen, score=%d", msg.Summary.Score)
		p := r.palette
		r.current = screens.NewResultsModel(
			msg.Summary,
			p.Bg, p.Fg, p.Accent, p.Muted, p.Correct, p.Wrong,
		)
		return r, r.current.Init()

	case screens.BackToMenuMsg:
		logger.Info("transitioning back to menu")
		p := r.palette
		r.current = screens.NewMenuModel(
			p.Bg, p.Fg, p.Accent, p.Muted, p.SelBg,
		)
		return r, r.current.Init()

	case screens.PlayAgainMsg:
		logger.Info("play again, mode=%d", msg.Mode)
		cfg := game.Config{
			Mode:       data.Mode(msg.Mode),
			Lives:      r.lives,
			StartLevel: 0,
		}
		p := r.palette
		r.current = screens.NewStudyModel(
			cfg,
			p.Bg, p.Fg, p.Accent, p.Muted, p.Correct, p.Wrong, p.Border,
		)
		return r, r.current.Init()
	}

	var cmd tea.Cmd
	r.current, cmd = r.current.Update(msg)
	return r, cmd
}

func (r Renderer) View() string {
	return r.current.View()
}
