package game

import (
	"fmt"

	"github.com/nuixyz/kanarenshu/internal/data"
)

func LevelLabel(level int) string {
	return fmt.Sprintf("Level %d", level+1)
}

func LevelProgress(s *Session) float64 {
	group := data.GroupForLevel(s.cfg.Mode, s.Level)
	if len(group) == 0 {
		return 0
	}

	const minAttempts = 5
	const targetAccuracy = 0.90

	totalScore := 0.0

	for _, c := range group {
		p, exists := s.store.Data[c.Kana]
		if !exists || p.Attempts == 0 {
			continue
		}

		att := p.Attempts
		cor := p.Correct
		acc := float64(cor) / float64(att)

		attemptProgress := float64(att) / float64(minAttempts)
		if attemptProgress > 1.0 {
			attemptProgress = 1.0
		}

		charScore := (acc / targetAccuracy) * attemptProgress
		if charScore > 1.0 {
			charScore = 1.0
		}

		totalScore += charScore
	}

	progress := totalScore / float64(len(group))
	if progress > 1.0 {
		return 1.0
	}
	return progress
}

func NewCharsForLevel(mode data.Mode, level int) []string {
	group := data.GroupForLevel(mode, level)
	labels := make([]string, 0, len(group))
	for _, c := range group {
		labels = append(labels, fmt.Sprintf("%s (%s)", c.Kana, c.Primary))
	}
	return labels
}

func IsMaxLevel(s *Session) bool {
	return s.Level >= s.MaxLevel()
}
