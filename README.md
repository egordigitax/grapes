# 🍇 Grapes

**Grapes** is a Go package for working with colors and generating rich color palettes using HSL transformations, hex parsing, and image-based extraction. It's built to be simple, efficient, and designer-friendly.

## ✨ Features

- Define and convert colors (Hex, RGBA, HSL)
- Generate:
  - Analogous
  - Complementary
  - Triadic
  - Tetradic
  - Accent palettes
  - Shades
- Extract distinct color palettes from images
- Sort and flatten color data for further use (e.g., shaders, graphics)

---

## 📦 Installation

```bash
go get github.com/egordigitax/grapes
```

---

## 📘 Quick Start

```go
package main

import (
	"fmt"
	"github.com/yourusername/grapes"
)

func main() {
	c := grapes.FromHex("#FF5733")
	fmt.Println("Base color:", c)

	analogues := c.Analogue()
	fmt.Println("Analogues:", analogues)

	shades := c.Shades(5, 0.1)
	fmt.Println("Shades:")
	for _, s := range shades {
		fmt.Println(s)
	}
}
```

---

## 🎨 Color API

### Create a Color

```go
c := grapes.FromHex("#AABBCC")
c := grapes.FromFloats(1, 0.5, 0, 1)
```

### Convert Color

```go
hex := c.Hex()
floats := c.ToFloats()
h, s, l := c.ToHSL()
```

### Color Relationships

```go
c.Complimentary()
c.Analogue()
c.Triade()
c.Tetradic()
c.AnalogueAccent()
```

### Shades

```go
c.Shades(5, 0.2) // 5 variants, strength 0.2
```

---

## 🌈 Palette API

### Create a Palette

```go
p := grapes.NewPalette(c1, c2, c3)
```

### From Image

```go
img, _ := os.Open("image.png")
defer img.Close()
decoded, _ := png.Decode(img)

palette := grapes.NewPaletteFromImage(decoded, 6)
```

### Palette Transforms

```go
analog := palette.Analogues()
shades := palette.Shades(4, 0.15)
comp := palette.Complementary()
triadic := palette.Triadic()
tetradic := palette.Tetradic()
accent := palette.AnalogousAccent()
```

### Flatten Palette to RGBA floats (useful for OpenGL, shaders, etc.)

```go
floats := palette.ToFlatRGBAFloats() // [r0, g0, b0, a0, r1, g1, ...]
```

---

## 🧠 Color Distance

```go
dist := grapes.ColorDistance(c1, c2) // perceptual difference metric
```

---

## 🖼️ Image Extraction Example

```go
f, _ := os.Open("some-image.jpg")
defer f.Close()
img, _, _ := image.Decode(f)

colors := grapes.FromImage(img, 8)
for _, c := range colors {
	fmt.Println(c.Hex())
}
```

---

## 📚 Example Use Cases

- UI color theme generation
- Dynamic visualizations
- Procedural art and shaders
- Game asset coloring
- Extracting branding palettes from logos

---

## 🔧 Dev Notes

- All colors use RGBA 8-bit internally
- Palette generation is HSL-based for perceptual consistency
- Sorting is used to stabilize palette outputs

---

## 📄 License

MIT — use freely, but attribution appreciated.
