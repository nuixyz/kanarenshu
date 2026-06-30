package game

import (
	"fmt"

	"github.com/nuixyz/kanarenshu/internal/data"
)

func KanjiLevelLabel(level int) string {
	return fmt.Sprintf("Level %d", level+1)
}

func KanjiLevelProgress(s *KanjiSession) float64 {
	groups := data.KanjiGroupsForJLPT(s.cfg.JLPT)
	if s.Level >= len(groups) {
		return 0
	}
	group := groups[s.Level]
	if len(group) == 0 {
		return 0
	}

	total := 0.0
	for _, k := range group {
		w := float64(s.WeightFor(k.Char))
		if w > initialWeight {
			w = initialWeight
		}
		if w < masteryWeight {
			w = masteryWeight
		}
		total += 1.0 - (w-masteryWeight)/(initialWeight-masteryWeight)
	}
	return total / float64(len(group))
}

func NewKanjiForLevel(jlpt string, level int) []string {
	groups := data.KanjiGroupsForJLPT(jlpt)
	if level >= len(groups) {
		return nil
	}
	group := groups[level]
	labels := make([]string, 0, len(group))
	for _, k := range group {
		reading := ""
		if len(k.Onyomi) > 0 {
			reading = k.Onyomi[0]
		}
		labels = append(labels, fmt.Sprintf("%s (%s)", k.Char, reading))
	}
	return labels
}

func IsMaxKanjiLevel(s *KanjiSession) bool {
	return s.Level > s.MaxLevel()
}

func SummariseKanji(s *KanjiSession) Summary {
	acc := s.Accuracy()
	return Summary{
		Score:    s.Score,
		Total:    s.Total,
		Accuracy: acc,
		Grade:    gradeFor(acc),
		MaxLevel: s.Level,
		Mode:     "Kanji " + s.cfg.JLPT,
	}
}
