package grapes

import (
	"errors"
	"fmt"
	"image"
	"math"
	"sort"
	"strconv"
	"strings"
)

type color struct {
	R, G, B, A uint8
}

func (c color) Hex() string {
	return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
}

func (c color) String() string {
	return c.Hex()
}

func (c color) ToHSL() (float64, float64, float64) {
	r := float64(c.R) / 255
	g := float64(c.G) / 255
	b := float64(c.B) / 255

	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)
	d := max - min

	l := (max + min) / 2

	var h, s float64
	if d != 0 {
		if l < 0.5 {
			s = d / (max + min)
		} else {
			s = d / (2.0 - max - min)
		}

		switch max {
		case r:
			h = (g - b) / d
			if g < b {
				h += 6
			}
		case g:
			h = (b-r)/d + 2
		case b:
			h = (r-g)/d + 4
		}
		h *= 60
	}

	return h, s, l
}

func FromHex(hex string) color {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) == 6 {
		hex += "FF"
	}
	if len(hex) != 8 {
		return color{}
	}
	return color{
		R: parseHexByte(hex[0:2]),
		G: parseHexByte(hex[2:4]),
		B: parseHexByte(hex[4:6]),
		A: parseHexByte(hex[6:8]),
	}
}

func parseHexByte(s string) uint8 {
	i, _ := strconv.ParseUint(s, 16, 8)
	return uint8(i)
}

func FromFloats(floatR, floatG, floatB, floatA float64) color {
	return color{
		R: uint8(clamp(floatR) * 255),
		G: uint8(clamp(floatG) * 255),
		B: uint8(clamp(floatB) * 255),
		A: uint8(clamp(floatA) * 255),
	}
}

func FromHSLf(h, s, l, a float64) color {
	h = math.Mod(h, 360)
	c := (1 - math.Abs(2*l-1)) * s
	x := c * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := l - c/2

	var rf, gf, bf float64
	switch {
	case h < 60:
		rf, gf, bf = c, x, 0
	case h < 120:
		rf, gf, bf = x, c, 0
	case h < 180:
		rf, gf, bf = 0, c, x
	case h < 240:
		rf, gf, bf = 0, x, c
	case h < 300:
		rf, gf, bf = x, 0, c
	default:
		rf, gf, bf = c, 0, x
	}

	return FromFloats(rf+m, gf+m, bf+m, a)
}

func ColorDistance(c1, c2 color) float64 {
	r1, g1, b1 := float64(c1.R), float64(c1.G), float64(c1.B)
	r2, g2, b2 := float64(c2.R), float64(c2.G), float64(c2.B)
	rMean := (r1 + r2) / 2
	return math.Sqrt(
		(2+(rMean/256))*(r1-r2)*(r1-r2) +
			4*(g1-g2)*(g1-g2) +
			(2+((255-rMean)/256))*(b1-b2)*(b1-b2),
	)
}

func FromImage(img image.Image, numColors int) ([]color, error) {
	if numColors <= 0 {
		return nil, errors.New("numColors must be > 0")
	}

	freq := countColors(img)
	sorted := sortColorFrequencies(freq)
	return filterDistinctColors(sorted, numColors), nil
}

func countColors(img image.Image) map[color]int {
	bounds := img.Bounds()
	counts := make(map[color]int)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			if a == 0 {
				continue
			}
			c := color{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
			counts[c]++
		}
	}
	return counts
}

func sortColorFrequencies(counts map[color]int) []colorFreq {
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
	c     color
	count int
}

func filterDistinctColors(sortedColors []colorFreq, limit int) []color {
	var distinctColors []color
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

func isColorDistinct(c color, existing []color) bool {
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
