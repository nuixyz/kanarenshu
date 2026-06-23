package components

import (
	"github.com/charmbracelet/lipgloss"
)

type CardStyle struct {
	box       lipgloss.Style
	character lipgloss.Style
	hint      lipgloss.Style
}

type CardState int

const (
	CardNeutral CardState = iota
	CardCorrect
	CardWrong
)

func NewCardStyle(
	borderColor string,
	charColor string,
	correctColor string,
	wrongColor string,
	hintColor string,
) CardStyle {
	base := lipgloss.NewStyle().Width(20).Height(7).Align(lipgloss.Center, lipgloss.Center).Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color(borderColor))

	return CardStyle{
		box:       base,
		character: lipgloss.NewStyle().Foreground(lipgloss.Color(charColor)).Bold(true),
		hint:      lipgloss.NewStyle().Foreground(lipgloss.Color(hintColor)).Italic(true),
	}
}

func (cs CardStyle) Render(kana string, state CardState, hint string) string {
	box := cs.box

	switch state {
	case CardCorrect:
		box = box.BorderForeground(lipgloss.Color("#9ece6a"))
	case CardWrong:
		box = box.BorderForeground(lipgloss.Color("#f7768e"))
	}

	char := cs.character.Render(kana)

	if hint != "" {
		hintLine := cs.hint.Render(hint)
		return box.Render(char + "\n\n" + hintLine)
	}

	return box.Render(char)
}
