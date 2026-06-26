package game

import (
	"math/rand"
	"time"

	"github.com/nuixyz/kanarenshu/internal/data"
	"github.com/nuixyz/kanarenshu/internal/logger"
	"github.com/nuixyz/kanarenshu/pkg/romaji"
)

type AnswerResult int

const (
	AnswerCorrect AnswerResult = iota
	AnswerWrong
	AnswerGameOver
	AnswerLevelUp
)

type Config struct {
	Mode       data.Mode
	Lives      int
	StartLevel int
}

func DefaultConfig() Config {
	return Config{
		Mode:       data.ModeHiragana,
		Lives:      3,
		StartLevel: 0,
	}
}

type Session struct {
	cfg Config

	Level  int
	Lives  int
	Score  int
	Streak int
	Total  int

	pool    []romaji.Character // Characters in rotation
	current romaji.Character   // Character currently being asked

	rng *rand.Rand

	// per character attempt for levelling up
	attempts map[string]int
	correct  map[string]int
}

func NewSession(cfg Config) *Session {
	s := &Session{
		cfg:      cfg,
		Level:    cfg.StartLevel,
		Lives:    cfg.Lives,
		attempts: make(map[string]int),
		correct:  make(map[string]int),
		rng:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	s.pool = data.PoolForLevel(cfg.Mode, s.Level)
	logger.Info("Session started: Mode=%s Level=%d Pool=%d Chars Lives=%d", cfg.Mode, s.Level, len(s.pool), cfg.Lives)

	s.pickNext()
	return s
}

func (s *Session) Current() romaji.Character {
	return s.current
}

func (s *Session) Submit(answer string) AnswerResult {
	s.Total++
	s.attempts[s.current.Kana]++

	if romaji.KanaChecker(s.current, answer) {
		return s.handleCorrect()
	}
	return s.handleWrong()
}

func (s *Session) handleCorrect() AnswerResult {
	s.Score++
	s.Streak++
	s.correct[s.current.Kana]++

	logger.Debug("Correct: %s (%s) Streak= %d Score= %d", s.current.Kana, s.current.Primary, s.Streak, s.Score)

	if s.shouldLevelUp() {
		s.levelUp()
		return AnswerLevelUp
	}

	s.pickNext()
	return AnswerCorrect
}

func (s *Session) handleWrong() AnswerResult {
	s.Streak = 0

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

func (s *Session) pickNext() {
	if len(s.pool) == 1 {
		s.current = s.pool[0]
		return
	}

	for {
		next := s.pool[s.rng.Intn(len(s.pool))]
		if next.Kana != s.current.Kana {
			s.current = next
			return
		}
	}
}

// At least 5 answers attempted per character in the current level group
// Accuracy accross group >= 90%
func (s *Session) shouldLevelUp() bool {
	nextLevel := s.Level + 1
	if nextLevel >= data.TotalLevels(s.cfg.Mode) {
		return false
	}

	group := data.GroupForLevel(s.cfg.Mode, s.Level)
	if len(group) == 0 {
		return false
	}

	const minAttempts = 5
	const minAccuracy = 0.90

	totalAttempts := 0
	totalCorrect := 0

	for _, c := range group {
		att := s.attempts[c.Kana]
		cor := s.correct[c.Kana]

		if att < minAttempts {
			return false
		}

		totalAttempts += att
		totalCorrect += cor
	}

	if totalAttempts == 0 {
		return false
	}

	accuracy := float64(totalCorrect) / float64(totalAttempts)
	return accuracy >= minAccuracy
}

func (s *Session) levelUp() {
	s.Level++
	s.pool = data.PoolForLevel(s.cfg.Mode, s.Level)

	logger.Info("Level Up! New Level= %d Pool= %d Chars", s.Level, len(s.pool))
	s.pickNext()
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
