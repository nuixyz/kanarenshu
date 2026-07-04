package data

import "github.com/nuixyz/kanarenshu/pkg/romaji"

var KatakanaGroups [][]romaji.Character

func init() {
	groupSize := 5
	totalChars := len(romaji.Katakana)

	for i := 0; i < totalChars; i += groupSize {
		end := i + groupSize
		if end > totalChars {
			end = totalChars
		}
		KatakanaGroups = append(KatakanaGroups, romaji.Katakana[i:end])
	}
}

func KatakanaFlat() []romaji.Character {
	var all []romaji.Character
	for _, group := range KatakanaGroups {
		all = append(all, group...)
	}
	return all
}

func KatakanaUpToLevel(level int) []romaji.Character {
	var pool []romaji.Character
	for i := 0; i <= level && i < len(KatakanaGroups); i++ {
		pool = append(pool, KatakanaGroups[i]...)
	}
	return pool
}
