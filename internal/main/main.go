package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"

	_ "github.com/silbinarywolf/preferdiscretegpu"
)

func main() {
	ebiten.SetFullscreen(true)

	g := NewGame()

	if err := ebiten.RunGame(g); err != nil {
		fmt.Println(err)
	}
}
