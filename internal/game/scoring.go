package game

import "fmt"

type Grade string

const (
	GradeS Grade = "S"
	GradeA Grade = "A"
	GradeB Grade = "B"
	GradeC Grade = "C"
	GradeD Grade = "D"
)

type Summary struct {
	Score    int
	Total    int
	Accuracy float64
	Grade    Grade
	MaxLevel int
	Mode     string
	JLPT     string
}

func Summarise(s *Session) Summary {
	acc := s.Accuracy()
	return Summary{
		Score:    s.Score,
		Total:    s.Total,
		Accuracy: acc,
		Grade:    gradeFor(acc),
		MaxLevel: s.Level,
		Mode:     s.cfg.Mode.String(),
		JLPT:     "",
	}
}

func gradeFor(accuracy float64) Grade {
	switch {
	case accuracy >= 1.0:
		return GradeS
	case accuracy >= 0.90:
		return GradeA
	case accuracy >= 0.75:
		return GradeB
	case accuracy >= 0.60:
		return GradeC
	default:
		return GradeD
	}
}

func (sum Summary) AccuracyPercent() string {
	return fmt.Sprintf("%.1f%%", sum.Accuracy*100)
}

func StreakBonus(streak int) int {
	return (streak / 10) * 5
}
