package kanji

import "strings"

type Kanji struct {
	Char     string
	Onyomi   []string
	Kunyomi  []string
	Meanings []string
}

func (k Kanji) Display() string {
	return k.Char
}

// Hint returns a the primary meaning and the first on'yomi reading
func (k Kanji) Hint() string {
	parts := []string{}
	if len(k.Meanings) > 0 {
		parts = append(parts, k.Meanings[0])
	}
	if len(k.Onyomi) > 0 {
		parts = append(parts, k.Onyomi[0])
	}
	return strings.Join(parts, " · ")
}

// Check returns true if the answer matches any accepted reading or meaning
// Accepts on'yomi romaji, kun'yomi romaji, English meaning
func (k Kanji) Check(answer string) bool {
	answer = normalize(answer)

	if answer == "" {
		return false
	}
	for _, r := range k.Onyomi {
		if answer == normalize(r) {
			return true
		}
	}
	for _, r := range k.Kunyomi {
		if answer == normalize(r) {
			return true
		}
	}
	for _, r := range k.Meanings {
		if answer == normalize(r) {
			return true
		}
	}
	return false
}

func normalize(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}
