package game

import (
	"math/rand"
	"time"

	"github.com/nuixyz/kanarenshu/internal/data"
	"github.com/nuixyz/kanarenshu/internal/logger"
	"github.com/nuixyz/kanarenshu/internal/storage"
	"github.com/nuixyz/kanarenshu/pkg/romaji"
)

type AnswerResult int

const (
	AnswerCorrect AnswerResult = iota
	AnswerWrong
	AnswerGameOver
	AnswerLevelUp
)

const (
	initialWeight = 400
	masteryWeight = 25
)

type Config struct {
	Mode         data.Mode
	Lives        int
	StartLevel   int
	RomajiStrict bool
}

func DefaultConfig() Config {
	return Config{
		Mode:         data.ModeHiragana,
		Lives:        3,
		StartLevel:   0,
		RomajiStrict: false,
	}
}

type Session struct {
	cfg Config

	Level  int
	Lives  int
	Score  int
	Streak int
	Total  int

	pool         []romaji.Character // Characters in rotation
	current      romaji.Character   // Character currently being asked
	WrongAnswers []WrongAnswer      // To track the wrong answers in a session

	rng *rand.Rand

	weights map[string]int //per-character weights
	// per character attempt for levelling up
	// attempts map[string]int
	// correct  map[string]int
}

func NewSession(cfg Config) *Session {
	if cfg.StartLevel == 0 {
		if highest, err := storage.HighestLevelFor(cfg.Mode); err == nil {
			cfg.StartLevel = highest
		}
	}

	s := &Session{
		cfg:     cfg,
		Level:   cfg.StartLevel,
		Lives:   cfg.Lives,
		weights: make(map[string]int),
		rng:     rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	s.pool = data.PoolForLevel(cfg.Mode, s.Level)
	s.resetWeights()
	logger.Info("Session started: Mode=%s Level=%d Pool=%d Chars Lives=%d", cfg.Mode, s.Level+1, len(s.pool), cfg.Lives)

	s.pickNext()
	return s
}

func (s *Session) Current() romaji.Character {
	return s.current
}

// Submit will log 0 for correct answer, 1 for wrong answer, 2 for game over and 3 for level up
func (s *Session) Submit(answer string) AnswerResult {
	s.Total++

	isCorrect := romaji.KanaChecker(s.current, answer, s.cfg.RomajiStrict)

	if err := storage.RecordAttempt(s.cfg.Mode, s.current.Kana, isCorrect); err != nil {
		logger.Error("Failed to record character stats: %v", err)
	}

	if isCorrect {
		return s.handleCorrect()
	}
	return s.handleWrong()
}

func (s *Session) handleCorrect() AnswerResult {
	s.Score++
	s.Streak++
	// s.correct[s.current.Kana]++

	kana := s.current.Kana
	s.weights[kana] = s.weights[kana] / 2

	logger.Debug("Correct: %s (%s) Weight= %d", s.current.Kana, s.current.Primary, s.weights[kana])

	if s.shouldLevelUp() {
		s.levelUp()
		return AnswerLevelUp
	}

	s.pickNext()
	return AnswerCorrect
}

func (s *Session) handleWrong() AnswerResult {
	s.Streak = 0

	kana := s.current.Kana
	w := s.weights[kana] * 2
	if w > initialWeight {
		w = initialWeight
	}
	s.weights[kana] = w

	s.WrongAnswers = append(s.WrongAnswers, WrongAnswer{
		Char:   s.current.Kana,
		Answer: s.current.Primary,
	})

	logger.Debug("Wrong: %s (%s)", s.current.Kana, s.current.Primary)

	if s.cfg.Lives > 0 {
		s.Lives--
		if s.Lives <= 0 {
			logger.Info("Game Over! \nScore= %d Total=%d", s.Score, s.Total)
			return AnswerGameOver
		}
	}
	return AnswerWrong
}

// includes a weighted random selection instead of pure gacha
func (s *Session) pickNext() {
	if len(s.pool) == 1 {
		s.current = s.pool[0]
		return
	}

	type candidate struct {
		char   romaji.Character
		weight int
	}

	candidates := make([]candidate, 0, len(s.pool)-1)
	total := 0
	for _, c := range s.pool {
		if c.Kana == s.current.Kana {
			continue
		}
		w := s.weights[c.Kana]
		if w < 1 {
			w = 1 // floor??????? i am not sure what this does
		}
		candidates = append(candidates, candidate{c, w})
		total += w
	}

	r := s.rng.Intn(total)
	cum := 0

	for _, cand := range candidates {
		cum += cand.weight
		if r < cum {
			s.current = cand.char
			return
		}
	}

	s.current = candidates[len(candidates)-1].char
}

// At least 5 answers attempted per character in the current level group
// Accuracy accross group >= 90%
// func (s *Session) shouldLevelUp() bool {
// 	nextLevel := s.Level + 1
// 	if nextLevel >= data.TotalLevels(s.cfg.Mode) {
// 		return false
// 	}

// 	group := data.GroupForLevel(s.cfg.Mode, s.Level)
// 	if len(group) == 0 {
// 		return false
// 	}

// 	const minAttempts = 5
// 	const minAccuracy = 0.90

// 	totalAttempts := 0
// 	totalCorrect := 0

// 	for _, c := range group {
// 		att := s.attempts[c.Kana]
// 		cor := s.correct[c.Kana]

// 		if att < minAttempts {
// 			return false
// 		}

// 		totalAttempts += att
// 		totalCorrect += cor
// 	}

// 	if totalAttempts == 0 {
// 		return false
// 	}

// 	accuracy := float64(totalCorrect) / float64(totalAttempts)
// 	return accuracy >= minAccuracy
// }

func (s *Session) shouldLevelUp() bool {
	nextLevel := s.Level + 1
	if nextLevel >= data.TotalLevels(s.cfg.Mode) {
		return false
	}

	group := data.GroupForLevel(s.cfg.Mode, s.Level)
	if len(group) == 0 {
		return false
	}

	for _, c := range group {
		if s.weights[c.Kana] > masteryWeight {
			return false
		}
	}
	return true
}

func (s *Session) levelUp() {
	s.Level++

	if err := storage.RecordLevel(s.cfg.Mode, s.Level); err != nil {
		logger.Error("Failed to save progress: %v", err)
	}

	s.pool = data.PoolForLevel(s.cfg.Mode, s.Level)
	s.resetWeights()

	logger.Info("Level Up! New Level= %d Pool= %d Chars", s.Level, len(s.pool))
	s.pickNext()
}

func (s *Session) resetWeights() {
	for _, c := range s.pool {
		s.weights[c.Kana] = initialWeight
	}
}

func (s *Session) Accuracy() float64 {
	if s.Total == 0 {
		return 0
	}
	return float64(s.Score) / float64(s.Total)
}

func (s *Session) IsGameOver() bool {
	return s.cfg.Lives > 0 && s.Lives <= 0
}

func (s *Session) PoolSize() int {
	return len(s.pool)
}

func (s *Session) MaxLevel() int {
	return data.TotalLevels(s.cfg.Mode) - 1
}

func (s *Session) Cfg() Config {
	return s.cfg
}

func (s *Session) WeightFor(kana string) int {
	return s.weights[kana]
}
