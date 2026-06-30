package screens

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nuixyz/kanarenshu/internal/data"
)

type OpenJLPTSelectMsg struct{}

type StartKanjiStudyMsg struct {
	JLPT string
}

type JLPTSelectModel struct {
	cursor int
	levels []string

	titleStyle     lipgloss.Style
	subtitleStyle  lipgloss.Style
	selectedStyle  lipgloss.Style
	normalStyle    lipgloss.Style
	footerStyle    lipgloss.Style
	containerStyle lipgloss.Style
}

func NewJLPTSelectModel(bgColor, fgColor, accentColor, mutedColor, selectedBg string) JLPTSelectModel {
	return JLPTSelectModel{
		cursor: 0,
		levels: data.AvailableJLPT(),

		titleStyle:     lipgloss.NewStyle().Foreground(lipgloss.Color(accentColor)).Bold(true).MarginBottom(1),
		subtitleStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color(mutedColor)).MarginBottom(2),
		selectedStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color(bgColor)).Background(lipgloss.Color(accentColor)).Bold(true).Padding(0, 2),
		normalStyle:    lipgloss.NewStyle().Foreground(lipgloss.Color(fgColor)).Padding(0, 2),
		footerStyle:    lipgloss.NewStyle().Foreground(lipgloss.Color(mutedColor)).MarginTop(2),
		containerStyle: lipgloss.NewStyle().Padding(2, 4),
	}
}

func (m JLPTSelectModel) Init() tea.Cmd {
	return nil
}

func (m JLPTSelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.levels)-1 {
				m.cursor++
			}
		case "enter", " ":
			if len(m.levels) == 0 {
				return m, nil
			}
			jlpt := m.levels[m.cursor]
			return m, func() tea.Msg { return StartKanjiStudyMsg{JLPT: jlpt} }
		case "esc":
			return m, func() tea.Msg { return BackToMenuMsg{} }
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m JLPTSelectModel) View() string {
	title := m.titleStyle.Render("Kanji - Select JLPT Level")

	if len(title) == 0 {
		body := title + "\n" + m.subtitleStyle.Render("No kanji data available yet.")
		return m.containerStyle.Render(body)
	}

	list := ""
	for i, lvl := range m.levels {
		if i == m.cursor {
			list += m.selectedStyle.Render(" "+lvl) + "\n"
		} else {
			list += m.normalStyle.Render(" "+lvl) + "\n"
		}
	}

	footer := m.footerStyle.Render("↑↓ / jk to move		enter to select		ctrl+t cycle theme		q to quit")

	body := fmt.Sprintf("%s\n%s\n%s", title, list, footer)
	return m.containerStyle.Render(body)
}
