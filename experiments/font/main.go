package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	internal_text "github.com/ongyx/teora/internal/text"
	"github.com/ongyx/teora/internal/vec"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

const (
	greeting = "Hello World!\nHello World!"
)

var (
	face = basicfont.Face7x13
)

type Game struct {
	p *internal_text.Printer
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	c := screen.Bounds().Size().Div(2)

	var v vec.Vec
	m := g.p.Measure(greeting).Add(c)
	v.Rect(m)
	v.Draw(screen, &vec.DrawOptions{Color: color.Gray{127}})

	g.p.Print(screen, greeting, c, color.White)
}

func (g *Game) Layout(w, h int) (lw, lh int) {
	return w, h
}

func main() {
	mt := face.Metrics()
	fmt.Printf("height: %d, ascent: %d, descent: %d\n", mt.Height.Ceil(), mt.Ascent.Ceil(), mt.Descent.Ceil())

	p := internal_text.NewPrinter(face)
	m := p.Measure(greeting)
	fmt.Printf("Printer.Measure: (%d, %d), (%d, %d)\n", m.Min.X, m.Min.Y, m.Max.X, m.Max.Y)

	b1, _ := font.BoundString(face, greeting)
	fmt.Printf("font.boundString: (%d, %d), (%d, %d)\n", b1.Min.X.Ceil(), b1.Min.Y.Ceil(), b1.Max.X.Ceil(), b1.Max.Y.Ceil())

	b2 := text.BoundString(face, greeting)
	fmt.Printf("text.boundString: (%d, %d), (%d, %d)\n", b2.Min.X, b2.Min.Y, b2.Max.X, b2.Max.Y)

	if err := ebiten.RunGame(&Game{p}); err != nil {
		panic(err)
	}
}
