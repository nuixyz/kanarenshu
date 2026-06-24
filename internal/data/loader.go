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
		var mixedPool []romaji.Character
		for i := 0; i <= level; i++ {
			mixedPool = append(mixedPool, GroupForLevel(ModeMixed, i)...)
		}
		return mixedPool
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
		var mixedGroup []romaji.Character
		if level < len(HiraganaGroups) {
			mixedGroup = append(mixedGroup, HiraganaGroups[level]...)
		}
		if level < len(KatakanaGroups) {
			mixedGroup = append(mixedGroup, KatakanaGroups[level]...)
		}
		return mixedGroup
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
