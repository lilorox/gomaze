package maze

import (
	"image"
)

func IsPath(im image.Image, x, y int) bool {
	r, _, _, _ := im.At(x, y).RGBA()
	return (r > 0)
}

type Point struct {
	image.Point
}

// (0,0) is the up-left corner of the image
func (p Point) Neighbors(im image.Image) (up, right, down, left bool, count int) {
	if p.Y > 0 {
		up = IsPath(im, p.X, p.Y-1)
		if up {
			count++
		}
	}
	if p.X < im.Bounds().Max.X-1 {
		right = IsPath(im, p.X+1, p.Y)
		if right {
			count++
		}
	}
	if p.Y < im.Bounds().Max.Y-1 {
		down = IsPath(im, p.X, p.Y+1)
		if down {
			count++
		}
	}
	if p.X > 0 {
		left = IsPath(im, p.X-1, p.Y)
		if left {
			count++
		}
	}
	return
}
