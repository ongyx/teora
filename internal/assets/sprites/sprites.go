package sprites

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/teora/internal/util"
)

var (
	//go:embed *.png
	embedFS embed.FS
	
	Demo, Arrow *ebiten.Image
)

func init() {
	Demo = util.Must(loadSprite("sprites/demo.png"))
	Arrow = util.Must(loadSprite("sprites/arrow.png"))
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
