package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetFullscreen(true)

	g := &Game{}

	if err := ebiten.RunGame(g); err != nil {
		fmt.Println(err)
	}
}
