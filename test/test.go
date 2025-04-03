package main

import (
	"fmt"

	"github.com/egordigitax/grapes"
)

func main() {
	color := grapes.FromFloats(0.2, 0.5, 0.6, 1.0)
	fmt.Println("color.Shades: ", color.Shades(5, 0.2))
	pal := grapes.NewPalette(color)
	analogues := pal.Shades(3, 0.2).Unpack()
	fmt.Println("palette.Shades: ", analogues)
	comp := color.Complimentary()
	newPal := grapes.NewPalette(color)
	palComp := newPal.Complementary()
	fmt.Println("color.Complimentary: ", comp)
	fmt.Println("palette.Complimentary: ", palComp)
	fmt.Println("Complementary floats: ", palComp.Complementary().ToFlatRGBFloats())
	fmt.Println("palette.Triadic: ", newPal.Triadic().ToFlatRGBAFloats())
	fmt.Println("color.AnalogueAccent: ", color.AnalogueAccent())
}
