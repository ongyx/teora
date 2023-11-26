package debug

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/teora/internal/text"
	"github.com/ongyx/teora/internal/vec"
)

var (
	backgroundColor = color.RGBA{A: 127}
)

// Overlay is a graphical overlay that shows debug information.
type Overlay struct {
	printer *text.Printer
	point   image.Point

	screenSize      image.Point
	graphicsLibrary string

	text       string
	background image.Rectangle
}

// NewOverlay creates a new overlay, printing debug information using the given printer at the point.
func NewOverlay(p *text.Printer, pt image.Point) *Overlay {
	return &Overlay{printer: p, point: pt}
}

// Update updates the overlay.
func (o *Overlay) Update() {
	if o.graphicsLibrary == "" {
		var di ebiten.DebugInfo
		ebiten.ReadDebugInfo(&di)

		o.graphicsLibrary = di.GraphicsLibrary.String()
	}

	o.text = fmt.Sprintf(
		`Running on %s
Screen size (%d, %d)
%f TPS
%f FPS`,
		o.graphicsLibrary,
		o.screenSize.X,
		o.screenSize.Y,
		ebiten.ActualTPS(),
		ebiten.ActualFPS(),
	)

	// Use the bounding box as the background.
	o.background = o.printer.Measure(o.text).Add(o.point)
}

// Draw draws debug information to the destination image.
func (o *Overlay) Draw(dst *ebiten.Image) {
	o.screenSize = dst.Bounds().Size()

	// Draw the background.
	var v vec.Vec
	v.Rect(o.background)
	v.Draw(dst, &vec.DrawOptions{Color: backgroundColor})

	// Draw the text text.
	o.printer.Print(dst, o.text, o.point, color.White)
}
