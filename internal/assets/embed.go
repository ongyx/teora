package assets

import (
	"embed"
)

var (
	//go:embed fonts/*
	//go:embed shaders/*.go
	//go:embed sprites/*.png
	embedFS embed.FS
)
