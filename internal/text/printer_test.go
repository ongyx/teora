package text

import (
	"image"
	"testing"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

const (
	greeting = "Hello World!"
)

var (
	printer = NewPrinter(basicfont.Face7x13)
)

func TestPrinterMeasure(t *testing.T) {
	m := printer.Face().Metrics()
	a := image.Point{Y: m.Ascent.Ceil()}

	b, _ := font.BoundString(printer.Face(), greeting)
	r1 := rtoi(b)

	// Subtract ascent to get the measurement relative to the origin.
	r2 := printer.Measure(greeting).Sub(a)

	if r1 != r2 {
		t.Errorf("measure results incorrect: expected %#v, got %#v", r1, r2)
	}
}
