package main

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/ongyx/teora/internal/vec"
)

const (
	screenWidth  = 256
	screenHeight = 256

	radius = 8
	stroke = 8
)

var (
	// Reference: https://en.wikipedia.org/wiki/Rainbow#Number_of_colours_in_a_spectrum_or_a_rainbow
	colors = hexToRGB(
		0xFF0000,
		0xFF8000,
		0xFFFF00,
		0x00FF00,
		0x00FFFF,
		0x0000FF,
		0x8000FF,
	)
)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	s := screen.Bounds().Size()
	c := s.Div(2)
	cx := float64(c.X)
	cy := float64(c.Y)

	var cv vec.Vec
	cv.Circle(c, float32(cy))
	cv.Draw(screen, &vec.DrawOptions{Color: color.Gray{127}, AntiAlias: true})

	var rv vec.Vec
	hl := (math.Sqrt(2) * cy) / 2
	r := image.Rect(
		int(cx-hl),
		int(cy-hl),
		int(cx+hl),
		int(cy+hl),
	)
	rv.Rect(r)
	rv.Draw(screen, &vec.DrawOptions{Color: color.White, AntiAlias: true})

	n := len(colors)

	for i, cl := range colors {
		var v vec.Vec

		r := (n - i) * radius
		v.Arc(c, float32(r), math.Pi, 0)

		o := &vec.DrawOptions{
			Stroke: vector.StrokeOptions{
				Width:    stroke,
				LineCap:  vector.LineCapRound,
				LineJoin: vector.LineJoinRound,
			},
			Color:     cl,
			AntiAlias: true,
		}
		v.Draw(screen, o)
	}
}

func (g *Game) Layout(w, h int) (lw, lh int) {
	return screenWidth, screenHeight
}

func main() {
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

func hexToRGB(hexColors ...int) (colors []color.Color) {
	colors = make([]color.Color, len(hexColors))

	for i, h := range hexColors {
		colors[i] = color.RGBA{
			R: uint8((h >> 16) & 0xFF),
			G: uint8((h >> 8) & 0xFF),
			B: uint8(h & 0xFF),
			A: 0xFF,
		}
	}

	return
}
