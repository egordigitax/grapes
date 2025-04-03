package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"

	"github.com/egordigitax/grapes"
)

func main() {

	file, err := os.Open("/Users/egordigitax/Documents/coin.png")
	if err != nil {
		fmt.Println("err while opening file")
	}

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("err while decoding image")
	}

	color, err := grapes.FromImage(img, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(color)
	palette := grapes.MakePalette(color...)
	newPalette := palette.Shades(5, 0.15)
	fmt.Println(newPalette)

}
