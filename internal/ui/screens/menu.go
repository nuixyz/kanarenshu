package screens

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type StartStudyMsg struct {
	Mode int // 0, 1, 2
}

type QuitMsg struct{}

type menuItem struct {
	label    string
	sublabel string
	msg      tea.Msg
}

type MenuModel struct {
	cursor int
	items  []menuItem

	titleStyle     lipgloss.Style
	subtitleStyle  lipgloss.Style
	selectedStyle  lipgloss.Style
	normalStyle    lipgloss.Style
	sublabelStyle  lipgloss.Style
	footerStyle    lipgloss.Style
	containerStyle lipgloss.Style
}

func NewMenuModel(
	bgColor, fgColor, accentColor, mutedColor, selectedBg string,
) MenuModel {
	items := []menuItem{
		{
			label:    "Hiragana",
			sublabel: "ひらがな",
			msg:      StartStudyMsg{Mode: 0},
		},
		{
			label:    "Katakana",
			sublabel: "カタカナ",
			msg:      StartStudyMsg{Mode: 1},
		},
		{
			label:    "Mixed",
			sublabel: "Both scripts interleaved",
			msg:      StartStudyMsg{Mode: 2},
		},
		{
			label:    "Quit",
			sublabel: "Have a great day!",
			msg:      QuitMsg{},
		},
	}

	return MenuModel{
		cursor: 0,
		items:  items,

		titleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color(accentColor)).Bold(true).MarginBottom(0),

		subtitleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color(mutedColor)).MarginBottom(2),

		selectedStyle: lipgloss.NewStyle().Foreground(lipgloss.Color(bgColor)).Background(lipgloss.Color(accentColor)).Bold(true).Padding(0, 2),

		normalStyle:    lipgloss.NewStyle().Foreground(lipgloss.Color(fgColor)).Padding(0, 2),
		sublabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color(mutedColor)).Padding(0, 3),
		footerStyle:    lipgloss.NewStyle().Foreground(lipgloss.Color(mutedColor)).MarginTop(2),
		containerStyle: lipgloss.NewStyle().Padding(2, 4),
	}
}

func (m MenuModel) Init() tea.Cmd {
	return nil
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		case "enter", " ":
			selected := m.items[m.cursor]
			if _, ok := selected.msg.(QuitMsg); ok {
				return m, tea.Quit
			}
			return m, func() tea.Msg { return selected.msg }
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m MenuModel) View() string {
	title := m.titleStyle.Render("かなれんしゅ") + "\n" + m.titleStyle.Render("kanarenshu")
	subtitle := m.sublabelStyle.Render("Japanese kana practise for the terminal")

	menu := ""
	for i, item := range m.items {
		var row string
		if i == m.cursor {
			row = m.selectedStyle.Render(" " + item.label)
		} else {
			row = m.normalStyle.Render(" " + item.label)
		}

		menu += row + "\n"
		menu += m.sublabelStyle.Render(item.sublabel) + "\n"

		if i < len(m.items)-1 {
			menu += "\n"
		}
	}

	footer := m.footerStyle.Render(" / jk to move		enter to select		q to quit")

	body := fmt.Sprintf("%s\n%s\n%s\n%s", title, subtitle, menu, footer)

	return m.containerStyle.Render(body)
}
