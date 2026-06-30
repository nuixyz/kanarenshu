package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/nuixyz/kanarenshu/internal/data"
	"github.com/nuixyz/kanarenshu/internal/logger"
)

type CharStat struct {
	Attempts int `json:"attempts"`
	Correct  int `json:"correct"`
}

type Progress struct {
	HighestLevel      map[string]int                 `json:"highest_level"`
	CharStats         map[string]map[string]CharStat `json:"char_stats"`
	HighestKanjiLevel map[string]int                 `json:"highest_kanji_level"`
}

func newProgress() *Progress {
	return &Progress{
		HighestLevel:      make(map[string]int),
		CharStats:         make(map[string]map[string]CharStat),
		HighestKanjiLevel: make(map[string]int),
	}
}

func progressFilePath() (string, error) {
	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("Could not determine home directory: %w", err)
		}
		dataHome = filepath.Join(home, ".local", "share")
	}
	dir := filepath.Join(dataHome, "kanarenshu")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("Could not create data directory %s: %w", dir, err)
	}
	return filepath.Join(dir, "progress.json"), nil
}

func Load() (*Progress, error) {
	path, err := progressFilePath()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(path)
	if os.IsNotExist(err) {
		logger.Info("No progress file found. Starting fresh.")
		return newProgress(), nil
	}

	if err != nil {
		return nil, fmt.Errorf("Could not open progress file: %w", err)
	}
	defer f.Close()

	p := newProgress()
	if err := json.NewDecoder(f).Decode(p); err != nil {
		return nil, fmt.Errorf("Could not decode progress file: %w", err)
	}

	if p.CharStats == nil {
		p.CharStats = make(map[string]map[string]CharStat)
	}
	if p.HighestKanjiLevel == nil {
		p.HighestKanjiLevel = make(map[string]int)
	}

	logger.Info("Progress loaded from path: %s", path)
	return p, nil
}

func Save(p *Progress) error {
	path, err := progressFilePath()
	if err != nil {
		return err
	}

	tmp := path + ".tmp"
	f, err := os.OpenFile(tmp, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("Could not create temporary progress file: %w", err)
	}

	enc := json.NewEncoder(f)
	enc.SetIndent("", " ")
	if err := enc.Encode(p); err != nil {
		_ = f.Close()
		_ = os.Remove(tmp)
		return fmt.Errorf("Could not encode progress: %w", err)
	}

	if err := f.Close(); err != nil {
		_ = os.Remove(tmp)
		return fmt.Errorf("Could not close temp progress file: %s", err)
	}

	if err := os.Rename(tmp, path); err != nil {
		_ = os.Remove(tmp)
		return fmt.Errorf("Could not save progress file: %s", err)
	}

	logger.Info("Progress saved to %s.", path)
	return nil
}

func RecordLevel(mode data.Mode, level int) error {
	p, err := Load()
	if err != nil {
		return err
	}

	key := mode.String()
	if level > p.HighestLevel[key] {
		p.HighestLevel[key] = level
		logger.Info("New highest level for %s: %d", key, level)
		return Save(p)
	}

	return nil
}

func HighestLevelFor(mode data.Mode) (int, error) {
	p, err := Load()
	if err != nil {
		return 0, err
	}
	return p.HighestLevel[mode.String()], nil
}

func RecordAttempt(mode data.Mode, kana string, correct bool) error {
	p, err := Load()
	if err != nil {
		return err
	}

	key := mode.String()
	if p.CharStats[key] == nil {
		p.CharStats[key] = make(map[string]CharStat)
	}

	stat := p.CharStats[key][kana]
	stat.Attempts++
	if correct {
		stat.Correct++
	}
	p.CharStats[key][kana] = stat

	return Save(p)
}

func CharStatsFor(mode data.Mode) (map[string]CharStat, error) {
	p, err := Load()
	if err != nil {
		return nil, err
	}
	return p.CharStats[mode.String()], nil
}

func RecordKanjiLevel(jlpt string, level int) error {
	p, err := Load()
	if err != nil {
		return err
	}
	if level > p.HighestKanjiLevel[jlpt] {
		p.HighestKanjiLevel[jlpt] = level
		logger.Info("New highest kanji level for %s: %d", jlpt, level)
		return Save(p)
	}
	return nil
}

func HighestKanjiLevelFor(jlpt string) (int, error) {
	p, err := Load()
	if err != nil {
		return 0, err
	}
	return p.HighestKanjiLevel[jlpt], nil
}

func RecordKanjiAttempts(jlpt string, char string, correct bool) error {
	p, err := Load()
	if err != nil {
		return err
	}

	key := "kanji_" + jlpt
	if p.CharStats[key] == nil {
		p.CharStats[key] = make(map[string]CharStat)
	}

	stat := p.CharStats[key][char]
	stat.Attempts++
	if correct {
		stat.Correct++
	}
	p.CharStats[key][char] = stat

	return Save(p)
}
