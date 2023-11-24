package main

import (
	"image"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/ongyx/teora/internal/assets/fonts"
	"github.com/ongyx/teora/internal/debug"
	"github.com/ongyx/teora/internal/text"
	"github.com/ongyx/teora/internal/util"
	"github.com/ongyx/teora/internal/vec"
)

var (
	_ ebiten.Game = &Game{}
)

type Game struct {
	box     *text.Box
	overlay *debug.Overlay
}

func NewGame() *Game {
	p := fonts.CommitMono
	t := "Hello World!\nThis should be on the next line."

	return &Game{
		box:     text.NewBox(p, t, time.Millisecond),
		overlay: debug.NewOverlay(p, image.Point{}),
	}
}

func (g *Game) Update() error {
	g.box.Update()
	g.overlay.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	pt := image.Point{Y: screen.Bounds().Dy()}
	g.box.Draw(screen, pt, color.White)

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

	g.overlay.Draw(screen)
}

func (g *Game) Layout(w, h int) (lw, lh int) {
	return util.DeviceScale(w), util.DeviceScale(h)
}
