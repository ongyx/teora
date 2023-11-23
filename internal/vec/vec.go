package vec

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	// Vector graphics require a source to render correctly.
	source = ebiten.NewImage(3, 3)

	sourcePixel = source.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

func init() {
	source.Fill(color.White)
}

// Vec is a helper for vector graphics.
type Vec struct {
	Path vector.Path
}

// Move moves the vector to a new position.
// This does not draw anything.
func (v *Vec) Move(to image.Point) {
	v.Path.MoveTo(float32(to.X), float32(to.Y))
}

// Line draws a line from the current position to another position.
func (v *Vec) Line(to image.Point) {
	v.Path.LineTo(float32(to.X), float32(to.Y))
}

// Rect draws a rectangle with bounds.
func (v *Vec) Rect(bounds image.Rectangle) {
	d := image.Pt(0, bounds.Dy())

	p1 := bounds.Min
	p2 := bounds.Min.Add(d)
	p3 := bounds.Max
	p4 := bounds.Max.Sub(d)

	v.Move(p1)
	v.Line(p2)
	v.Line(p3)
	v.Line(p4)
}

// Arc draws a clockwise arc with a center and radius.
// from and to are angles in radians.
func (v *Vec) Arc(center image.Point, radius, from, to float32) {
	cx := float32(center.X)
	cy := float32(center.Y)

	v.Path.Arc(cx, cy, radius, from, to, vector.Clockwise)
}

// Circle draws a circle with a center and radius.
func (v *Vec) Circle(center image.Point, radius float32) {
	v.Arc(center, radius, 0, 2*math.Pi)
}

// Close closes the current vector path and creates a new one.
// This can be used to begin a separate drawing.
func (v *Vec) Close() {
	v.Path.Close()
}

// Draw draws the vector to the destination image.
func (v *Vec) Draw(dst *ebiten.Image, o *DrawOptions) {
	op := &ebiten.DrawTrianglesOptions{
		// color.Color.RGBA() returns pre-multiplied alpha values,
		// so use the correct scale mode.
		ColorScaleMode: ebiten.ColorScaleModePremultipliedAlpha,
		AntiAlias:      o.AntiAlias,
	}

	var vs []ebiten.Vertex
	var is []uint16

	// Draw a stroke if the appropriate options are set.
	if o.Stroke != (vector.StrokeOptions{}) {
		vs, is = v.Path.AppendVerticesAndIndicesForStroke(nil, nil, &o.Stroke)
	} else {
		vs, is = v.Path.AppendVerticesAndIndicesForFilling(nil, nil)

		// Using EvenOdd ensures complex polygons render properly.
		op.FillRule = ebiten.EvenOdd
	}

	scale(vs, o.Color)

	dst.DrawTriangles(vs, is, sourcePixel, op)
}

func scale(vs []ebiten.Vertex, c color.Color) {
	ir, ig, ib, ia := c.RGBA()

	r := float32(ir) / 0xffff
	g := float32(ig) / 0xffff
	b := float32(ib) / 0xffff
	a := float32(ia) / 0xffff

	for i := range vs {
		v := &vs[i]
		v.SrcX = 1
		v.SrcY = 1
		v.ColorR = r
		v.ColorG = g
		v.ColorB = b
		v.ColorA = a
	}
}
