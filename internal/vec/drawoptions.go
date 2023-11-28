package vec

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/vector"
)

// DrawOptions represents options for drawing a vector.
type DrawOptions struct {
	// Options for drawing the vector as a stroke.
	// If not set, the vector drawing is filled.
	StrokeOp vector.StrokeOptions

	// The color that the vector should be drawn in.
	Color color.Color

	// Whether or not the vector drawing should be anti-aliased.
	AntiAlias bool
}

// DefaultDrawOptions returns the default drawing options used in vec.Draw.
func DefaultDrawOptions() *DrawOptions {
	return &DrawOptions{
		Color: color.White,
	}
}
