package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

type InputField struct {
	inner       textinput.Model
	labelStyle  lipgloss.Style
	borderStyle lipgloss.Style
}

func NewInputField(accentColor string, fgColor string, mutedColor string) InputField {
	ti := textinput.New()
	ti.Placeholder = "romaji..."
	ti.CharLimit = 10
	ti.Width = 18
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(fgColor))
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(mutedColor))
	ti.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color(accentColor))
	ti.Focus()

	return InputField{
		inner:       ti,
		labelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color(mutedColor)).MarginBottom(1),
		borderStyle: lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color(accentColor)).Padding(0, 1).Width(22),
	}
}

func (f *InputField) Render() string {
	label := f.labelStyle.Render("Type Reading: ")
	field := f.borderStyle.Render(f.inner.View())
	return label + "\n" + field
}

func (f *InputField) Value() string {
	return f.inner.Value()
}

func (f *InputField) Reset() {
	f.inner.SetValue("")
}

func (f *InputField) Blur() {
	f.inner.Blur()
}

func (f *InputField) Update(msg interface{}) {
	type teaMsg = interface{}
	if m, ok := msg.(interface{}); ok {
		_ = m
	}

}

func (f *InputField) InnerModel() *textinput.Model {
	return &f.inner
}
