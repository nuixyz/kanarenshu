package romaji

import "strings"

type Character struct {
	Kana    string
	Primary string
	Alts    []string
}

func Normalize(input string) string {
	return strings.ToLower(strings.TrimSpace(input))
}

func kanaChecker(c Character, answer string) bool {
	answer = Normalize(answer)
	if answer == c.Primary {
		return true
	}
	for _, alt := range c.Alts {
		if answer == alt {
			return true
		}
	}
	return false
}

var Hiragama = []Character{
	{Kana: "あ", Primary: "a"},
	{Kana: "い", Primary: "i"},
	{Kana: "う", Primary: "u"},
	{Kana: "え", Primary: "e"},
	{Kana: "お", Primary: "o"},

	{Kana: "か", Primary: "ka"},
	{Kana: "き", Primary: "ki"},
	{Kana: "く", Primary: "ku"},
	{Kana: "け", Primary: "ke"},
	{Kana: "こ", Primary: "ko"},

	{Kana: "さ", Primary: "sa"},
	{Kana: "し", Primary: "shi", Alts: []string{"si"}},
	{Kana: "す", Primary: "su"},
	{Kana: "せ", Primary: "se"},
	{Kana: "そ", Primary: "so"},

	{Kana: "た", Primary: "ko"},
	{Kana: "ち", Primary: "ko"},
	{Kana: "つ", Primary: "ko"},
	{Kana: "て", Primary: "ko"},
	{Kana: "と", Primary: "ko"},

	{Kana: "こ", Primary: "ko"},
	{Kana: "こ", Primary: "ko"},
	{Kana: "こ", Primary: "ko"},
	{Kana: "こ", Primary: "ko"},
	{Kana: "こ", Primary: "ko"},
}
