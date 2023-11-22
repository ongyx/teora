package assets

import (
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	Demo, Arrow *ebiten.Image
)

func init() {
	Demo = must(loadSprite("sprites/demo.png"))
	Arrow = must(loadSprite("sprites/arrow.png"))
}

func loadSprite(path string) (*ebiten.Image, error) {
	f, err := embedFS.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if img, _, err := image.Decode(f); err != nil {
		return nil, err
	} else {
		return ebiten.NewImageFromImage(img), nil
	}
}
