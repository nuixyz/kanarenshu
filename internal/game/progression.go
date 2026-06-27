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
	total := 0.0
	for _, c := range group {
		w := float64(s.WeightFor(c.Kana))

		if w > initialWeight {
			w = initialWeight
		}
		if w < masteryWeight {
			w = masteryWeight
		}

		charProgress := 1.0 - (w-masteryWeight)/(initialWeight-masteryWeight)
		total += charProgress
	}

	return total / float64(len(group))
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
