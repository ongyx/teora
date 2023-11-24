package text

import (
	"testing"

	"golang.org/x/image/font/basicfont"

	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	greeting = "Hello World!"
)

var (
	printer = &Printer{Face: basicfont.Face7x13}
)

func TestPrinterMeasure(t *testing.T) {
	r1 := text.BoundString(printer.Face, greeting)
	r2 := printer.Measure(greeting)

	if r1 != r2 {
		t.Errorf("measure results incorrect: expected %#v, got %#v", r1, r2)
	}
}
