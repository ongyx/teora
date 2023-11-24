package debug

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/teora/internal/text"
	"github.com/ongyx/teora/internal/vec"
)

const (
	overlayTemplate = `%f TPS
%f FPS`
)

var (
	backgroundColor = color.RGBA{R: 127, G: 127, B: 127, A: 127}
)

// Overlay is a graphical overlay that shows debug information.
type Overlay struct {
	printer *text.Printer
	point   image.Point

	info   string
	bounds image.Rectangle
}

// NewOverlay creates a new overlay, printing debug information using the given printer at the point.
func NewOverlay(p *text.Printer, pt image.Point) *Overlay {
	// Add height to compensate for the origin.
	m := p.Face.Metrics()
	pt.Y += m.Height.Ceil()

	return &Overlay{printer: p, point: pt}
}

// Update updates the overlay.
func (o *Overlay) Update() {
	o.info = fmt.Sprintf(overlayTemplate, ebiten.ActualTPS(), ebiten.ActualFPS())

	// Use the bounding box of the info text to draw a background.
	o.bounds = o.printer.Measure(o.info)
}

// Draw draws debug information to the destination image.
func (o *Overlay) Draw(dst *ebiten.Image) {
	// Draw the background.
	var v vec.Vec
	v.Rect(o.bounds.Add(o.point))
	v.Draw(dst, &vec.DrawOptions{Color: backgroundColor})

	// Draw the info text.
	o.printer.Print(dst, o.info, o.point, color.White)
}
