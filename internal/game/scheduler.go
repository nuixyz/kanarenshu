package game

import (
	"math"
	"time"

	"github.com/nuixyz/kanarenshu/internal/storage"
)

func EvaluateSM2(p *storage.CharacterProgress, quality int) {
	p.Attempts++

	if quality < 3 {
		p.Repetitions = 0
		p.Interval = 1
	} else {
		p.Correct++
		p.Repetitions++

		switch p.Repetitions {
		case 1:
			p.Interval = 1
		case 2:
			p.Interval = 6
		default:
			p.Interval = int(math.Round(float64(p.Interval) * p.EaseFactor))
		}
	}

	qFactor := float64(5 - quality)
	p.EaseFactor = p.EaseFactor + (0.1 - qFactor*(0.08+qFactor*0.02))

	// 1.3 to prevent the intervals from collapsing completely
	if p.EaseFactor < 1.3 {
		p.EaseFactor = 1.3
	}

	p.NextReview = time.Now().Add(time.Duration(p.Interval) * 24 * time.Hour)
}

func GetOrCreateProgress(data map[string]*storage.CharacterProgress, kana string) *storage.CharacterProgress {
	// safety check here; initializes internal maps if they don't exists
	if p, exists := data[kana]; exists {
		if p.ModeAttempts == nil {
			p.ModeAttempts = make(map[string]int)
		}
		if p.ModeCorrect == nil {
			p.ModeCorrect = make(map[string]int)
		}
		return p
	}
	p := &storage.CharacterProgress{
		Kana:         kana,
		EaseFactor:   2.5,
		NextReview:   time.Now(),
		ModeAttempts: make(map[string]int),
		ModeCorrect:  make(map[string]int),
	}
	data[kana] = p
	return p
}
