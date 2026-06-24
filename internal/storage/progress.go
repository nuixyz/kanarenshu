package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/nuixyz/kanarenshu/internal/logger"
)

type CharacterProgress struct {
	Kana        string    `json:"kana"`
	Repetitions int       `json:"repetitions"`
	Interval    int       `json:"interval"`
	EaseFactor  float64   `json:"ease_factor"`
	NextReview  time.Time `json:"next_review"`
	Attempts    int       `json:"attempts"`
	Correct     int       `json:"correct"`
}

// ProgressStore handles persisting state to progress.json
type ProgressStore struct {
	filePath              string
	HighestUnlockedLevels map[string]int                `json:"highest_unlocked_level"`
	Data                  map[string]*CharacterProgress `json:"data"`
}

func NewProgressStore() (*ProgressStore, error) {
	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("could not determine home dir: %w", err)
		}
		dataHome = filepath.Join(home, ".local", "share")
	}

	dir := filepath.Join(dataHome, "kanarenshu")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("could not create storage dir: %w", err)
	}

	return &ProgressStore{
		filePath:              filepath.Join(dir, "progress.json"),
		HighestUnlockedLevels: make(map[string]int),
		Data:                  make(map[string]*CharacterProgress),
	}, nil
}

func (ps *ProgressStore) Load() error {
	if _, err := os.Stat(ps.filePath); os.IsNotExist(err) {
		logger.Info("Progress file not found. Starting fresh.")
		return nil
	}

	bytes, err := os.ReadFile(ps.filePath)
	if err != nil {
		return fmt.Errorf("failed to read progress: %w", err)
	}

	if err := json.Unmarshal(bytes, ps); err != nil {
		return fmt.Errorf("failed to parse progress JSON: %w", err)
	}

	logger.Info("Progress loaded. Highest Level: %v, Tracked characters: %d", ps.HighestUnlockedLevels, len(ps.Data)) //
	return nil
}

func (ps *ProgressStore) Save() error {
	bytes, err := json.MarshalIndent(ps, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal progress: %w", err)
	}

	if err := os.WriteFile(ps.filePath, bytes, 0644); err != nil {
		return fmt.Errorf("failed to write progress file: %w", err)
	}

	logger.Debug("Progress autosaved to disk.")
	return nil
}
