package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/nuixyz/kanarenshu/internal/logger"
	"github.com/nuixyz/kanarenshu/internal/storage"
	"github.com/nuixyz/kanarenshu/internal/ui"
)

func main() {
	cleanup, err := logger.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not initialize logger: %v\n", err)
	} else {
		defer cleanup()
	}

	defer logger.RecoverAndLog(func(reason string) {
		fmt.Fprintln(os.Stderr, reason)
		os.Exit(1)
	})

	// Initialize the persistent storage layer
	store, err := storage.NewProgressStore()
	if err != nil {
		logger.Error("Could not initialize progress store: %v", err)
		fmt.Fprintf(os.Stderr, "Error: Could not initialize storage directory: %v\n", err)
		os.Exit(1)
	}

	// Load existing user progress data from progress.json
	if err := store.Load(); err != nil {
		logger.Error("Could not load progress data: %v", err)
		fmt.Fprintf(os.Stderr, "Error: Failed to read learning records: %v\n", err)
		os.Exit(1)
	}

	palette := ui.DefaultPalette()
	lives := 3

	root := ui.NewRenderer(palette, lives, store)

	p := tea.NewProgram(
		root,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	logger.Info("Starting kanarenshu...")

	if _, err := p.Run(); err != nil {
		logger.Error("Program exited with an error: %v", err)
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
