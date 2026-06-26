package storage

import (
	"fmt"
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/nuixyz/kanarenshu/internal/logger"
)

type CharacterProgress struct {
	Kana string `json:"kana"`
	Repetitions int `json:"repetitions"`
	Interval int `json:"interval"`
	EaseFactor float64 `json:"ease_factor"`
	NextReview time.Time `json:"next_review"`
	Attempts int `json:"attempts"`
	Correct int `json:"correct"`
}

type ProgressStore struct {
	filePath string
	Data map[string]*CharacterProgress `json:"data"`
}

func NewProgressStore() (*ProgressStore, error) {
	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("Could not find home directory: %w", err)
		}
		dataHome = filepath.Join(home, ".local", "share")
	}

	dir := filepath.Join(dataHome, "kanarenshu")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("Could not create storage dir: %w", err)
	}

	return &ProgressStore{
		filePath: filepath.Join(dir, "progress.json"),
		Data: make(map[string]*CharacterProgress),
	}, nil
}

func (ps *ProgressStore) Load() error {
	if _, err := os.Stat(ps.filePath); os.IsNotExist(err) {
		logger.Info("No previous progress. Starting a new progress.")
		return nil
	}

	bytes, err := os.ReadFile(ps.filePath)
	if err != nil {
		return fmt.Errorf("Failed to read progress: %w", err)
	}

	if err := json.Unmarshal(bytes, &ps.Data); err != nil {
		return fmt.Errorf("Failed to parse progress JSON: %w", err)
	}

	logger.Info("Progress loaded successfully. Tracked characters: %d", len(ps.Data))
	return nil
}


func (ps *ProgressStore) Save() error {
	bytes, err := json.MarshalIndent(ps.Data, "", " ")
	if err != nil {
		return fmt.Errorf("Failed to marshall progress: %w", err)
	}

	if err := os.WriteFile(ps.filePath, bytes, 0644); err != nil {
		return fmt.Errorf("Failed to write program file: %w", err)
	}

	logger.Debug("Progress autosaved to disk.")
	return nil
}