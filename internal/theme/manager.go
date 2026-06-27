package theme

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"
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

type baseTheme struct {
	Name   string `toml:"name"`
	Author string `toml:"author"`

	Bg      string `toml:"bg"`
	Fg      string `toml:"fg"`
	Accent  string `toml:"accent"`
	Correct string `toml:"correct"`
	Wrong   string `toml:"wrong"`
	Muted   string `toml:"muted"`
	Border  string `toml:"border"`
	SelBg   string `toml:"sel_bg"`
}

func (b baseTheme) toPalette() Palette {
	return Palette{
		Bg:      b.Bg,
		Fg:      b.Fg,
		Accent:  b.Accent,
		Correct: b.Correct,
		Wrong:   b.Wrong,
		Muted:   b.Muted,
		Border:  b.Border,
		SelBg:   b.SelBg,
	}
}

//go:embed themes/tokyo-night.toml
var tokyoNightRaw []byte

//go:embed themes/catppuccin.toml
var catppuccinRaw []byte

//go:embed themes/nord.toml
var nordRaw []byte

var registry = map[string][]byte{
	"tokyo-night": tokyoNightRaw,
	"catppuccin":  catppuccinRaw,
	// "gruvbox":     gruvboxRaw,
	"nord": nordRaw,
	// "dracula":     draculaRaw,
}

func Available() []string {
	names := make([]string, 0, len(registry))
	for k := range registry {
		names = append(names, k)
	}
	return names
}

func Load(name string) (Palette, error) {
	if name == "" {
		name = "tokyo-night"
	}

	key := strings.ToLower(name)
	raw, ok := registry[key]
	if !ok {
		return DefaultPalette(), fmt.Errorf("theme %q not found", name)
	}

	var t baseTheme
	if err := toml.Unmarshal(raw, &t); err != nil {
		return DefaultPalette(), fmt.Errorf("failed to parse theme %q: %w", name, err)
	}

	return t.toPalette(), nil
}
