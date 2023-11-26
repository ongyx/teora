package main

import (
	"image"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/teora/internal/assets/fonts"
	"github.com/ongyx/teora/internal/debug"
	"github.com/ongyx/teora/internal/text"
	"github.com/ongyx/teora/internal/util"
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
	t := `This is some scrolling text
over multiple lines.`

	return &Game{
		box:     text.NewBox(p, t, 50*time.Millisecond),
		overlay: debug.NewOverlay(p, image.Point{}),
	}
}

func (g *Game) Update() error {
	g.box.Update()
	g.overlay.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Gray{63})

	pt := image.Point{X: screen.Bounds().Dx() / 2}

	g.box.Draw(screen, pt, color.White)
	g.overlay.Draw(screen)
}

func (g *Game) Layout(w, h int) (lw, lh int) {
	return util.DeviceScale(w), util.DeviceScale(h)
}
