package text

import (
	"image"
	"image/color"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Box is a scrolling text box.
type Box struct {
	printer *Printer

	mu       sync.Mutex
	text     string
	idx, len int

	ticker *time.Ticker
}

// NewBox creates a new text box with the given printer.
// txt is the initial text to display.
// d indicates the duration to wait before displaying the next character.
func NewBox(p *Printer, txt string, d time.Duration) *Box {
	b := &Box{
		printer: p,
		ticker:  time.NewTicker(d),
	}
	b.SetText(txt)

	return b
}

// SetText sets the new text to display.
func (b *Box) SetText(txt string) {
	b.text = txt
	b.idx = 0
	b.len = len(txt)
}

// SetTick sets the new duration to wait before displaying the next character.
func (b *Box) SetTick(d time.Duration) {
	b.ticker.Reset(d)
}

// Update updates the text box's state.
func (b *Box) Update() {
	select {
	case <-b.ticker.C:
		// Show one more character per tick.
		if b.idx < b.len {
			b.idx += 1
		}
	default:
	}
}

// Draw draws the text box to the destination image at the given point.
//
// NOTE: The point represents the bottom-left point of the text box, not the top-left as with a normal image.
func (b *Box) Draw(dst *ebiten.Image, pt image.Point, c color.Color) {
	b.printer.Print(dst, b.text[:b.idx], pt, c)
}

// Close releases resources used by the text box.
func (b *Box) Close() {
	b.ticker.Stop()
}
