package grapes

import (
	"image"
	"strconv"
)


func parseHexByte(s string) uint8 {
	i, _ := strconv.ParseUint(s, 16, 8)
	return uint8(i)
}


func countColors(img image.Image) map[Color]int {
	bounds := img.Bounds()
	counts := make(map[Color]int)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			if a == 0 {
				continue
			}
			c := Color{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
			counts[c]++
		}
	}
	return counts
}


func filterDistinctColors(sortedColors []colorFreq, limit int) []Color {
	var distinctColors []Color
	if len(sortedColors) == 0 {
		return distinctColors
	}
	distinctColors = append(distinctColors, sortedColors[0].c)
	for _, item := range sortedColors[1:] {
		if len(distinctColors) >= limit {
			break
		}
		if isColorDistinct(item.c, distinctColors) {
			distinctColors = append(distinctColors, item.c)
		}
	}
	return distinctColors
}

func isColorDistinct(c Color, existing []Color) bool {
	for _, e := range existing {
		if ColorDistance(c, e) < 100 {
			return false
		}
	}
	return true
}

func clamp(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}
