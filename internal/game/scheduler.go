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
	if p.EaseFactor < 1.3 {
		p.EaseFactor = 1.3
	}

	p.NextReview = time.Now().Add(time.Duration(p.Interval) * 24 * time.Hour)
}

func GetOrCreateProgress(data map[string]*storage.CharacterProgress, kana string) *storage.CharacterProgress {
	if _, exists := data[kana]; !exists {
		data[kana] = &storage.CharacterProgress{
			Kana:        kana,
			Repetitions: 0,
			Interval:    0,
			EaseFactor:  2.5, // Standard SM-2 starting constant
			NextReview:  time.Now(),
		}
	}
	return data[kana]
}
