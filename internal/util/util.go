package util

import "github.com/hajimehoshi/ebiten/v2"

// Must unwraps a two-tuple of (T, error) into T.
// If the error is non-nil, a panic occurs.
func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}

	return t
}

// DeviceScale scales the given int with the device scale factor for high DPI rendering.
func DeviceScale(i int) int {
	return int(float64(i) * ebiten.DeviceScaleFactor())
}
