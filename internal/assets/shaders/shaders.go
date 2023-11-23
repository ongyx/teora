package shaders

import (
	"embed"
	
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/teora/internal/util"
)

var (
	//go:embed *.kage
	embedFS embed.FS
	
	Gradient *ebiten.Shader
)

func init() {
	Gradient = util.Must(loadShader("gradient.kage"))
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
