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

var version string // will be set by ldflags during build

func resolvedVersion() string {
	if version == "" {
		return "0.1.1"
	}
	return version
}

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		fmt.Printf("kanarenshu version %s\n", resolvedVersion())
		os.Exit(0)
	}

	cleanup, err := logger.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: could not initialise logger: %v\n", err)
	} else {
		defer cleanup()
	}

	// When the application panics, exit the application and print the log file
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
	logger.Info("kanarenshu %s started. Config loaded with theme=%s", version, cfg.Theme)

	// Resolve palette from config theme.
	palette, err := theme.Load(cfg.Theme)
	if err != nil {
		logger.Error("Could not load theme %q: %v — falling back to default", cfg.Theme, err)
		palette = theme.DefaultPalette()
	}

	root := ui.NewRenderer(palette, cfg)

	p := tea.NewProgram(root, tea.WithAltScreen(), tea.WithMouseCellMotion())

	if _, err := p.Run(); err != nil {
		logger.Error("Program exited with an error: %v", err)
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
