package grapes

import (
	"fmt"
	"image"
)

type Palette struct {
	Colors []Color
}

func (p Palette) String() string {
	return fmt.Sprint(p.Colors)
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

func (p *Palette) Analogues() Palette {
	var result []Color
	result = append(result, p.Colors...)

	for _, c := range p.Colors {
		colors := c.Analogue()
		result = append(result,
			colors[0],
			colors[1],
		)
	}

	np := Palette{Colors: result}
	SortByHSL(np.Colors, true, false, false)

	return np
}

func (p *Palette) Shades(n int, strength float64) Palette {
	var result []Color

	for _, c := range p.Colors {
		result = append(result, c.Shades(n, strength)...)
	}

	np := Palette{Colors: result}
	SortByHSL(np.Colors, false, false, true)

	return np
}

func (p *Palette) Complementary() Palette {
	var result []Color
	result = append(result, p.Colors...)

	for _, c := range p.Colors {
		colors := c.Complimentary()
		result = append(result,
			colors,
		)
	}

	np := Palette{Colors: result}
	SortByHSL(np.Colors, false, false, true)

	return np
}

func (p *Palette) Triadic() Palette {
	var result []Color
	result = append(result, p.Colors...)

	for _, c := range p.Colors {
		colors := c.Triade()
		result = append(result,
			colors[0],
			colors[1],
		)
	}
	return Palette{Colors: result}
}

func (p *Palette) Tetradic() Palette {
	var result []Color
	result = append(result, p.Colors...)

	for _, c := range p.Colors {
		colors := c.Tetradic()
		result = append(result,
			colors[0],
			colors[1],
			colors[2],
		)
	}
	return Palette{Colors: result}
}

func (p *Palette) AnalogousAccent() Palette {
	var result []Color
	result = append(result, p.Colors...)

	for _, c := range p.Colors {
		colors := c.AnalogueAccent()
		result = append(result,
			colors[0],
			colors[1],
			colors[2],
		)
	}
	return Palette{Colors: result}
}
