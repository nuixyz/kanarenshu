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

func KanaChecker(c Character, answer string) bool {
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

var Hiragana = []Character{
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

	{Kana: "た", Primary: "ta"},
	{Kana: "ち", Primary: "chi", Alts: []string{"ti"}},
	{Kana: "つ", Primary: "tsu", Alts: []string{"tu"}},
	{Kana: "て", Primary: "te"},
	{Kana: "と", Primary: "to"},

	{Kana: "な", Primary: "na"},
	{Kana: "に", Primary: "ni"},
	{Kana: "ぬ", Primary: "nu"},
	{Kana: "ね", Primary: "ne"},
	{Kana: "の", Primary: "no"},

	{Kana: "は", Primary: "ha"},
	{Kana: "ひ", Primary: "hi"},
	{Kana: "ふ", Primary: "fu", Alts: []string{"hu"}},
	{Kana: "へ", Primary: "he"},
	{Kana: "ほ", Primary: "ho"},

	{Kana: "ま", Primary: "ma"},
	{Kana: "み", Primary: "mi"},
	{Kana: "む", Primary: "mu"},
	{Kana: "め", Primary: "me"},
	{Kana: "も", Primary: "mo"},

	{Kana: "や", Primary: "ya"},
	{Kana: "ゆ", Primary: "yu"},
	{Kana: "よ", Primary: "yo"},

	{Kana: "ら", Primary: "ra"},
	{Kana: "り", Primary: "ri"},
	{Kana: "る", Primary: "ru"},
	{Kana: "れ", Primary: "re"},
	{Kana: "ろ", Primary: "ro"},

	{Kana: "わ", Primary: "wa"},
	{Kana: "を", Primary: "wo", Alts: []string{"o"}},

	{Kana: "ん", Primary: "n", Alts: []string{"nn"}},

	{Kana: "が", Primary: "ga"},
	{Kana: "ぎ", Primary: "gi"},
	{Kana: "ぐ", Primary: "gu"},
	{Kana: "げ", Primary: "ge"},
	{Kana: "ご", Primary: "go"},

	{Kana: "ざ", Primary: "za"},
	{Kana: "じ", Primary: "ji", Alts: []string{"zi"}},
	{Kana: "ず", Primary: "zu"},
	{Kana: "ぜ", Primary: "ze"},
	{Kana: "ぞ", Primary: "zo"},

	{Kana: "だ", Primary: "da"},
	{Kana: "ぢ", Primary: "di", Alts: []string{"zi", "ji"}},
	{Kana: "づ", Primary: "du", Alts: []string{"zu"}},
	{Kana: "で", Primary: "de"},
	{Kana: "ど", Primary: "do"},

	{Kana: "ば", Primary: "ba"},
	{Kana: "び", Primary: "bi"},
	{Kana: "ぶ", Primary: "bu"},
	{Kana: "べ", Primary: "be"},
	{Kana: "ぼ", Primary: "bo"},

	{Kana: "ぱ", Primary: "pa"},
	{Kana: "ぴ", Primary: "pi"},
	{Kana: "ぷ", Primary: "pu"},
	{Kana: "ぺ", Primary: "pe"},
	{Kana: "ぽ", Primary: "po"},

	{Kana: "きゃ", Primary: "kya"},
	{Kana: "きゅ", Primary: "kyu"},
	{Kana: "きょ", Primary: "kyo"},
	{Kana: "ぎゃ", Primary: "gya"},
	{Kana: "ぎゅ", Primary: "gyu"},
	{Kana: "ぎょ", Primary: "gyo"},

	{Kana: "しゃ", Primary: "sha"},
	{Kana: "しゅ", Primary: "shu"},
	{Kana: "しょ", Primary: "sho"},
	{Kana: "じゃ", Primary: "ja"},
	{Kana: "じゅ", Primary: "ju"},
	{Kana: "じょ", Primary: "jo"},

	{Kana: "ちゃ", Primary: "cha"},
	{Kana: "ちゅ", Primary: "chu"},
	{Kana: "ちょ", Primary: "cho"},

	{Kana: "にゃ", Primary: "nya"},
	{Kana: "にゅ", Primary: "nyu"},
	{Kana: "にょ", Primary: "nyo"},

	{Kana: "ひゃ", Primary: "hya"},
	{Kana: "ひゅ", Primary: "hyu"},
	{Kana: "ひょ", Primary: "hyo"},
	{Kana: "びゃ", Primary: "bya"},
	{Kana: "びゅ", Primary: "byu"},
	{Kana: "びょ", Primary: "byo"},
	{Kana: "ぴゃ", Primary: "pya"},
	{Kana: "ぴゅ", Primary: "pyu"},
	{Kana: "ぴょ", Primary: "pyo"},

	{Kana: "みょ", Primary: "mya"},
	{Kana: "みゅ", Primary: "myu"},
	{Kana: "みょ", Primary: "myo"},

	{Kana: "りゃ", Primary: "rya"},
	{Kana: "りゅ", Primary: "ryu"},
	{Kana: "りょ", Primary: "ryo"},
}

var Katakana = []Character{
	{Kana: "ア", Primary: "a"},
	{Kana: "イ", Primary: "i"},
	{Kana: "ウ", Primary: "u"},
	{Kana: "エ", Primary: "e"},
	{Kana: "オ", Primary: "o"},

	{Kana: "カ", Primary: "ka"},
	{Kana: "キ", Primary: "ki"},
	{Kana: "ク", Primary: "ku"},
	{Kana: "ケ", Primary: "ke"},
	{Kana: "コ", Primary: "ko"},

	{Kana: "サ", Primary: "sa"},
	{Kana: "シ", Primary: "shi", Alts: []string{"si"}},
	{Kana: "ス", Primary: "su"},
	{Kana: "セ", Primary: "se"},
	{Kana: "ソ", Primary: "so"},

	{Kana: "タ", Primary: "ta"},
	{Kana: "チ", Primary: "chi", Alts: []string{"ti"}},
	{Kana: "ツ", Primary: "tsu", Alts: []string{"tu"}},
	{Kana: "テ", Primary: "te"},
	{Kana: "ト", Primary: "to"},

	{Kana: "ナ", Primary: "na"},
	{Kana: "ニ", Primary: "ni"},
	{Kana: "ヌ", Primary: "nu"},
	{Kana: "ネ", Primary: "ne"},
	{Kana: "ノ", Primary: "no"},

	{Kana: "ハ", Primary: "ha"},
	{Kana: "ヒ", Primary: "hi"},
	{Kana: "フ", Primary: "fu", Alts: []string{"hu"}},
	{Kana: "ヘ", Primary: "he"},
	{Kana: "ホ", Primary: "ho"},

	{Kana: "マ", Primary: "ma"},
	{Kana: "ミ", Primary: "mi"},
	{Kana: "ム", Primary: "mu"},
	{Kana: "メ", Primary: "me"},
	{Kana: "モ", Primary: "mo"},

	{Kana: "ヤ", Primary: "ya"},
	{Kana: "ユ", Primary: "yu"},
	{Kana: "ヨ", Primary: "yo"},

	{Kana: "ラ", Primary: "ra"},
	{Kana: "リ", Primary: "ri"},
	{Kana: "ル", Primary: "ru"},
	{Kana: "レ", Primary: "re"},
	{Kana: "ロ", Primary: "ro"},

	{Kana: "ワ", Primary: "wa"},
	{Kana: "ヲ", Primary: "wo", Alts: []string{"o"}},

	{Kana: "ン", Primary: "n", Alts: []string{"nn"}},

	{Kana: "ガ", Primary: "ga"},
	{Kana: "ギ", Primary: "gi"},
	{Kana: "グ", Primary: "gu"},
	{Kana: "ゲ", Primary: "ge"},
	{Kana: "ゴ", Primary: "go"},

	{Kana: "ザ", Primary: "za"},
	{Kana: "ジ", Primary: "ji", Alts: []string{"zi"}},
	{Kana: "ズ", Primary: "zu"},
	{Kana: "ゼ", Primary: "ze"},
	{Kana: "ゾ", Primary: "zo"},

	{Kana: "ダ", Primary: "da"},
	{Kana: "ヂ", Primary: "di", Alts: []string{"zi", "ji"}},
	{Kana: "ヅ", Primary: "du", Alts: []string{"zu"}},
	{Kana: "デ", Primary: "de"},
	{Kana: "ド", Primary: "do"},

	{Kana: "バ", Primary: "ba"},
	{Kana: "ビ", Primary: "bi"},
	{Kana: "ブ", Primary: "bu"},
	{Kana: "ベ", Primary: "be"},
	{Kana: "ボ", Primary: "bo"},

	{Kana: "パ", Primary: "pa"},
	{Kana: "ピ", Primary: "pi"},
	{Kana: "プ", Primary: "pu"},
	{Kana: "ペ", Primary: "pe"},
	{Kana: "ポ", Primary: "po"},

	{Kana: "キャ", Primary: "kya"},
	{Kana: "キュ", Primary: "kyu"},
	{Kana: "キョ", Primary: "kyo"},
	{Kana: "ギャ", Primary: "gya"},
	{Kana: "ギュ", Primary: "gyu"},
	{Kana: "ギョ", Primary: "gyo"},

	{Kana: "シャ", Primary: "sha", Alts: []string{"sya"}},
	{Kana: "シュ", Primary: "shu", Alts: []string{"syu"}},
	{Kana: "ショ", Primary: "sho", Alts: []string{"syo"}},
	{Kana: "ジャ", Primary: "ja", Alts: []string{"zya", "jya"}},
	{Kana: "ジュ", Primary: "ju", Alts: []string{"zyu", "jyu"}},
	{Kana: "ジョ", Primary: "jo", Alts: []string{"zyo", "jyo"}},

	{Kana: "チャ", Primary: "cha", Alts: []string{"tya"}},
	{Kana: "チュ", Primary: "chu", Alts: []string{"tyu"}},
	{Kana: "チョ", Primary: "cho", Alts: []string{"tyo"}},

	{Kana: "ニャ", Primary: "nya"},
	{Kana: "ニュ", Primary: "nyu"},
	{Kana: "ニョ", Primary: "nyo"},

	{Kana: "ヒャ", Primary: "hya"},
	{Kana: "ヒュ", Primary: "hyu"},
	{Kana: "ヒョ", Primary: "hyo"},
	{Kana: "ビャ", Primary: "bya"},
	{Kana: "ビュ", Primary: "byu"},
	{Kana: "ビョ", Primary: "byo"},
	{Kana: "ピャ", Primary: "pya"},
	{Kana: "ピュ", Primary: "pyu"},
	{Kana: "ピョ", Primary: "pyo"},

	{Kana: "ミャ", Primary: "mya"},
	{Kana: "ミュ", Primary: "myu"},
	{Kana: "ミョ", Primary: "myo"},

	{Kana: "リャ", Primary: "rya"},
	{Kana: "リュ", Primary: "ryu"},
	{Kana: "リョ", Primary: "ryo"},
}
