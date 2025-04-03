package grapes

import (
	"fmt"
	"image"
	"math"
	"strings"
)

type Color struct {
	R, G, B, A uint8
}

func (c Color) Hex() string {
	return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
}

func (c Color) String() string {
	return c.Hex()
}

func (c Color) ToFloats() [4]float64 {
	return [4]float64{
		float64(c.R) / 255,
		float64(c.G) / 255,
		float64(c.B) / 255,
		float64(c.A) / 255,
	}
}

func (c Color) Analogue() [2]Color {
	h, s, l := c.ToHSL()
	h1 := math.Mod(h+30, 360)
	h2 := math.Mod(h-30+360, 360)
	return [2]Color{
		FromHSLf(h1, s, l, float64(c.A)/255),
		FromHSLf(h2, s, l, float64(c.A)/255),
	}
}

func (c Color) Complimentary() Color {
	h, s, l := c.ToHSL()
	hComp := math.Mod(h+180, 360)
	return FromHSLf(hComp, s, l, float64(c.A)/255)
}

func (c Color) Triade() [2]Color {
	h, s, l := c.ToHSL()
	h1 := math.Mod(h+120, 360)
	h2 := math.Mod(h+240, 360)
	return [2]Color{
		FromHSLf(h1, s, l, float64(c.A)/255),
		FromHSLf(h2, s, l, float64(c.A)/255),
	}
}

func (c Color) Tetradic() [3]Color {
	h, s, l := c.ToHSL()
	h1 := math.Mod(h+90, 360)
	h2 := math.Mod(h+180, 360)
	h3 := math.Mod(h+270, 360)
	return [3]Color{
		FromHSLf(h1, s, l, float64(c.A)/255),
		FromHSLf(h2, s, l, float64(c.A)/255),
		FromHSLf(h3, s, l, float64(c.A)/255),
	}
}

func (c Color) AnalogueAccent() [3]Color {
	h, s, l := c.ToHSL()
	hAcc := math.Mod(h+180, 360)
	h1 := math.Mod(hAcc+30, 360)
	h2 := math.Mod(hAcc-30+360, 360)
	return [3]Color{
		FromHSLf(h1, s, l, float64(c.A)/255),
		FromHSLf(hAcc, s, l, float64(c.A)/255),
		FromHSLf(h2, s, l, float64(c.A)/255),
	}
}

func (c Color) Shades(num int, strength float64) []Color {
	if num <= 0 {
		return nil
	}

	h, s, l0 := c.ToHSL()
	a := float64(c.A) / 255

	result := make([]Color, 0, num)
	result = append(result, FromHSLf(h, s, clamp(l0), a))

	if num == 1 {
		return result
	}

	steps := num - 1
	for i := 1; i <= steps; i++ {
		t := float64(i) / float64(steps)
		offset := strength * t

		var l float64
		if i%2 == 1 {
			l = clamp(l0 + offset) // сначала светлее
		} else {
			l = clamp(l0 - offset) // потом темнее
		}

		result = append(result, FromHSLf(h, s, l, a))
	}

	SortByHSL(result, false, false, true)

	return result
}

func (c Color) ToHSL() (float64, float64, float64) {
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

func FromHex(hex string) Color {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) == 6 {
		hex += "FF"
	}
	if len(hex) != 8 {
		return Color{}
	}
	return Color{
		R: parseHexByte(hex[0:2]),
		G: parseHexByte(hex[2:4]),
		B: parseHexByte(hex[4:6]),
		A: parseHexByte(hex[6:8]),
	}
}

func FromFloats(floatR, floatG, floatB, floatA float64) Color {
	return Color{
		R: uint8(clamp(floatR) * 255),
		G: uint8(clamp(floatG) * 255),
		B: uint8(clamp(floatB) * 255),
		A: uint8(clamp(floatA) * 255),
	}
}

func FromHSLf(h, s, l, a float64) Color {
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

func ColorDistance(c1, c2 Color) float64 {
	r1, g1, b1 := float64(c1.R), float64(c1.G), float64(c1.B)
	r2, g2, b2 := float64(c2.R), float64(c2.G), float64(c2.B)
	rMean := (r1 + r2) / 2
	return math.Sqrt(
		(2+(rMean/256))*(r1-r2)*(r1-r2) +
			4*(g1-g2)*(g1-g2) +
			(2+((255-rMean)/256))*(b1-b2)*(b1-b2),
	)
}

func FromImage(img image.Image, numColors int) []Color {
	freq := countColors(img)
	sorted := sortColorFrequencies(freq)
	return filterDistinctColors(sorted, numColors)
}
