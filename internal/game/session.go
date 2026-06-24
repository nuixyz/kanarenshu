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

	store *storage.ProgressStore
}

func NewSession(cfg Config, store *storage.ProgressStore) *Session {
	startLevel := cfg.StartLevel
	modeKey := cfg.Mode.String()

	if startLevel == 0 && store.HighestUnlockedLevels != nil {
		if unlockedLevel, exists := store.HighestUnlockedLevels[modeKey]; exists && unlockedLevel > 0 {
			startLevel = unlockedLevel
		}
	}

	s := &Session{
		cfg:   cfg,
		Level: startLevel,
		Lives: cfg.Lives,
		store: store,
		rng:   rand.New(rand.NewSource(time.Now().UnixNano())),
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

	if romaji.KanaChecker(s.current, answer) {
		return s.handleCorrect()
	}
	return s.handleWrong()
}

func (s *Session) handleCorrect() AnswerResult {
	s.Score++
	s.Streak++

	p := GetOrCreateProgress(s.store.Data, s.current.Kana)
	quality := 4

	if s.Streak > 5 {
		quality = 5
	}

	EvaluateSM2(p, quality)
	_ = s.store.Save()

	logger.Debug("Correct: %s (%s) Streak= %d Score= %d | SM-2 Interval= %d days", s.current.Kana, s.current.Primary, s.Streak, s.Score, p.Interval)

	if s.shouldLevelUp() {
		s.levelUp()
		return AnswerLevelUp
	}

	s.pickNext()
	return AnswerCorrect
}

func (s *Session) handleWrong() AnswerResult {
	s.Streak = 0

	p := GetOrCreateProgress(s.store.Data, s.current.Kana)
	EvaluateSM2(p, 1) // Quality = 1 for incorrect answer
	_ = s.store.Save()

	logger.Debug("Wrong: %s (%s)", s.current.Kana, s.current.Primary)

	if s.cfg.Lives > 0 {
		s.Lives--
		if s.Lives <= 0 {
			logger.Info("Game Over! \nScore= %d Total=%d", s.Score, s.Total)
			return AnswerGameOver
		}
	}
	s.pickNext()
	return AnswerWrong
}

func (s *Session) pickNext() {
	if len(s.pool) == 1 {
		s.current = s.pool[0]
		return
	}

	now := time.Now()

	for {
		var next romaji.Character

		if s.rng.Intn(10) < 7 {
			var overdue []romaji.Character

			for _, c := range s.pool {
				if p, exists := s.store.Data[c.Kana]; exists {
					if now.After(p.NextReview) {
						overdue = append(overdue, c)
					}
				}
			}

			if len(overdue) > 0 {
				next = overdue[s.rng.Intn(len(overdue))]
			}
		}

		if next.Kana == "" {
			next = s.pool[s.rng.Intn(len(s.pool))]
		}

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
		p, exists := s.store.Data[c.Kana]
		if !exists || p.Attempts < minAttempts {
			return false
		}

		totalAttempts += p.Attempts
		totalCorrect += p.Correct
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
	modeKey := s.cfg.Mode.String()

	if s.store.HighestUnlockedLevels == nil {
		s.store.HighestUnlockedLevels = make(map[string]int)
	}

	// Update milestone tracking only for the active mode category
	if s.Level > s.store.HighestUnlockedLevels[modeKey] {
		s.store.HighestUnlockedLevels[modeKey] = s.Level
		_ = s.store.Save()
	}

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
