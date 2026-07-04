package data

import "github.com/nuixyz/kanarenshu/pkg/romaji"

var HiraganaGroups [][]romaji.Character

func init() {
	groupSize := 5
	totalChars := len(romaji.Hiragana)

	for i := 0; i < totalChars; i += groupSize {
		end := i + groupSize
		if end > totalChars {
			end = totalChars
		}
		HiraganaGroups = append(HiraganaGroups, romaji.Hiragana[i:end])
	}
}

func HiraganaFlat() []romaji.Character {
	var all []romaji.Character
	for _, group := range HiraganaGroups {
		all = append(all, group...)
	}
	return all
}

func HiraganaUpToLevel(level int) []romaji.Character {
	var pool []romaji.Character
	for i := 0; i <= level && i < len(HiraganaGroups); i++ {
		pool = append(pool, HiraganaGroups[i]...)
	}
	return pool
}
