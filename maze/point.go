package maze

import (
	"image"
)

const (
	N_UP    = 0x01
	N_RIGHT = 0x02
	N_DOWN  = 0x04
	N_LEFT  = 0x08
)

func IsPath(im image.Image, x, y int) bool {
	r, _, _, _ := im.At(x, y).RGBA()
	return (r > 0)
}

type Point struct {
	image.Point
}

// (0,0) is the up-left corner of the image
func (p Point) Neighbors(im image.Image) (neighbors uint8, count int) {
	if p.Y > 0 && IsPath(im, p.X, p.Y-1) {
		neighbors |= N_UP
		count++
	}
	if p.X < im.Bounds().Max.X-1 && IsPath(im, p.X+1, p.Y) {
		neighbors |= N_RIGHT
		count++
	}
	if p.Y < im.Bounds().Max.Y-1 && IsPath(im, p.X, p.Y+1) {
		neighbors |= N_DOWN
		count++
	}
	if p.X > 0 && IsPath(im, p.X-1, p.Y) {
		neighbors |= N_LEFT
		count++
	}
	return
}

func (p Point) DistanceTo(other Point) (d int) {
	if p.X == other.X {
		if p.Y > other.Y {
			d = p.Y - other.Y
		} else {
			d = other.Y - p.Y
		}
	} else if p.Y == other.Y {
		if p.X > other.X {
			d = p.X - other.X
		} else {
			d = other.X - p.X
		}
	} else {
		panic("There cannot be a link between Nodes that are not horizontally or vertically aligned")
	}

	return
}
