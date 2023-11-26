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

	info       info
	text       string
	background image.Rectangle
}

// NewOverlay creates a new overlay, printing debug information using the given printer at the point.
func NewOverlay(p *text.Printer, pt image.Point) *Overlay {
	return &Overlay{printer: p, point: pt}
}

// Update updates the overlay.
func (o *Overlay) Update() {
	o.text = o.info.String()

	// Use the bounding box as the background.
	o.background = o.printer.Measure(o.text).Add(o.point)
}

// Draw draws debug information to the destination image.
func (o *Overlay) Draw(dst *ebiten.Image) {
	// Draw the background.
	var v vec.Vec
	v.Rect(o.background)
	v.Draw(dst, &vec.DrawOptions{Color: backgroundColor})

	// Draw the text text.
	o.printer.Print(dst, o.text, o.point, color.White)
}

type info struct {
	glib     string
	tps, fps float64
}

func (i info) String() string {
	if i.glib == "" {
		var di ebiten.DebugInfo
		ebiten.ReadDebugInfo(&di)

		i.glib = di.GraphicsLibrary.String()
	}

	i.fps = ebiten.ActualFPS()
	i.tps = ebiten.ActualTPS()

	return fmt.Sprintf(`Running on %s
%f TPS
%f FPS`, i.glib, i.tps, i.fps)
}
