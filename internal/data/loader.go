package data

import "github.com/nuixyz/kanarenshu/pkg/romaji"

type Mode int

const (
	ModeHiragana Mode = iota
	ModeKatakana
	ModeMixed
)

// Return a human readable label for the mode.
func (m Mode) String() string {
	switch m {
	case ModeHiragana:
		return "Hiragana"
	case ModeKatakana:
		return "Katakana"
	case ModeMixed:
		return "Mixed"
	default:
		return "Unknown"
	}
}

// TotalLevels returns the total number of levels for the given mode.
func TotalLevels(mode Mode) int {
	switch mode {
	case ModeHiragana:
		return len(HiraganaGroups)
	case ModeKatakana:
		return len(KatakanaGroups)
	case ModeMixed:
		h := len(HiraganaGroups)
		k := len(KatakanaGroups)
		if h > k {
			return h
		}
		return k
	default:
		return 0
	}
}

// PoolForLevel returns all characters unlocked up to and including the given level.
func PoolForLevel(mode Mode, level int) []romaji.Character {
	switch mode {
	case ModeHiragana:
		return HiraganaUpToLevel(level)
	case ModeKatakana:
		return KatakanaUpToLevel(level)
	case ModeMixed:
		h := HiraganaUpToLevel(level)
		k := KatakanaUpToLevel(level)
		return append(h, k...)
	default:
		return nil
	}
}

// GroupForLevel returns the group of characters for the given mode and level.
func GroupForLevel(mode Mode, level int) []romaji.Character {
	switch mode {
	case ModeHiragana:
		if level < len(HiraganaGroups) {
			return HiraganaGroups[level]
		}
		return nil
	case ModeKatakana:
		if level < len(KatakanaGroups) {
			return KatakanaGroups[level]
		}
		return nil
	case ModeMixed:
		var group []romaji.Character
		if level < len(HiraganaGroups) {
			group = append(group, HiraganaGroups[level]...)
		}
		if level < len(KatakanaGroups) {
			group = append(group, KatakanaGroups[level]...)
		}
		return group
	default:
		return nil
	}
}

// AllCharacters returns every character in a mode as a flat slice.
func AllCharacters(mode Mode) []romaji.Character {
	switch mode {
	case ModeHiragana:
		return HiraganaFlat()
	case ModeKatakana:
		return KatakanaFlat()
	case ModeMixed:
		return append(HiraganaFlat(), KatakanaFlat()...)
	default:
		return nil
	}
}
