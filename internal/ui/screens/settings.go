package screens

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nuixyz/kanarenshu/internal/storage"
	"github.com/nuixyz/kanarenshu/internal/theme"
)

// upon theme change
type ApplyThemeMsg struct {
	ThemeName string
}

// when the user saves and quits
type SaveSettingsMsg struct {
	Config storage.Config
}

type settingKind int

const (
	kindTheme settingKind = iota
	kindMode
	kindLives
	kindHints
)

type settingRow struct {
	label   string
	kind    settingKind
	idx     int
	options []string
	boolVal bool
	intVal  int
	intMin  int
	intMax  int
}

func (r settingRow) displayValue() string {
	switch r.kind {
	case kindTheme, kindMode:
		return r.options[r.idx]
	case kindHints:
		if r.boolVal {
			return "on"
		}
		return "off"
	case kindLives:
		if r.intVal == 0 {
			return "∞"
		}
		return fmt.Sprintf("%d", r.intVal)
	}
	return ""
}

// model
type SettingsModel struct {
	rows   []settingRow
	cursor int

	original storage.Config

	titleStyle       lipgloss.Style
	selectedStyle    lipgloss.Style
	normalStyle      lipgloss.Style
	labelStyle       lipgloss.Style
	valueStyle       lipgloss.Style
	selectedValStyle lipgloss.Style
	mutedStyle       lipgloss.Style
	footerStyle      lipgloss.Style
	containerStyle   lipgloss.Style
	dividerStyle     lipgloss.Style
}

func NewSettingsModel(
	cfg storage.Config,
	bgColor, fgColor, accentColor, mutedColor, selectedBg string,
) SettingsModel {
	themeOptions := sortedThemes()
	themeIdx := indexOfOrZero(themeOptions, cfg.Theme)

	modeOptions := []string{"hiragana", "katakana", "mixed"}
	modeIdx := indexOfOrZero(modeOptions, cfg.Mode)

	rows := []settingRow{
		{
			label:   "Theme",
			kind:    kindTheme,
			idx:     themeIdx,
			options: themeOptions,
		},
		{
			label:   "Mode",
			kind:    kindMode,
			idx:     modeIdx,
			options: modeOptions,
		},
		{
			label:  "Lives",
			kind:   kindLives,
			intVal: cfg.Lives,
			intMin: 0,
			intMax: 9,
		},
		{
			label:   "Show hints",
			kind:    kindHints,
			boolVal: cfg.ShowHints,
		},
	}

	return SettingsModel{
		rows:     rows,
		cursor:   0,
		original: cfg,

		titleStyle:       lipgloss.NewStyle().Foreground(lipgloss.Color(accentColor)).Bold(true).MarginBottom(1),
		selectedStyle:    lipgloss.NewStyle().Foreground(lipgloss.Color(bgColor)).Background(lipgloss.Color(accentColor)).Bold(true).Padding(0, 1),
		normalStyle:      lipgloss.NewStyle().Foreground(lipgloss.Color(fgColor)).Padding(0, 1),
		labelStyle:       lipgloss.NewStyle().Foreground(lipgloss.Color(fgColor)).Width(16),
		valueStyle:       lipgloss.NewStyle().Foreground(lipgloss.Color(mutedColor)),
		selectedValStyle: lipgloss.NewStyle().Foreground(lipgloss.Color(accentColor)).Bold(true),
		mutedStyle:       lipgloss.NewStyle().Foreground(lipgloss.Color(mutedColor)),
		footerStyle:      lipgloss.NewStyle().Foreground(lipgloss.Color(mutedColor)).MarginTop(2),
		containerStyle:   lipgloss.NewStyle().Padding(2, 4),
		dividerStyle:     lipgloss.NewStyle().Foreground(lipgloss.Color(mutedColor)),
	}
}

func (m SettingsModel) Init() tea.Cmd {
	return nil
}

func (m SettingsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c":
			return m, tea.Quit

		case "esc", "q":
			// Discard changes — go back to menu.
			return m, func() tea.Msg { return BackToMenuMsg{} }

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.rows)-1 {
				m.cursor++
			}

		case "left", "h":
			m.stepRow(-1)
			return m, m.themeChangedCmd()

		case "right", "l", " ", "enter":
			m.stepRow(+1)
			return m, m.themeChangedCmd()

		case "s":
			cfg := m.toConfig()
			if err := storage.SaveConfig(cfg); err != nil {
				// Non-fatal: still navigate away.
			}
			return m, func() tea.Msg { return SaveSettingsMsg{Config: cfg} }
		}
	}
	return m, nil
}

func (m *SettingsModel) stepRow(delta int) {
	row := &m.rows[m.cursor]
	switch row.kind {
	case kindTheme, kindMode:
		row.idx = clamp(row.idx+delta, 0, len(row.options)-1)
	case kindHints:
		row.boolVal = !row.boolVal
	case kindLives:
		row.intVal = clamp(row.intVal+delta, row.intMin, row.intMax)
	}
}

func (m *SettingsModel) themeChangedCmd() tea.Cmd {
	row := m.rows[m.cursor]
	if row.kind != kindTheme {
		return nil
	}
	name := row.options[row.idx]
	return func() tea.Msg { return ApplyThemeMsg{ThemeName: name} }
}

func (m SettingsModel) View() string {
	title := m.titleStyle.Render("Settings")
	divider := m.dividerStyle.Render("────────────────────────")

	var sb strings.Builder
	for i, row := range m.rows {
		selected := i == m.cursor

		var label, value string
		if selected {
			label = m.selectedStyle.Render(row.label)
			value = m.selectedValStyle.Render(m.arrowWrap(row))
		} else {
			label = m.normalStyle.Render(row.label)
			value = m.valueStyle.Render(row.displayValue())
		}

		sb.WriteString(lipgloss.JoinHorizontal(lipgloss.Center,
			lipgloss.NewStyle().Width(20).Render(label),
			value,
		))
		sb.WriteString("\n\n")
	}

	footer := m.footerStyle.Render(
		"↑↓ / jk  navigate		←→ / hl  change		s  save		esc  cancel",
	)

	body := fmt.Sprintf("%s\n%s\n\n%s%s", title, divider, sb.String(), footer)
	return m.containerStyle.Render(body)
}

func (m *SettingsModel) arrowWrap(row settingRow) string {
	v := row.displayValue()
	switch row.kind {
	case kindTheme, kindMode:
		canLeft := row.idx > 0
		canRight := row.idx < len(row.options)-1
		left := " "
		right := " "
		if canLeft {
			left = "←"
		}
		if canRight {
			right = "→"
		}
		return fmt.Sprintf("%s %s %s", left, v, right)
	case kindLives:
		return fmt.Sprintf("← %s →", v)
	default:
		return fmt.Sprintf("← %s →", v)
	}
}

func (m *SettingsModel) toConfig() storage.Config {
	cfg := m.original
	for _, row := range m.rows {
		switch row.kind {
		case kindTheme:
			cfg.Theme = row.options[row.idx]
		case kindMode:
			cfg.Mode = row.options[row.idx]
		case kindLives:
			cfg.Lives = row.intVal
		case kindHints:
			cfg.ShowHints = row.boolVal
		}
	}
	return cfg
}

func sortedThemes() []string {
	// Deterministic order regardless of map iteration.
	all := theme.Available()
	// Sort alphabetically so the cycle order is predictable.
	for i := 1; i < len(all); i++ {
		for j := i; j > 0 && all[j] < all[j-1]; j-- {
			all[j], all[j-1] = all[j-1], all[j]
		}
	}
	return all
}

func indexOfOrZero(slice []string, val string) int {
	for i, s := range slice {
		if s == val {
			return i
		}
	}
	return 0
}

func clamp(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
