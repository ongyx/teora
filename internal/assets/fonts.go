package assets

import (
	"io"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/teora/internal/text"
)

var (
	CommitMono *text.Printer

	Teoran *text.Printer
)

func init() {
	o := &opentype.FaceOptions{
		Size:    12,
		DPI:     96 * ebiten.DeviceScaleFactor(),
		Hinting: font.HintingFull,
	}

	CommitMono = must(loadFont("fonts/CommitMono-400-Regular.otf", o))

	Teoran = must(loadFont("fonts/teoran.ttf", o))
}

func loadFont(name string, fo *opentype.FaceOptions) (*text.Printer, error) {
	f, err := embedFS.Open(name)
	if err != nil {
		return nil, err
	}

	// This should never panic; embed.FS files implement io.ReaderAt
	r := f.(io.ReaderAt)

	return text.NewPrinter(r, fo)
}
