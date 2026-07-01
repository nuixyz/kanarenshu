package screens

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/nuixyz/kanarenshu/internal/game"
)

type PlayAgainMsg struct {
	Mode int
	JLPT string
}

type ResultsModel struct {
	summary game.Summary

	titleStyle     lipgloss.Style
	gradeStyle     lipgloss.Style
	statLabelStyle lipgloss.Style
	statValueStyle lipgloss.Style
	accentStyle    lipgloss.Style
	footerStyle    lipgloss.Style
	containerStyle lipgloss.Style
	dividerStyle   lipgloss.Style
}

func NewResultsModel(
	summary game.Summary,
	bgColor, fgColor, accentColor, mutedColor, correctColor, wrongColor string,
) ResultsModel {
	gradeColor := gradeToColor(summary.Grade, correctColor, accentColor, mutedColor, wrongColor)

	return ResultsModel{
		summary: summary,

		titleStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color(accentColor)).
			Bold(true).
			MarginBottom(1),

		gradeStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color(bgColor)).
			Background(lipgloss.Color(gradeColor)).
			Bold(true).
			Width(5).
			Align(lipgloss.Center).
			Padding(0, 1),

		statLabelStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color(mutedColor)).
			Width(14),

		statValueStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color(fgColor)).
			Bold(true),

		accentStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color(accentColor)),

		footerStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color(mutedColor)).
			MarginTop(2),

		containerStyle: lipgloss.NewStyle().
			Padding(2, 4),

		dividerStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color(mutedColor)),
	}
}

func (m ResultsModel) Init() tea.Cmd {
	return nil
}

func (m ResultsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "r", "R":
			mode := modeStringToInt(m.summary.Mode)
			return m, func() tea.Msg { return PlayAgainMsg{Mode: mode, JLPT: m.summary.JLPT} }
		case "m", "M":
			return m, func() tea.Msg { return BackToMenuMsg{} }
		case "q", "Q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m ResultsModel) View() string {
	sum := m.summary

	title := m.titleStyle.Render("Session Complete")
	divider := m.dividerStyle.Render("────────────────────────")

	grade := m.gradeStyle.Render(string(sum.Grade))

	statLine := func(label, value string) string {
		return m.statLabelStyle.Render(label) + m.statValueStyle.Render(value)
	}

	stats := fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		statLine("Score:   ", fmt.Sprintf("%d / %d", sum.Score, sum.Total)),
		statLine("Accuracy:", sum.AccuracyPercent()),
		statLine("Level:   ", fmt.Sprintf("%d", sum.MaxLevel+1)),
		statLine("Mode:    ", sum.Mode),
	)

	gradeBlock := fmt.Sprintf(
		"\n%s\n%s\n\n%s\n",
		m.accentStyle.Render("Grade"),
		grade,
		stats,
	)

	footer := m.footerStyle.Render(
		"r  play again		m  main menu		q  quit",
	)

	body := fmt.Sprintf("%s\n%s%s\n%s", title, gradeBlock, divider, footer)
	return m.containerStyle.Render(body)
}

func gradeToColor(g game.Grade, correct, accent, muted, wrong string) string {
	switch g {
	case game.GradeS:
		return accent
	case game.GradeA:
		return correct
	case game.GradeB:
		return accent
	case game.GradeC:
		return muted
	default:
		return wrong
	}
}

func modeStringToInt(mode string) int {
	switch {
	case strings.HasPrefix(mode, "Katakana"):
		return 1
	case strings.HasPrefix(mode, "Mixed"):
		return 2
	case strings.HasPrefix(mode, "Kanji"):
		return 3
	default:
		return 0
	}
}
