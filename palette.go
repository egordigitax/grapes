package grapes

import (
	"fmt"
	"image"
	"math"
	"sort"
	"strings"
)

type Palette struct {
	Colors []Color
}

func (p Palette) String() string {
	var parts []string
	for _, c := range p.Colors {
		parts = append(parts, fmt.Sprint(c))
	}
	return "colors: " + strings.Join(parts, ", ")
}

func (p Palette) Unpack() []Color {
	return p.Colors
}

func (p Palette) ToFlatRGBFloats() []float64 {
	r := make([]float64, len(p.Colors)*3)
	idx := 0
	for i := 0; i < len(p.Colors); i++ {
		floats := p.Colors[i].ToFloats()
		r[idx], r[idx+1], r[idx+2] = floats[0], floats[1], floats[2]
		idx += 3
	}
	return r
}

func (p Palette) ToFlatRGBAFloats() []float64 {
	r := make([]float64, len(p.Colors)*4)
	idx := 0
	for i := 0; i < len(p.Colors); i++ {
		floats := p.Colors[i].ToFloats()
		r[idx], r[idx+1], r[idx+2], r[idx+3] = floats[0], floats[1], floats[2], floats[3]
		idx += 4
	}
	return r
}

func NewPalette(colors ...Color) Palette {
	return Palette{
		Colors: colors,
	}
}

func NewPaletteFromImage(img image.Image, colorsNum int) Palette {
	return Palette{
		Colors: FromImage(img, colorsNum),
	}
}

func (p *Palette) SortByHSL(sortHue, sortSaturation, sortLightness bool) {
	sort.Slice(p.Colors, func(i, j int) bool {
		h1, s1, l1 := p.Colors[i].ToHSL()
		h2, s2, l2 := p.Colors[j].ToHSL()

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

func (p *Palette) Analogues() Palette {
	var result []Color
	result = append(result, p.Colors...)

	for _, c := range p.Colors {
		h, s, l := c.ToHSL()
		h1 := math.Mod(h+30, 360)
		h2 := math.Mod(h-30+360, 360)
		result = append(result,
			FromHSLf(h1, s, l, float64(c.A)/255),
			FromHSLf(h2, s, l, float64(c.A)/255),
		)
	}

	np := Palette{Colors: result}
	np.SortByHSL(false, false, true)

	return np
}

func (p *Palette) Shades(n int, strength float64) Palette {
	var result []Color

	for _, c := range p.Colors {
		h, s, l0 := c.ToHSL()
		result = append(result, FromHSLf(h, s, clamp(l0), float64(c.A)/255))

		steps := n - 1
		if steps == 0 {
			continue
		}

		for i := 1; i <= steps; i++ {
			shift := strength * float64(i) / float64(steps)
			if i%2 == 1 {
				lighter := clamp(l0 + shift)
				result = append(result, FromHSLf(h, s, lighter, float64(c.A)/255))
			} else {
				darker := clamp(l0 - shift)
				result = append(result, FromHSLf(h, s, darker, float64(c.A)/255))
			}
		}
	}

	np := Palette{Colors: result}
	np.SortByHSL(false, false, true)

	return np
}

func (p *Palette) Complementary() Palette {
	var result []Color
	result = append(result, p.Colors...)

	for _, c := range p.Colors {
		h, s, l := c.ToHSL()
		hComp := math.Mod(h+180, 360)
		result = append(result,
			FromHSLf(hComp, s, l, float64(c.A)/255),
		)
	}
	return Palette{Colors: result}
}

func (p *Palette) Triadic() Palette {
	var result []Color
	result = append(result, p.Colors...)

	for _, c := range p.Colors {
		h, s, l := c.ToHSL()
		h1 := math.Mod(h+120, 360)
		h2 := math.Mod(h+240, 360)
		result = append(result,
			FromHSLf(h1, s, l, float64(c.A)/255),
			FromHSLf(h2, s, l, float64(c.A)/255),
		)
	}
	return Palette{Colors: result}
}

func (p *Palette) Tetradic() Palette {
	var result []Color
	result = append(result, p.Colors...)

	for _, c := range p.Colors {
		h, s, l := c.ToHSL()
		h1 := math.Mod(h+90, 360)
		h2 := math.Mod(h+180, 360)
		h3 := math.Mod(h+270, 360)
		result = append(result,
			FromHSLf(h1, s, l, float64(c.A)/255),
			FromHSLf(h2, s, l, float64(c.A)/255),
			FromHSLf(h3, s, l, float64(c.A)/255),
		)
	}
	return Palette{Colors: result}
}

func (p *Palette) AnalogousAccent() Palette {
	var result []Color
	result = append(result, p.Colors...)

	for _, c := range p.Colors {
		h, s, l := c.ToHSL()
		hAcc := math.Mod(h+180, 360)
		h1 := math.Mod(hAcc+30, 360)
		h2 := math.Mod(hAcc-30+360, 360)
		result = append(result,
			FromHSLf(h1, s, l, float64(c.A)/255),
			FromHSLf(hAcc, s, l, float64(c.A)/255),
			FromHSLf(h2, s, l, float64(c.A)/255),
		)
	}
	return Palette{Colors: result}
}
