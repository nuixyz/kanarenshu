package screens

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nuixyz/kanarenshu/internal/game"
	"github.com/nuixyz/kanarenshu/internal/logger"
	"github.com/nuixyz/kanarenshu/internal/theme"
	"github.com/nuixyz/kanarenshu/internal/ui/components"
)

type KanjiSessionEndMsg struct {
	Summary game.Summary
}

type KanjiStudyModel struct {
	session   *game.KanjiSession
	input     textinput.Model
	cardState components.CardState
	card      components.CardStyle
	stats     components.StatsPanel
	progress  components.ProgressBar

	palette theme.Palette

	showLevelUp bool
	newChars    []string
	showHint    bool

	width  int
	height int

	levelUpStyle   lipgloss.Style
	newCharStyle   lipgloss.Style
	footerStyle    lipgloss.Style
	containerStyle lipgloss.Style
	modeStyle      lipgloss.Style
}

func NewKanjiStudyModel(cfg game.KanjiConfig, bgColor, fgColor, accentColor, mutedColor, correctColor, wrongColor, borderColor string) KanjiStudyModel {
	ti := textinput.New()
	ti.Placeholder = "Type reading or meaning"
	ti.CharLimit = 24
	ti.Width = 24
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(fgColor))
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(mutedColor))
	ti.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color(accentColor))
	ti.Focus()

	return KanjiStudyModel{
		session:   game.NewKanjiSession(cfg),
		input:     ti,
		cardState: components.CardNeutral,

		palette: theme.Palette{
			Bg:      bgColor,
			Fg:      fgColor,
			Accent:  accentColor,
			Muted:   mutedColor,
			Correct: correctColor,
			Wrong:   wrongColor,
			Border:  borderColor,
		},

		card:     components.NewCardStyle(borderColor, fgColor, correctColor, wrongColor, mutedColor),
		stats:    components.NewStatsPanel(wrongColor, mutedColor, accentColor, correctColor, mutedColor),
		progress: components.NewProgressBar(20, accentColor, mutedColor, mutedColor),

		levelUpStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color(bgColor)).Background(lipgloss.Color(accentColor)).Bold(true).Padding(0, 2).MarginBottom(1),
		newCharStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color(accentColor)),
		footerStyle:    lipgloss.NewStyle().Foreground(lipgloss.Color(mutedColor)).MarginTop(1),
		containerStyle: lipgloss.NewStyle().Padding(1, 4),
		modeStyle:      lipgloss.NewStyle().Foreground(lipgloss.Color(mutedColor)),
	}
}

func (m KanjiStudyModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m KanjiStudyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case flashTimeoutMsg:
		m.cardState = components.CardNeutral
		m.showHint = false

	case levelUpTimeoutMsg:
		m.showLevelUp = false
		m.newChars = nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "esc":
			return m, func() tea.Msg { return BackToMenuMsg{} }

		case "enter":
			answer := strings.TrimSpace(m.input.Value())
			if answer == "" {
				break
			}
			m.input.SetValue("")
			cmds = append(cmds, textinput.Blink)

			result := m.session.Submit(answer)
			logger.Debug("Kanji submitted=%q result=%d", answer, result)

			switch result {
			case game.AnswerCorrect:
				m.cardState = components.CardCorrect
				cmds = append(cmds, flashAfter(300*time.Millisecond))

			case game.AnswerLevelUp:
				m.cardState = components.CardCorrect
				m.showLevelUp = true
				m.newChars = game.NewKanjiForLevel(
					m.session.Cfg().JLPT,
					m.session.Level,
				)
				cmds = append(cmds, flashAfter(300*time.Millisecond))
				cmds = append(cmds, levelUpAfter(2*time.Second))

			case game.AnswerWrong:
				m.cardState = components.CardWrong
				m.showHint = true
				cmds = append(cmds, flashAfter(600*time.Millisecond))

			case game.AnswerGameOver:
				sum := game.SummariseKanji(m.session)
				return m, func() tea.Msg { return KanjiSessionEndMsg{Summary: sum} }
			}
		}
	}

	var tiCmd tea.Cmd
	m.input, tiCmd = m.input.Update(msg)
	cmds = append(cmds, tiCmd)

	return m, tea.Batch(cmds...)
}

func (m KanjiStudyModel) View() string {
	s := m.session
	cur := s.Current()

	levelLabel := game.KanjiLevelLabel(s.Level)
	progress := game.KanjiLevelProgress(s)
	progressBar := m.progress.Render(progress, levelLabel)
	modeLabel := m.modeStyle.Render("Kanji " + s.Cfg().JLPT)

	topBar := lipgloss.JoinHorizontal(lipgloss.Top, progressBar, strings.Repeat(" ", 4), modeLabel)

	statsLine := m.stats.Render(
		s.Lives,
		m.session.Cfg().Lives,
		s.Streak,
		s.Score,
	)

	var cardView string
	if m.showHint {
		meaning := ""
		if len(cur.Meanings) > 0 {
			meaning = cur.Meanings[0]
		}
		cardView = m.card.RenderKanji(cur.Char, m.cardState, cur.Onyomi, cur.Kunyomi, meaning, m.palette)
	} else {
		cardView = m.card.RenderKanji(cur.Char, m.cardState, nil, nil, "", m.palette)
	}

	inputLabel := lipgloss.NewStyle().Foreground(lipgloss.Color(m.palette.Muted)).Render("Reading: ")

	inputBox := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color(m.palette.Accent)).Padding(0, 1).Width(26).Render(m.input.View())

	inputArea := inputLabel + "\n" + inputBox

	levelUpBanner := ""
	if m.showLevelUp {
		banner := m.levelUpStyle.Render(
			fmt.Sprintf("Level %d unlocked!", s.Level+1),
		)
		newCharLine := ""
		if len(m.newChars) > 0 {
			newCharLine = "\n" + m.newCharStyle.Render(
				"New: "+strings.Join(m.newChars, " "),
			)
		}
		levelUpBanner = banner + newCharLine + "\n"
	}

	footer := m.footerStyle.Render("enter to submit		esc to menu		ctrl+c to quit")

	body := fmt.Sprintf(
		"%s\n\n%s\n\n%s\n\n%s\n\n%s%s",
		topBar,
		statsLine,
		cardView,
		inputArea,
		levelUpBanner,
		footer,
	)
	return m.containerStyle.Render(body)
}
