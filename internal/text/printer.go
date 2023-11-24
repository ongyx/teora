package text

import (
	"image"
	"image/color"
	"io"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// Printer is a text drawer.
type Printer struct {
	Face font.Face
}

// NewPrinterFromReader creates a new printer using the OpenType or TrueType font in src.
func NewPrinterFromReader(src io.ReaderAt, fo *opentype.FaceOptions) (*Printer, error) {
	f, err := opentype.ParseReaderAt(src)
	if err != nil {
		return nil, err
	}

	fc, err := opentype.NewFace(f, fo)
	if err != nil {
		return nil, err
	}

	return &Printer{Face: fc}, nil
}

// Print prints text to the destination image at the given point.
func (p *Printer) Print(dst *ebiten.Image, txt string, pt image.Point, c color.Color) {
	text.Draw(dst, txt, p.Face, pt.X, pt.Y, c)
}

// DebugPrint prints text to the upper left corner of the destination image.
func (p *Printer) DebugPrint(dst *ebiten.Image, txt string) {
	h := p.Face.Metrics().Height

	p.Print(dst, txt, image.Point{Y: h.Ceil()}, color.White)
}

// Measure returns the drawn size of the text in pixels.
// This function is similar to MeasureString in the x/image/font package, except that newlines are accounted for.
func (p *Printer) Measure(txt string) image.Rectangle {
	var (
		bounds  fixed.Rectangle26_6
		advance fixed.Point26_6

		prevR rune = -1
	)

	height := p.Face.Metrics().Height

	for _, r := range txt {
		if prevR >= 0 {
			// Advance X by the kern between the previous and current rune.
			advance.X += p.Face.Kern(prevR, r)
		}

		if r == '\n' {
			// Reset X, advance Y, and process the next rune.
			advance.X = 0
			advance.Y += height
			prevR = -1

			continue
		}

		b, a, _ := p.Face.GlyphBounds(r)

		// Extend bounds by the advance.
		bounds = bounds.Union(b.Add(advance))

		// Advance X by the glyph's width.
		advance.X += a

		// Set the previous rune to the current one.
		prevR = r
	}

	return image.Rect(
		bounds.Min.X.Floor(),
		bounds.Min.Y.Floor(),
		bounds.Max.X.Ceil(),
		bounds.Max.Y.Ceil(),
	)
}
