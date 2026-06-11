package data

import "github.com/nuixyz/kanarenshu/pkg/romaji"

var HiraganaGroups = [][]romaji.Character{
	{
		romaji.Hiragana[0],
		romaji.Hiragana[1],
		romaji.Hiragana[2],
		romaji.Hiragana[3],
		romaji.Hiragana[4],
	},
	{
		romaji.Hiragana[5],
		romaji.Hiragana[6],
		romaji.Hiragana[7],
		romaji.Hiragana[8],
		romaji.Hiragana[9],
	},
	{
		romaji.Hiragana[10],
		romaji.Hiragana[11],
		romaji.Hiragana[12],
		romaji.Hiragana[13],
		romaji.Hiragana[14],
	},
	{
		romaji.Hiragana[15],
		romaji.Hiragana[16],
		romaji.Hiragana[17],
		romaji.Hiragana[18],
		romaji.Hiragana[19],
	},
	{
		romaji.Hiragana[20],
		romaji.Hiragana[21],
		romaji.Hiragana[22],
		romaji.Hiragana[23],
		romaji.Hiragana[24],
	},
	{
		romaji.Hiragana[25],
		romaji.Hiragana[26],
		romaji.Hiragana[27],
		romaji.Hiragana[28],
		romaji.Hiragana[29],
	},
	{
		romaji.Hiragana[30],
		romaji.Hiragana[31],
		romaji.Hiragana[32],
		romaji.Hiragana[33],
		romaji.Hiragana[34],
	},
	{
		romaji.Hiragana[35],
		romaji.Hiragana[36],
		romaji.Hiragana[37],
		romaji.Hiragana[38],
		romaji.Hiragana[39],
	},
	{
		romaji.Hiragana[40],
		romaji.Hiragana[41],
		romaji.Hiragana[42],
		romaji.Hiragana[43],
		romaji.Hiragana[44],
	},
	{
		romaji.Hiragana[45],
		romaji.Hiragana[46],
		romaji.Hiragana[47],
		romaji.Hiragana[48],
		romaji.Hiragana[49],
	},
	{
		romaji.Hiragana[50],
		romaji.Hiragana[51],
		romaji.Hiragana[52],
		romaji.Hiragana[53],
		romaji.Hiragana[54],
	},
	{
		romaji.Hiragana[55],
		romaji.Hiragana[56],
		romaji.Hiragana[57],
		romaji.Hiragana[58],
		romaji.Hiragana[59],
	},
	{
		romaji.Hiragana[60],
		romaji.Hiragana[61],
		romaji.Hiragana[62],
		romaji.Hiragana[63],
		romaji.Hiragana[64],
	},
	{
		romaji.Hiragana[65],
		romaji.Hiragana[66],
		romaji.Hiragana[67],
		romaji.Hiragana[68],
		romaji.Hiragana[69],
	},
	{
		romaji.Hiragana[70],
		romaji.Hiragana[71],
		romaji.Hiragana[72],
		romaji.Hiragana[73],
		romaji.Hiragana[74],
	},
	{
		romaji.Hiragana[75],
		romaji.Hiragana[76],
		romaji.Hiragana[77],
		romaji.Hiragana[78],
		romaji.Hiragana[79],
	},
	{
		romaji.Hiragana[80],
		romaji.Hiragana[81],
		romaji.Hiragana[82],
		romaji.Hiragana[83],
		romaji.Hiragana[84],
	},
	{
		romaji.Hiragana[85],
		romaji.Hiragana[86],
		romaji.Hiragana[87],
		romaji.Hiragana[88],
		romaji.Hiragana[89],
	},
	{
		romaji.Hiragana[90],
		romaji.Hiragana[91],
		romaji.Hiragana[92],
		romaji.Hiragana[93],
		romaji.Hiragana[94],
	},
	{
		romaji.Hiragana[95],
		romaji.Hiragana[96],
		romaji.Hiragana[97],
		romaji.Hiragana[98],
		romaji.Hiragana[99],
	},
	{
		romaji.Hiragana[100],
		romaji.Hiragana[101],
		romaji.Hiragana[102],
		romaji.Hiragana[103],
	},
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
