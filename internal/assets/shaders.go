package assets

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	Gradient *ebiten.Shader
)

func init() {
	if g, err := loadShader("shaders/gradient.go"); err != nil {
		panic(err)
	} else {
		Gradient = g
	}
}

func loadShader(path string) (*ebiten.Shader, error) {
	c, err := embedFS.ReadFile(path)
	if err != nil {
		return nil, err
	}

	s, err := ebiten.NewShader(c)
	if err != nil {
		return nil, err
	}

	return s, nil
}
