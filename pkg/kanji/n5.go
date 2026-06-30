package kanji

// N5 contains all 80 JLPT N5 kanji, ordered by frequency.
// Accepted answers for each character: any on'yomi romaji,
// any kun'yomi romaji, or any English meaning (case-insensitive).
var N5 = []Kanji{
	// Group 1 (1–10)
	{
		Char:     "日",
		Onyomi:   []string{"nichi", "jitsu"},
		Kunyomi:  []string{"hi", "bi", "ka"},
		Meanings: []string{"day", "sun", "japan", "counter for days"},
	},
	{
		Char:     "一",
		Onyomi:   []string{"ichi"},
		Kunyomi:  []string{"hito", "hitotsu"},
		Meanings: []string{"one", "1"},
	},
	{
		Char:     "国",
		Onyomi:   []string{"koku"},
		Kunyomi:  []string{"kuni"},
		Meanings: []string{"country", "state", "nation"},
	},
	{
		Char:     "人",
		Onyomi:   []string{"jin", "nin"},
		Kunyomi:  []string{"hito"},
		Meanings: []string{"person", "human"},
	},
	{
		Char:     "年",
		Onyomi:   []string{"nen"},
		Kunyomi:  []string{"toshi"},
		Meanings: []string{"year", "counter for years"},
	},
	{
		Char:     "大",
		Onyomi:   []string{"dai", "tai"},
		Kunyomi:  []string{"oo", "ooki"},
		Meanings: []string{"large", "big"},
	},
	{
		Char:     "十",
		Onyomi:   []string{"juu"},
		Kunyomi:  []string{"tou", "to"},
		Meanings: []string{"ten", "10"},
	},
	{
		Char:     "二",
		Onyomi:   []string{"ni", "ji"},
		Kunyomi:  []string{"futa", "futatsu"},
		Meanings: []string{"two", "2"},
	},
	{
		Char:     "本",
		Onyomi:   []string{"hon"},
		Kunyomi:  []string{"moto"},
		Meanings: []string{"book", "origin", "present", "true"},
	},
	{
		Char:     "中",
		Onyomi:   []string{"chuu"},
		Kunyomi:  []string{"naka", "uchi"},
		Meanings: []string{"middle", "in", "inside", "center"},
	},

	// Group 2 (11–20)
	{
		Char:     "長",
		Onyomi:   []string{"chou"},
		Kunyomi:  []string{"naga", "nagai", "osa"},
		Meanings: []string{"long", "leader", "chief", "boss"},
	},
	{
		Char:     "出",
		Onyomi:   []string{"shutsu", "sui"},
		Kunyomi:  []string{"de", "deru", "da", "dasu", "i", "ideru"},
		Meanings: []string{"exit", "leave", "go out"},
	},
	{
		Char:     "三",
		Onyomi:   []string{"san"},
		Kunyomi:  []string{"mi", "mittsu"},
		Meanings: []string{"three", "3"},
	},
	{
		Char:     "時",
		Onyomi:   []string{"ji"},
		Kunyomi:  []string{"toki", "doki"},
		Meanings: []string{"time", "hour"},
	},
	{
		Char:     "行",
		Onyomi:   []string{"kou", "gyou", "an"},
		Kunyomi:  []string{"i", "iku", "yu", "yuku", "okona", "okonau"},
		Meanings: []string{"going", "journey", "line", "row"},
	},
	{
		Char:     "見",
		Onyomi:   []string{"ken"},
		Kunyomi:  []string{"mi", "miru", "meru"},
		Meanings: []string{"see", "look", "visible"},
	},
	{
		Char:     "月",
		Onyomi:   []string{"getsu", "gatsu"},
		Kunyomi:  []string{"tsuki"},
		Meanings: []string{"month", "moon"},
	},
	{
		Char:     "分",
		Onyomi:   []string{"bun", "fun", "bu"},
		Kunyomi:  []string{"wa", "waru", "wakeru"},
		Meanings: []string{"part", "minute", "understand"},
	},
	{
		Char:     "後",
		Onyomi:   []string{"go", "kou"},
		Kunyomi:  []string{"nochi", "ushi", "ato"},
		Meanings: []string{"behind", "back", "later"},
	},
	{
		Char:     "前",
		Onyomi:   []string{"zen"},
		Kunyomi:  []string{"mae"},
		Meanings: []string{"in front", "before"},
	},

	// Group 3 (21–30)
	{
		Char:     "生",
		Onyomi:   []string{"sei", "shou"},
		Kunyomi:  []string{"i", "iru", "mu", "umareru", "o", "ha", "haeru", "nama"},
		Meanings: []string{"life", "birth", "genuine"},
	},
	{
		Char:     "五",
		Onyomi:   []string{"go"},
		Kunyomi:  []string{"itsu", "itsutsu"},
		Meanings: []string{"five", "5"},
	},
	{
		Char:     "間",
		Onyomi:   []string{"kan", "ken"},
		Kunyomi:  []string{"aida", "ma", "ai"},
		Meanings: []string{"interval", "space"},
	},
	{
		Char:     "上",
		Onyomi:   []string{"jou", "shou"},
		Kunyomi:  []string{"ue", "uwa", "kami", "ageru", "noboru"},
		Meanings: []string{"above", "up"},
	},
	{
		Char:     "東",
		Onyomi:   []string{"tou"},
		Kunyomi:  []string{"higashi"},
		Meanings: []string{"east"},
	},
	{
		Char:     "四",
		Onyomi:   []string{"shi"},
		Kunyomi:  []string{"yo", "yotsu", "yon"},
		Meanings: []string{"four", "4"},
	},
	{
		Char:     "今",
		Onyomi:   []string{"kon", "kin"},
		Kunyomi:  []string{"ima"},
		Meanings: []string{"now", "the present"},
	},
	{
		Char:     "金",
		Onyomi:   []string{"kin", "kon"},
		Kunyomi:  []string{"kane", "kana"},
		Meanings: []string{"gold", "money", "friday"},
	},
	{
		Char:     "九",
		Onyomi:   []string{"kyuu", "ku"},
		Kunyomi:  []string{"kokono", "kokonotsu"},
		Meanings: []string{"nine", "9"},
	},
	{
		Char:     "入",
		Onyomi:   []string{"nyuu"},
		Kunyomi:  []string{"i", "iru", "hairu"},
		Meanings: []string{"enter", "insert"},
	},

	// Group 4 (31–40)
	{
		Char:     "学",
		Onyomi:   []string{"gaku"},
		Kunyomi:  []string{"mana", "manabu"},
		Meanings: []string{"study", "learn", "science"},
	},
	{
		Char:     "高",
		Onyomi:   []string{"kou"},
		Kunyomi:  []string{"taka", "takai"},
		Meanings: []string{"tall", "high", "expensive"},
	},
	{
		Char:     "円",
		Onyomi:   []string{"en"},
		Kunyomi:  []string{"maru", "marui"},
		Meanings: []string{"circle", "yen", "round"},
	},
	{
		Char:     "子",
		Onyomi:   []string{"shi", "su"},
		Kunyomi:  []string{"ko", "ne"},
		Meanings: []string{"child", "kid"},
	},
	{
		Char:     "外",
		Onyomi:   []string{"gai", "ge"},
		Kunyomi:  []string{"soto", "hoka", "hazu", "to"},
		Meanings: []string{"outside"},
	},
	{
		Char:     "八",
		Onyomi:   []string{"hachi"},
		Kunyomi:  []string{"ya", "yattsu", "you"},
		Meanings: []string{"eight", "8"},
	},
	{
		Char:     "六",
		Onyomi:   []string{"roku"},
		Kunyomi:  []string{"mu", "muttsu", "mui"},
		Meanings: []string{"six", "6"},
	},
	{
		Char:     "下",
		Onyomi:   []string{"ka", "ge"},
		Kunyomi:  []string{"shita", "shimo", "moto", "sageru", "kudaru", "orosu"},
		Meanings: []string{"below", "down", "low", "inferior"},
	},
	{
		Char:     "来",
		Onyomi:   []string{"rai"},
		Kunyomi:  []string{"kuru", "kitaru", "ki", "ko"},
		Meanings: []string{"come", "next", "due"},
	},
	{
		Char:     "気",
		Onyomi:   []string{"ki", "ke"},
		Kunyomi:  []string{"iki"},
		Meanings: []string{"spirit", "mind", "air", "mood", "energy"},
	},

	// Group 5 (41–50)
	{
		Char:     "小",
		Onyomi:   []string{"shou"},
		Kunyomi:  []string{"chii", "chiisai", "ko", "o", "sa"},
		Meanings: []string{"little", "small"},
	},
	{
		Char:     "七",
		Onyomi:   []string{"shichi"},
		Kunyomi:  []string{"nana", "nanatsu", "nano"},
		Meanings: []string{"seven", "7"},
	},
	{
		Char:     "山",
		Onyomi:   []string{"san", "sen"},
		Kunyomi:  []string{"yama"},
		Meanings: []string{"mountain"},
	},
	{
		Char:     "話",
		Onyomi:   []string{"wa"},
		Kunyomi:  []string{"hana", "hanasu", "hanashi"},
		Meanings: []string{"talk", "tale", "speech"},
	},
	{
		Char:     "女",
		Onyomi:   []string{"jo", "nyo"},
		Kunyomi:  []string{"onna", "me"},
		Meanings: []string{"woman", "female"},
	},
	{
		Char:     "北",
		Onyomi:   []string{"hoku"},
		Kunyomi:  []string{"kita"},
		Meanings: []string{"north"},
	},
	{
		Char:     "午",
		Onyomi:   []string{"go"},
		Kunyomi:  []string{"uma"},
		Meanings: []string{"noon"},
	},
	{
		Char:     "百",
		Onyomi:   []string{"hyaku", "byaku"},
		Kunyomi:  []string{"momo"},
		Meanings: []string{"hundred", "100"},
	},
	{
		Char:     "書",
		Onyomi:   []string{"sho"},
		Kunyomi:  []string{"kaku"},
		Meanings: []string{"write"},
	},
	{
		Char:     "先",
		Onyomi:   []string{"sen"},
		Kunyomi:  []string{"saki", "mazu"},
		Meanings: []string{"before", "ahead", "previous", "future"},
	},

	// Group 6 (51–60)
	{
		Char:     "名",
		Onyomi:   []string{"mei", "myou"},
		Kunyomi:  []string{"na"},
		Meanings: []string{"name", "noted", "reputation"},
	},
	{
		Char:     "川",
		Onyomi:   []string{"sen"},
		Kunyomi:  []string{"kawa"},
		Meanings: []string{"river", "stream"},
	},
	{
		Char:     "千",
		Onyomi:   []string{"sen"},
		Kunyomi:  []string{"chi"},
		Meanings: []string{"thousand", "1000"},
	},
	{
		Char:     "水",
		Onyomi:   []string{"sui"},
		Kunyomi:  []string{"mizu"},
		Meanings: []string{"water"},
	},
	{
		Char:     "半",
		Onyomi:   []string{"han"},
		Kunyomi:  []string{"naka", "nakaba"},
		Meanings: []string{"half", "middle", "semi"},
	},
	{
		Char:     "男",
		Onyomi:   []string{"dan", "nan"},
		Kunyomi:  []string{"otoko"},
		Meanings: []string{"male", "man"},
	},
	{
		Char:     "西",
		Onyomi:   []string{"sei", "sai"},
		Kunyomi:  []string{"nishi"},
		Meanings: []string{"west"},
	},
	{
		Char:     "電",
		Onyomi:   []string{"den"},
		Kunyomi:  []string{},
		Meanings: []string{"electricity", "electric"},
	},
	{
		Char:     "校",
		Onyomi:   []string{"kou"},
		Kunyomi:  []string{},
		Meanings: []string{"school"},
	},
	{
		Char:     "語",
		Onyomi:   []string{"go"},
		Kunyomi:  []string{"kata", "kataru"},
		Meanings: []string{"language", "word", "speech"},
	},

	// Group 7 (61–70)
	{
		Char:     "土",
		Onyomi:   []string{"do", "to"},
		Kunyomi:  []string{"tsuchi"},
		Meanings: []string{"soil", "earth", "ground"},
	},
	{
		Char:     "木",
		Onyomi:   []string{"boku", "moku"},
		Kunyomi:  []string{"ki", "ko"},
		Meanings: []string{"tree", "wood"},
	},
	{
		Char:     "聞",
		Onyomi:   []string{"bun", "mon"},
		Kunyomi:  []string{"ki", "kiku", "ku", "kuu"},
		Meanings: []string{"hear", "listen", "ask"},
	},
	{
		Char:     "食",
		Onyomi:   []string{"shoku", "jiki"},
		Kunyomi:  []string{"ku", "kuu", "ta", "taberu"},
		Meanings: []string{"eat", "food"},
	},
	{
		Char:     "車",
		Onyomi:   []string{"sha"},
		Kunyomi:  []string{"kuruma"},
		Meanings: []string{"car", "wheel"},
	},
	{
		Char:     "何",
		Onyomi:   []string{"ka"},
		Kunyomi:  []string{"nani", "nan"},
		Meanings: []string{"what"},
	},
	{
		Char:     "南",
		Onyomi:   []string{"nan", "na"},
		Kunyomi:  []string{"minami"},
		Meanings: []string{"south"},
	},
	{
		Char:     "万",
		Onyomi:   []string{"man", "ban"},
		Kunyomi:  []string{},
		Meanings: []string{"ten thousand", "10000"},
	},
	{
		Char:     "毎",
		Onyomi:   []string{"mai"},
		Kunyomi:  []string{"goto", "gotoni"},
		Meanings: []string{"every"},
	},
	{
		Char:     "白",
		Onyomi:   []string{"haku", "byaku"},
		Kunyomi:  []string{"shiro", "shiroi"},
		Meanings: []string{"white"},
	},

	// Group 8 (71–80)
	{
		Char:     "天",
		Onyomi:   []string{"ten"},
		Kunyomi:  []string{"ama", "amatsu"},
		Meanings: []string{"heavens", "sky", "heaven"},
	},
	{
		Char:     "母",
		Onyomi:   []string{"bo"},
		Kunyomi:  []string{"haha", "ka"},
		Meanings: []string{"mother"},
	},
	{
		Char:     "火",
		Onyomi:   []string{"ka"},
		Kunyomi:  []string{"hi", "bi", "ho"},
		Meanings: []string{"fire", "flame"},
	},
	{
		Char:     "右",
		Onyomi:   []string{"u", "yuu"},
		Kunyomi:  []string{"migi"},
		Meanings: []string{"right", "right direction"},
	},
	{
		Char:     "読",
		Onyomi:   []string{"doku", "toku", "tou"},
		Kunyomi:  []string{"yo", "yomu"},
		Meanings: []string{"read"},
	},
	{
		Char:     "友",
		Onyomi:   []string{"yuu"},
		Kunyomi:  []string{"tomo"},
		Meanings: []string{"friend", "buddy"},
	},
	{
		Char:     "左",
		Onyomi:   []string{"sa"},
		Kunyomi:  []string{"hidari"},
		Meanings: []string{"left", "left direction"},
	},
	{
		Char:     "休",
		Onyomi:   []string{"kyuu"},
		Kunyomi:  []string{"yasu", "yasumu"},
		Meanings: []string{"rest", "day off", "sleep"},
	},
	{
		Char:     "父",
		Onyomi:   []string{"fu"},
		Kunyomi:  []string{"chichi", "tou"},
		Meanings: []string{"father", "dad"},
	},
	{
		Char:     "雨",
		Onyomi:   []string{"u"},
		Kunyomi:  []string{"ame", "ama"},
		Meanings: []string{"rain"},
	},
}

// N5Groups divides the N5 kanji into groups of 10 for the level progression system.
var N5Groups = [][]Kanji{
	N5[0:10],
	N5[10:20],
	N5[20:30],
	N5[30:40],
	N5[40:50],
	N5[50:60],
	N5[60:70],
	N5[70:80],
}
