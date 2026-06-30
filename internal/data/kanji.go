package data

import "github.com/nuixyz/kanarenshu/pkg/kanji"

var KanjiSets = map[string][]kanji.Kanji{
	"N5": kanji.N5,
}

// This divides each inner slice into groups of 10 characters
var KanjiGroups = map[string][][]kanji.Kanji{
	"N5": kanji.N5Groups,
}

// Returns a sorted list of JLPT levels that contain some data
func AvailableJLPT() []string {
	order := []string{"N5", "N4", "N3", "N2", "N1"}
	var available []string
	for _, level := range order {
		if _, ok := KanjiSets[level]; ok {
			available = append(available, level)
		}
	}
	return available
}

// Returns the level groups for a given JLPT level
func KanjiGroupsForJLPT(jlpt string) [][]kanji.Kanji {
	return KanjiGroups[jlpt]
}

// Returns all kanji unlocked upto and including that level
func KanjiUptoLevel(jlpt string, level int) []kanji.Kanji {
	groups, ok := KanjiGroups[jlpt]
	if !ok {
		return nil
	}
	var pool []kanji.Kanji
	for i := 0; i <= level && i < len(groups); i++ {
		pool = append(pool, groups[i]...)
	}
	return pool
}

// Return the total number of levels in a JLPT tier
func TotalKanjiLevels(jlpt string) int {
	return len(KanjiGroups[jlpt])
}
