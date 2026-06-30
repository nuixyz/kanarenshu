package game

import (
	"math/rand"
	"time"

	"github.com/nuixyz/kanarenshu/internal/data"
	"github.com/nuixyz/kanarenshu/internal/logger"
	"github.com/nuixyz/kanarenshu/internal/storage"
	"github.com/nuixyz/kanarenshu/pkg/kanji"
)

type KanjiConfig struct {
	JLPT       string
	Lives      int
	StartLevel int
}

func DefaultKanjiConfig(jlpt string) KanjiConfig {
	return KanjiConfig{
		JLPT:       jlpt,
		Lives:      3,
		StartLevel: 0,
	}
}

type KanjiSession struct {
	cfg KanjiConfig

	Level  int
	Lives  int
	Score  int
	Streak int
	Total  int

	pool    []kanji.Kanji
	current kanji.Kanji

	rng *rand.Rand

	weights map[string]int
}

func NewKanjiSession(cfg KanjiConfig) *KanjiSession {
	if cfg.StartLevel == 0 {
		if highest, err := storage.HighestKanjiLevelFor(cfg.JLPT); err == nil {
			cfg.StartLevel = highest
		}
	}

	s := &KanjiSession{
		cfg:     cfg,
		Level:   cfg.StartLevel,
		Lives:   cfg.Lives,
		weights: make(map[string]int),
		rng:     rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	s.pool = data.KanjiUptoLevel(cfg.JLPT, s.Level)
	s.resetWeights()
	logger.Info("Kanji session started: JLPT=%s Level=%d Pool=%d chars Lives=%d", cfg.JLPT, s.Level+1, len(s.pool), cfg.Lives)

	s.pickNext()
	return s
}

func (s *KanjiSession) Current() kanji.Kanji {
	return s.current
}

func (s *KanjiSession) Submit(answer string) AnswerResult {
	s.Total++

	isCorrect := s.current.Check(answer)

	if err := storage.RecordKanjiAttempts(s.cfg.JLPT, s.current.Char, isCorrect); err != nil {
		logger.Error("Failed to record kanji stats: %v", err)
	}

	if isCorrect {
		return s.handleCorrect()
	}
	return s.handleWrong()
}

func (s *KanjiSession) handleCorrect() AnswerResult {
	s.Score++
	s.Streak++

	char := s.current.Char
	s.weights[char] = s.weights[char] / 2

	logger.Debug("Correct: %s Weight: %d", char, s.weights[char])

	if s.shouldLevelUp() {
		s.levelUp()
		return AnswerLevelUp
	}

	s.pickNext()
	return AnswerCorrect
}

func (s *KanjiSession) handleWrong() AnswerResult {
	s.Streak = 0

	char := s.current.Char
	w := s.weights[char] * 2
	if w > initialWeight {
		w = initialWeight
	}
	s.weights[char] = w

	logger.Debug("Wrong: %s", char)

	if s.cfg.Lives > 0 {
		s.Lives--
		if s.Lives <= 0 {
			logger.Info("Game Over! Score= %d Total= %d", s.Score, s.Total)
			return AnswerGameOver
		}
	}
	return AnswerWrong
}

func (s *KanjiSession) pickNext() {
	if len(s.pool) == 1 {
		s.current = s.pool[0]
		return
	}

	type candidate struct {
		k      kanji.Kanji
		weight int
	}

	candidates := make([]candidate, 0, len(s.pool)-1)
	total := 0
	for _, k := range s.pool {
		if k.Char == s.current.Char {
			continue
		}
		w := s.weights[k.Char]
		if w < 1 {
			w = 1
		}
		candidates = append(candidates, candidate{k, w})
		total += w
	}

	r := s.rng.Intn(total)
	cum := 0
	for _, cand := range candidates {
		cum += cand.weight
		if r < cum {
			s.current = cand.k
			return
		}
	}
	s.current = candidates[len(candidates)-1].k
}

func (s *KanjiSession) shouldLevelUp() bool {
	nextLevel := s.Level + 1
	if nextLevel >= data.TotalKanjiLevels(s.cfg.JLPT) {
		return false
	}

	groups := data.KanjiGroupsForJLPT(s.cfg.JLPT)
	if s.Level >= len(groups) {
		return false
	}
	group := groups[s.Level]
	if len(group) == 0 {
		return false
	}

	for _, k := range group {
		if s.weights[k.Char] > masteryWeight {
			return false
		}
	}
	return true
}

func (s *KanjiSession) levelUp() {
	s.Level++

	if err := storage.RecordKanjiLevel(s.cfg.JLPT, s.Level); err != nil {
		logger.Error("Failed to save kanji progress: %s", err)
	}

	s.pool = data.KanjiUptoLevel(s.cfg.JLPT, s.Level)
	s.resetWeights()

	logger.Info("Kanji level up! New Level=%d Pool=%d chars", s.Level, len(s.pool))
	s.pickNext()
}

func (s *KanjiSession) resetWeights() {
	for _, k := range s.pool {
		s.weights[k.Char] = initialWeight
	}
}
func (s *KanjiSession) Accuracy() float64 {
	if s.Total == 0 {
		return 0
	}
	return float64(s.Score) / float64(s.Total)
}

func (s *KanjiSession) IsGameOver() bool {
	return s.cfg.Lives > 0 && s.Lives <= 0
}

func (s *KanjiSession) PoolSize() int {
	return len(s.pool)
}

func (s *KanjiSession) MaxLevel() int {
	return data.TotalKanjiLevels(s.cfg.JLPT) - 1
}

func (s *KanjiSession) Cfg() KanjiConfig {
	return s.cfg
}

func (s *KanjiSession) WeightFor(char string) int {
	return s.weights[char]
}
