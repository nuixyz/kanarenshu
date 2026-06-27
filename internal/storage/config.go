package storage

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Theme        string `toml:"theme"`
	Lives        int    `toml:"lives"`
	RomajiStrict bool   `toml:"romaji_strict"`
	ShowHints    bool   `toml:"show_hints"`
}

func DefaultConfig() Config {
	return Config{
		Theme:        "tokyo-night",
		Lives:        3,
		RomajiStrict: false,
		ShowHints:    true,
	}
}

func configFilePath() (string, error) {
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("Could not determine home directory: %w", err)
		}
		configHome = filepath.Join(home, ".config")
	}
	dir := filepath.Join(configHome, "kanarenshu")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("Could not create config dir: %w", err)
	}
	return filepath.Join(dir, "config.toml"), nil
}

func LoadConfig() (Config, error) {
	path, err := configFilePath()
	if err != nil {
		return DefaultConfig(), err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		cfg := DefaultConfig()
		if writeErr := writeConfig(path, cfg); writeErr != nil {
			return cfg, nil
		}
		return cfg, nil
	}

	var cfg Config
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return DefaultConfig(), fmt.Errorf("Could not parse config file %s: %w", path, err)
	}

	applyDefaults(&cfg)
	return cfg, nil
}

func SaveConfig(cfg Config) error {
	path, err := configFilePath()
	if err != nil {
		return err
	}
	return writeConfig(path, cfg)
}

func writeConfig(path string, cfg Config) error {
	tmp := path + ".tmp"
	f, err := os.OpenFile(tmp, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("Could not open temp config file: %w", err)
	}

	content := fmt.Sprintf(`# kanarenshu configuration
# github.com/nuixyz/kanarenshu

theme         = %q  # tokyo-night | catppuccin | gruvbox | nord | dracula
lives         = %d              # 0 = infinite
romaji_strict = %t          # true = Hepburn only (e.g. "shi", not "si")
show_hints    = %t           # show reading hint after a wrong answer
`,
		cfg.Theme, cfg.Lives, cfg.RomajiStrict, cfg.ShowHints,
	)

	if _, err := f.WriteString(content); err != nil {
		_ = f.Close()
		_ = os.Remove(tmp)
		return fmt.Errorf("Could not close temp config file: %w", err)
	}

	if err := f.Close(); err != nil {
		_ = os.Remove(tmp)
		return fmt.Errorf("could not close temp config file: %w", err)
	}
	if err := os.Rename(tmp, path); err != nil {
		_ = os.Remove(tmp)
		return fmt.Errorf("could not save config file: %w", err)
	}
	return nil
}

// applyDefaults fills zero-value fields so configs written by older
// versions of the app stay valid after an upgrade.
func applyDefaults(cfg *Config) {
	d := DefaultConfig()
	if cfg.Theme == "" {
		cfg.Theme = d.Theme
	}
	if cfg.Lives == 0 {
		cfg.Lives = d.Lives
	}
}
