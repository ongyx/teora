package fonts

import (
	"embed"
	"io"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/ongyx/teora/internal/text"
	"github.com/ongyx/teora/internal/util"
)

var (
	//go:embed *.otf *.ttf
	embedFS embed.FS

	CommitMono, Teoran *text.Printer
)

func init() {
	o := &opentype.FaceOptions{
		Size:    12,
		DPI:     96,
		Hinting: font.HintingFull,
	}

	CommitMono = util.Must(loadFont("CommitMono-400-Regular.otf", o))

	Teoran = util.Must(loadFont("teoran.ttf", o))
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
