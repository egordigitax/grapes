package grapes

import "sort"

func SortByHSL(colors []Color, sortHue, sortSaturation, sortLightness bool) {
	sort.Slice(colors, func(i, j int) bool {
		h1, s1, l1 := colors[i].ToHSL()
		h2, s2, l2 := colors[j].ToHSL()

		if sortHue && h1 != h2 {
			return h1 < h2
		}
		if sortSaturation && s1 != s2 {
			return s1 < s2
		}
		if sortLightness && l1 != l2 {
			return l1 < l2
		}
		return false
	})
}

func sortColorFrequencies(counts map[Color]int) []colorFreq {
	var freq []colorFreq
	for c, count := range counts {
		freq = append(freq, colorFreq{c, count})
	}
	sort.Slice(freq, func(i, j int) bool {
		return freq[i].count > freq[j].count
	})
	return freq
}

type colorFreq struct {
	c     Color
	count int
}
