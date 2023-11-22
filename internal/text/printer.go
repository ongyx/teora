package text

import (
	"image"
	"image/color"
	"io"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// Printer is a text drawer.
type Printer struct {
	face font.Face
}

// NewPrinter creates a new printer using the OpenType or TrueType font in src.
func NewPrinter(src io.ReaderAt, fo *opentype.FaceOptions) (*Printer, error) {
	f, err := opentype.ParseReaderAt(src)
	if err != nil {
		return nil, err
	}

	fc, err := opentype.NewFace(f, fo)
	if err != nil {
		return nil, err
	}

	return &Printer{face: fc}, nil
}

// Print prints text to the destination image at the given point.
func (p *Printer) Print(dst *ebiten.Image, txt string, pt image.Point, c color.Color) {
	text.Draw(dst, txt, p.face, pt.X, pt.Y, c)
}

// DebugPrint prints text to the upper left corner of the destination image.
func (p *Printer) DebugPrint(dst *ebiten.Image, txt string) {
	h := p.face.Metrics().Height.Ceil()

	p.Print(dst, txt, image.Point{0, h}, color.White)
}

// Bound returns the drawn size of the text in pixels.
func (p *Printer) Bound(txt string) image.Rectangle {
	b, _ := font.BoundString(p.face, txt)

	return image.Rect(
		b.Min.X.Floor(),
		b.Min.Y.Floor(),
		b.Max.X.Ceil(),
		b.Max.Y.Ceil(),
	)
}
