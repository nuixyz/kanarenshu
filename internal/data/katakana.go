package data

import "github.com/nuixyz/kanarenshu/pkg/romaji"

var KatakanaGroups = [][]romaji.Character{
	{
		romaji.Katakana[0],
		romaji.Katakana[1],
		romaji.Katakana[2],
		romaji.Katakana[3],
		romaji.Katakana[4],
	},
	{
		romaji.Katakana[5],
		romaji.Katakana[6],
		romaji.Katakana[7],
		romaji.Katakana[8],
		romaji.Katakana[9],
	},
	{
		romaji.Katakana[10],
		romaji.Katakana[11],
		romaji.Katakana[12],
		romaji.Katakana[13],
		romaji.Katakana[14],
	},
	{
		romaji.Katakana[15],
		romaji.Katakana[16],
		romaji.Katakana[17],
		romaji.Katakana[18],
		romaji.Katakana[19],
	},
	{
		romaji.Katakana[20],
		romaji.Katakana[21],
		romaji.Katakana[22],
		romaji.Katakana[23],
		romaji.Katakana[24],
	},
	{
		romaji.Katakana[25],
		romaji.Katakana[26],
		romaji.Katakana[27],
		romaji.Katakana[28],
		romaji.Katakana[29],
	},
	{
		romaji.Katakana[30],
		romaji.Katakana[31],
		romaji.Katakana[32],
		romaji.Katakana[33],
		romaji.Katakana[34],
	},
	{
		romaji.Katakana[35],
		romaji.Katakana[36],
		romaji.Katakana[37],
		romaji.Katakana[38],
		romaji.Katakana[39],
	},
	{
		romaji.Katakana[40],
		romaji.Katakana[41],
		romaji.Katakana[42],
		romaji.Katakana[43],
		romaji.Katakana[44],
	},
	{
		romaji.Katakana[45],
		romaji.Katakana[46],
		romaji.Katakana[47],
		romaji.Katakana[48],
		romaji.Katakana[49],
	},
	{
		romaji.Katakana[50],
		romaji.Katakana[51],
		romaji.Katakana[52],
		romaji.Katakana[53],
		romaji.Katakana[54],
	},
	{
		romaji.Katakana[55],
		romaji.Katakana[56],
		romaji.Katakana[57],
		romaji.Katakana[58],
		romaji.Katakana[59],
	},
	{
		romaji.Katakana[60],
		romaji.Katakana[61],
		romaji.Katakana[62],
		romaji.Katakana[63],
		romaji.Katakana[64],
	},
	{
		romaji.Katakana[65],
		romaji.Katakana[66],
		romaji.Katakana[67],
		romaji.Katakana[68],
		romaji.Katakana[69],
	},
	{
		romaji.Katakana[70],
		romaji.Katakana[71],
		romaji.Katakana[72],
		romaji.Katakana[73],
		romaji.Katakana[74],
	},
	{
		romaji.Katakana[75],
		romaji.Katakana[76],
		romaji.Katakana[77],
		romaji.Katakana[78],
		romaji.Katakana[79],
	},
	{
		romaji.Katakana[80],
		romaji.Katakana[81],
		romaji.Katakana[82],
		romaji.Katakana[83],
		romaji.Katakana[84],
	},
	{
		romaji.Katakana[85],
		romaji.Katakana[86],
		romaji.Katakana[87],
		romaji.Katakana[88],
		romaji.Katakana[89],
	},
	{
		romaji.Katakana[90],
		romaji.Katakana[91],
		romaji.Katakana[92],
		romaji.Katakana[93],
		romaji.Katakana[94],
	},
	{
		romaji.Katakana[95],
		romaji.Katakana[96],
		romaji.Katakana[97],
		romaji.Katakana[98],
		romaji.Katakana[99],
	},
	{
		romaji.Katakana[100],
		romaji.Katakana[101],
		romaji.Katakana[102],
		romaji.Katakana[103],
	},
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
