package components

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/nuixyz/kanarenshu/internal/theme"
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

func (cs CardStyle) Render(kana string, state CardState, hint string, palette theme.Palette) string {
	box := cs.box

	switch state {
	case CardCorrect:
		box = box.BorderForeground(lipgloss.Color(palette.Correct))
	case CardWrong:
		box = box.BorderForeground(lipgloss.Color(palette.Wrong))
	}

	char := cs.character.Render(kana)

	if hint != "" {
		hintLine := cs.hint.Render(hint)
		return box.Render(char + "\n\n" + hintLine)
	}

	return box.Render(char)
}

func (cs CardStyle) RenderKanji(char string, state CardState, onyomi, kunyomi []string, meaning string, palette theme.Palette) string {
	box := cs.box

	switch state {
	case CardCorrect:
		box = box.BorderForeground(lipgloss.Color(palette.Correct))
	case CardWrong:
		box = box.BorderForeground(lipgloss.Color(palette.Wrong))
	}

	body := cs.character.Render(char)

	if len(onyomi) > 0 || len(kunyomi) > 0 || meaning != "" {
		hint := ""
		if len(onyomi) > 0 {
			hint += "音: " + joinComma(onyomi)
		}
		if len(kunyomi) > 0 {
			if hint != "" {
				hint += "\n"
			}
			hint += "訓: " + joinComma(kunyomi)
		}
		if meaning != "" {
			if hint != "" {
				hint += "\n"
			}
			hint += meaning
		}
		body += "\n\n" + cs.hint.Render(hint)
	}
	return box.Render(body)
}

func joinComma(items []string) string {
	out := ""
	for i, s := range items {
		if i > 0 {
			out += ", "
		}
		out += s
	}
	return out
}
