package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/teora/internal/assets/fonts"
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
	fonts.CommitMono.DebugPrint(screen, "Hello World!")

	s := screen.Bounds().Size()
	c := s.Div(2)

	var v vec.Vec
	v.Circle(c, float32(c.Y))
	v.Draw(screen, &vec.DrawOptions{Color: color.White, AntiAlias: true})
}

func (g *Game) Layout(w, h int) (lw, lh int) {
	return w, h
}
