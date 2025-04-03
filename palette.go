package grapes

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type palette struct {
	colors []color
}

func (p palette) String() string {
	var parts []string
	for _, c := range p.colors {
		parts = append(parts, fmt.Sprint(c))
	}
	return "colors: " + strings.Join(parts, ", ")
}

func MakePalette(colors ...color) palette {
	return palette{
		colors: colors,
	}
}

func (p *palette) SortByHSL(sortHue, sortSaturation, sortLightness bool) {
	sort.Slice(p.colors, func(i, j int) bool {
		h1, s1, l1 := p.colors[i].ToHSL()
		h2, s2, l2 := p.colors[j].ToHSL()

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

func (p *palette) Analogues() palette {
	var result []color
	result = append(result, p.colors...)

	for _, c := range p.colors {
		h, s, l := c.ToHSL()
		h1 := math.Mod(h+30, 360)
		h2 := math.Mod(h-30+360, 360)
		result = append(result,
			FromHSLf(h1, s, l, float64(c.A)/255),
			FromHSLf(h2, s, l, float64(c.A)/255),
		)
	}

	np := palette{colors: result}
	np.SortByHSL(false, false, true)

	return np
}

func (p *palette) Shades(n int, strength float64) palette {
	var result []color

	for _, c := range p.colors {
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
	
	np := palette{colors: result}
	np.SortByHSL(false, false, true)

	return np
}

func (p *palette) Complementary() palette {
	var result []color
	result = append(result, p.colors...)

	for _, c := range p.colors {
		h, s, l := c.ToHSL()
		hComp := math.Mod(h+180, 360)
		result = append(result,
			FromHSLf(hComp, s, l, float64(c.A)/255),
		)
	}
	return palette{colors: result}
}

func (p *palette) Triadic() palette {
	var result []color
	result = append(result, p.colors...)

	for _, c := range p.colors {
		h, s, l := c.ToHSL()
		h1 := math.Mod(h+120, 360)
		h2 := math.Mod(h+240, 360)
		result = append(result,
			FromHSLf(h1, s, l, float64(c.A)/255),
			FromHSLf(h2, s, l, float64(c.A)/255),
		)
	}
	return palette{colors: result}
}

func (p *palette) Tetradic() palette {
	var result []color
	result = append(result, p.colors...)

	for _, c := range p.colors {
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
	return palette{colors: result}
}

func (p *palette) AnalogousAccent() palette {
	var result []color
	result = append(result, p.colors...)

	for _, c := range p.colors {
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
	return palette{colors: result}
}
