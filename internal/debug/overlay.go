package debug

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/ongyx/teora/internal/text"
	"github.com/ongyx/teora/internal/vec"
)

var (
	overlayBg = color.RGBA{A: 127}
	overlayFg = color.White
	overlayX  = color.RGBA{R: 255, A: 255}
)

// Overlay is a graphical overlay that shows debug information.
type Overlay struct {
	printer *text.Printer
	point   image.Point

	graphicsLibrary string
	screenSize      image.Point
	deviceScale     float64
	cursorPosition  image.Point

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

	if o.deviceScale == 0 {
		o.deviceScale = ebiten.DeviceScaleFactor()
	}

	o.cursorPosition = image.Pt(ebiten.CursorPosition())

	o.text = fmt.Sprintf(
		`Running on %s
Screen (%d, %d)
DScale %0.2f
Cursor (%d, %d)
TPS    %f
FPS    %f`,
		o.graphicsLibrary,
		o.screenSize.X,
		o.screenSize.Y,
		o.deviceScale,
		o.cursorPosition.X,
		o.cursorPosition.Y,
		ebiten.ActualTPS(),
		ebiten.ActualFPS(),
	)

	// Use the bounding box as the background.
	o.background = o.printer.Measure(o.text).Add(o.point)
}

// Draw draws debug information to the destination image.
func (o *Overlay) Draw(dst *ebiten.Image) {
	o.screenSize = dst.Bounds().Size()

	// Draw a crosshair.
	var ch vec.Vec

	// The vertical line of the crosshair.
	vl := image.Point{X: o.cursorPosition.X}
	ch.Move(vl)
	vl.Y = o.screenSize.Y
	ch.Line(vl)

	// The horizontal line of the crosshair.
	hl := image.Point{Y: o.cursorPosition.Y}
	ch.Move(hl)
	hl.X = o.screenSize.X
	ch.Line(hl)

	ch.Draw(dst, &vec.DrawOptions{
		StrokeOp: vector.StrokeOptions{Width: 1},
		Color:    overlayX,
	})

	// Draw the background.
	var bg vec.Vec

	bg.Rect(o.background)
	bg.Draw(dst, &vec.DrawOptions{Color: overlayBg})

	// Draw the text.
	o.printer.Print(dst, o.text, o.point, overlayFg)
}
