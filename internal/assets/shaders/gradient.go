package main

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	pos := position.xy / imageDstTextureSize()
	origin, size := imageDstRegionOnTexture()
	pos -= origin
	pos /= size
	return vec4(pos.x, pos.y, 0, 1)
}
