# かな練習 · kanarenshu

> Japanese kana practise for the terminal.

![screenshot of main menu](/assets/main_menu.png)

![screenshot of hiragana mode](/assets/hiragana_session.png)

![screenshot of settings menu](/assets/settings_menu.png)

kanarenshu is a terminal-based app for learning hiragana and katakana. 

---

## Features
- **Hiragana, Katakana, and Mixed** modes to choose from.
- **Weight based selection** - Characters you answer wrong appear more often, while mastered ones fade out.
- **Level based progression system** - Unlock new characters as you level up.
- **5 built-in themes** - Tokyo Night, Catppuccin, Gruvbox, Nord, Dracula.
- **Configurable** - lives, hints all editable via config file or the settings menu.

## Installation

### Nix
Nix users can directly use this repository to get the latest kanarenshu for their system.

Add in your `flake.nix`:
```nix
inputs.kanarenshu.url = "github:nuixyz/kanarenshu";
```

Pass inputs to your modules and then in `configuration.nix`:
```nix
environment.systemPackages = [
  inputs.kanarenshu.packages.${pkgs.stdenv.hostPlatform.system}.default
];
```

### From Source
```base
git clone https://github.com/nuixyz/kanarenshu
cd kanarenshu
go build ./cmd/kanarenshu
```

### Requirements

- Go 1.21 or later (uses `//go:embed`)

## Usage

```bash
./kanarenshu
```

Use the arrow keys or `j/k` to navigate the menu, 
`enter` to select.

### Keyboard shortcuts
 
| Key | Action |
|-----|--------|
| `enter` | Submit answer |
| `esc` | Return to menu |
| `ctrl+t` | Cycle through themes |
| `ctrl+c` | Quit |
| `q` | Quit (menu) |
| `s` | Open settings (menu) |

### Romaji input

Both strict Hepburn and common alternatives are accepted by default:
 
| Kana | Accepted |
|------|----------|
| し | `shi`, `si` |
| ち | `chi`, `ti` |
| つ | `tsu`, `tu` |
| ふ | `fu`, `hu` |
| じ | `ji`, `zi` |
| ん | `n`, `nn` |


## Progression

Characters are grouped into levels of 5. The program uses an adaptive learning algorithm to determine the weights of each character in a group.
Within a level, the character selection is weighted. A correct answer halves a character's weight; a wrong answer doubles it (capped to its initial value). Characters reach mastery when their weights fall below a certain threshold. All characters in a group must fall under the threshold to level up. 

The highest unlocked level is saved automatically and resumed on next launch.

---

## Configuration

On first run, kananrenshu creates a config file at:

```
~/.config/kanarenshu/config.toml
```

```toml
theme         = "tokyo-night"   # tokyo-night | catppuccin | gruvbox | nord | dracula
lives         = 3               # 0 = infinite
romaji_strict = false           # true = Hepburn only (e.g. "shi", not "si")
show_hints    = true            # show reading hint after a wrong answer
```

Changes take affect on next launch, or can be applied live via the in-app settings menu

---

## Themes

Themes can be changed in the settings menu (`s` from the main menu) or cycled in-session with `ctrl+t`. The selected theme is saved to config automatically on save.
 
| Name | |
|------|-|
| `tokyo-night` | Dark blue — default |
| `catppuccin` | Catppuccin Mocha |
| `gruvbox` | Gruvbox Dark |
| `nord` | Nord |
| `dracula` | Dracula |
 
---

## Data files
 
| Path | Contents |
|------|----------|
| `~/.config/kanarenshu/config.toml` | User configuration |
| `~/.local/share/kanarenshu/progress.json` | Level progress and per-character statistics |
| `~/.local/share/kanarenshu/debug.log` | Debug log |
 
XDG environment variables (`$XDG_CONFIG_HOME`, `$XDG_DATA_HOME`) are respected if set.
 
---

## Dependencies
 
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) — TUI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) — UI components
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) — Styles and layout
- [BurntSushi/toml](https://github.com/BurntSushi/toml) — TOML parsing

---
 
## License
 
MIT
