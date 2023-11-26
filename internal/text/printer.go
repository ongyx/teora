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
//
// This implementation differs from ebiten's text package in the way text is drawn.
// Instead of drawing relative to a glyph's origin,
// the upper-left corner of the bounding box is used as the origin (as with images).
//
// For reference, see the following diagram:
// https://developer.apple.com/library/archive/documentation/TextFonts/Conceptual/CocoaTextArchitecture/Art/glyphterms_2x.png
type Printer struct {
	face font.Face

	// Cached font ascent.
	ascent int
}

// NewPrinter creates a new printer with a fontface.
func NewPrinter(face font.Face) *Printer {
	return &Printer{
		face:   face,
		ascent: face.Metrics().Ascent.Ceil(),
	}
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

	return NewPrinter(fc), nil
}

// Face returns the printer's fontface.
func (p *Printer) Face() font.Face {
	return p.face
}

// Print prints text to the destination image at the given point.
func (p *Printer) Print(dst *ebiten.Image, txt string, pt image.Point, c color.Color) {
	// Add ascent to get the origin.
	pt.Y += p.ascent

	text.Draw(dst, txt, p.face, pt.X, pt.Y, c)
}

// DebugPrint prints text to the upper left corner of the destination image.
func (p *Printer) DebugPrint(dst *ebiten.Image, txt string) {
	p.Print(dst, txt, image.Point{}, color.White)
}

// Measure returns the drawn size of the text in pixels.
// This function is similar to MeasureString in the x/image/font package, except that newlines are accounted for.
func (p *Printer) Measure(txt string) image.Rectangle {
	var (
		bounds  fixed.Rectangle26_6
		advance fixed.Point26_6

		prevR rune = -1
	)

	m := p.face.Metrics()
	height := m.Height
	ascent := m.Ascent

	for _, r := range txt {
		if prevR >= 0 {
			// Advance X by the kern between the previous and current rune.
			advance.X += p.face.Kern(prevR, r)
		}

		if r == '\n' {
			// Reset X, advance Y, and process the next rune.
			advance.X = 0
			advance.Y += height
			prevR = -1

			continue
		}

		b, a, _ := p.face.GlyphBounds(r)

		// Extend bounds by the advance.
		bounds = bounds.Union(b.Add(advance))

		// Advance X by the glyph's width.
		advance.X += a

		// Set the previous rune to the current one.
		prevR = r
	}

	// Add ascent to get the bounds relative to the upper-left corner.
	bounds = bounds.Add(fixed.Point26_6{Y: ascent})

	return rtoi(bounds)
}

func rtoi(r fixed.Rectangle26_6) image.Rectangle {
	return image.Rect(
		r.Min.X.Floor(),
		r.Min.Y.Floor(),
		r.Max.X.Ceil(),
		r.Max.Y.Ceil(),
	)
}
