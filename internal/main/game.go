package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/ongyx/teora/internal/assets/fonts"
	"github.com/ongyx/teora/internal/util"
	"github.com/ongyx/teora/internal/vec"
)

var (
	_ ebiten.Game = &Game{}
)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	fonts.CommitMono.DebugPrint(screen, "Hello World!\nThis should be on the next line.")

	s := screen.Bounds().Size()
	c := s.Div(2)

	var v vec.Vec

	v.Circle(c, float32(c.Y))

	do := &vec.DrawOptions{
		Stroke: vector.StrokeOptions{
			Width: 10,
		},
		Color:     color.White,
		AntiAlias: true,
	}
	v.Draw(screen, do)
}

func (g *Game) Layout(w, h int) (lw, lh int) {
	return util.DeviceScale(w), util.DeviceScale(h)
}
