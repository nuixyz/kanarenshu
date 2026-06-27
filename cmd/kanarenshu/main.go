package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/nuixyz/kanarenshu/internal/logger"
	"github.com/nuixyz/kanarenshu/internal/storage"
	"github.com/nuixyz/kanarenshu/internal/theme"
	"github.com/nuixyz/kanarenshu/internal/ui"
)

func main() {
	cleanup, err := logger.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: could not initialise logger: %v\n", err)
	} else {
		defer cleanup()
	}

	defer logger.RecoverAndLog(func(reason string) {
		fmt.Fprintln(os.Stderr, reason)
		os.Exit(1)
	})

	// Load config — writes defaults on first run.
	cfg, err := storage.LoadConfig()
	if err != nil {
		logger.Error("Could not load config: %v — using defaults", err)
		cfg = storage.DefaultConfig()
	}
	logger.Info("Config loaded: theme=%s mode=%s lives=%d", cfg.Theme, cfg.Lives)

	// Resolve palette from config theme.
	palette, err := theme.Load(cfg.Theme)
	if err != nil {
		logger.Error("Could not load theme %q: %v — falling back to default", cfg.Theme, err)
		palette = theme.DefaultPalette()
	}

	root := ui.NewRenderer(palette, cfg)

	p := tea.NewProgram(
		root,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	logger.Info("Starting kanarenshu…")

	if _, err := p.Run(); err != nil {
		logger.Error("Program exited with an error: %v", err)
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
