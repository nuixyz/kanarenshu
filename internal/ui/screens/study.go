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
	"github.com/nuixyz/kanarenshu/internal/storage"
	"github.com/nuixyz/kanarenshu/internal/ui/components"
)

type SessionEndMsg struct {
	Summary game.Summary
}

type BackToMenuMsg struct{}

type flashTimeoutMsg struct{}

type levelUpTimeoutMsg struct{}

type StudyModel struct {
	session   *game.Session
	input     textinput.Model
	cardState components.CardState
	card      components.CardStyle
	stats     components.StatsPanel
	progress  components.ProgressBar

	showLevelUp bool
	newChars    []string
	showHint    bool

	hintText string

	width  int
	height int

	levelUpStyle   lipgloss.Style
	newCharStyle   lipgloss.Style
	hintStyle      lipgloss.Style
	footerStyle    lipgloss.Style
	containerStyle lipgloss.Style
	modeStyle      lipgloss.Style
}

// Pass the fully loaded pointer from your main application orchestrator
func NewStudyModel(
	cfg game.Config, store *storage.ProgressStore, bgColor, fgColor, accentColor, mutedColor, correctColor, wrongColor, borderColor string,
) StudyModel {
	ti := textinput.New()
	ti.Placeholder = "romaji..."
	ti.CharLimit = 10
	ti.Width = 18
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(fgColor))
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(mutedColor))
	ti.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color(accentColor))
	ti.Focus()

	return StudyModel{
		session:   game.NewSession(cfg, store), // Injects persistence down to game loop
		input:     ti,
		cardState: components.CardNeutral,

		card:      components.NewCardStyle(borderColor, fgColor, correctColor, wrongColor, mutedColor),
		stats:     components.NewStatsPanel(wrongColor, mutedColor, accentColor, correctColor, mutedColor),
		progress:  components.NewProgressBar(20, accentColor, mutedColor, mutedColor),

		levelUpStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color(bgColor)).Background(lipgloss.Color(accentColor)).Bold(true).Padding(0, 2).MarginBottom(1),
		newCharStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color(accentColor)),
		hintStyle:      lipgloss.NewStyle().Foreground(lipgloss.Color(wrongColor)).Italic(true),
		footerStyle:    lipgloss.NewStyle().Foreground(lipgloss.Color(mutedColor)).MarginTop(1),
		containerStyle: lipgloss.NewStyle().Padding(1, 4),
		modeStyle:      lipgloss.NewStyle().Foreground(lipgloss.Color(mutedColor)),
	}
}

func (m StudyModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m StudyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case flashTimeoutMsg:
		m.cardState = components.CardNeutral
		m.showHint = false
		m.hintText = ""

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
			logger.Debug("submit answer=%q result=%d", answer, result)

			switch result {
			case game.AnswerCorrect:
				m.cardState = components.CardCorrect
				cmds = append(cmds, flashAfter(300*time.Millisecond))

			case game.AnswerLevelUp:
				m.cardState = components.CardCorrect
				m.showLevelUp = true
				m.newChars = game.NewCharsForLevel(
					m.session.Cfg().Mode,
					m.session.Level,
				)
				cmds = append(cmds, flashAfter(300*time.Millisecond))
				cmds = append(cmds, levelUpAfter(2*time.Second))

			case game.AnswerWrong:
				m.cardState = components.CardWrong
				m.showHint = true
				m.hintText = m.session.Current().Primary
				cmds = append(cmds, flashAfter(600*time.Millisecond))

			case game.AnswerGameOver:
				sum := game.Summarise(m.session)
				return m, func() tea.Msg { return SessionEndMsg{Summary: sum} }
			}
		}
	}

	var tiCmd tea.Cmd
	m.input, tiCmd = m.input.Update(msg)
	cmds = append(cmds, tiCmd)

	return m, tea.Batch(cmds...)
}

func (m StudyModel) View() string {
	s := m.session

	levelLabel := game.LevelLabel(s.Level)
	progress := game.LevelProgress(s)
	progressBar := m.progress.Render(progress, levelLabel)
	modeLabel := m.modeStyle.Render(s.Cfg().Mode.String())

	topBar := lipgloss.JoinHorizontal(lipgloss.Top, progressBar, strings.Repeat(" ", 4), modeLabel)

	statsLien := m.stats.Render(
		s.Lives,
		m.session.Cfg().Lives,
		s.Streak,
		s.Score,
	)

	hint := ""
	if m.showHint {
		hint = m.hintText
	}
	cardView := m.card.Render(s.Current().Kana, m.cardState, hint)

	inputLabel := lipgloss.NewStyle().Foreground(lipgloss.Color("565f89")).Render("Type Reading: ")

	inputBox := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#7aa2f7")).Padding(0, 1).Width(22).Render(m.input.View())

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

	footer := m.footerStyle.Render("enter to submit        esc to menu        ctrl+c to quit")

	body := fmt.Sprintf(
		"%s\n\n%s\n\n%s\n\n%s\n\n%s%s",
		topBar,
		statsLien,
		cardView,
		inputArea,
		levelUpBanner,
		footer,
	)
	return m.containerStyle.Render(body)
}

func flashAfter(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(_ time.Time) tea.Msg {
		return flashTimeoutMsg{}
	})
}

func levelUpAfter(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(_ time.Time) tea.Msg {
		return levelUpTimeoutMsg{}
	})
}