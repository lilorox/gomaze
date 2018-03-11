package maze

import "image"

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

func Neighbors(im image.Image, x, y int) (neighbors uint8, count int) {
	if y > 0 && IsPath(im, x, y-1) {
		neighbors |= N_UP
		count++
	}
	if x < im.Bounds().Max.X-1 && IsPath(im, x+1, y) {
		neighbors |= N_RIGHT
		count++
	}
	if y < im.Bounds().Max.Y-1 && IsPath(im, x, y+1) {
		neighbors |= N_DOWN
		count++
	}
	if x > 0 && IsPath(im, x-1, y) {
		neighbors |= N_LEFT
		count++
	}
	return
}
