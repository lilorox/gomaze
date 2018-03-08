package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
)

var imageFile = "examples/small.png"

func main() {
	fmt.Println("gomaze")

	f, err := os.Open(imageFile)
	if err != nil {
		log.Fatalf("Cannot open file %s", imageFile)
	}

	im, format, err := image.Decode(f)
	if err != nil {
		log.Fatalf("Cannot decode image %s", imageFile)
	}
	log.Printf("Image: %s, Format: %s", imageFile, format)

	width := im.Bounds().Max.X
	height := im.Bounds().Max.Y
	log.Printf("Width: %d, Height: %d", width, height)

	// Find start and end
	start := image.Point{X: 0, Y: 0}
	end := image.Point{X: 0, Y: height - 1}

	for x := 1; x < height-1; x++ {
		r, _, _, _ := im.At(x, 0).RGBA()
		if r > 0 {
			start.X = x
		}

		r, _, _, _ = im.At(x, height-1).RGBA()
		if r > 0 {
			end.X = x
		}
	}

	log.Printf("Start: %s, End: %s", start, end)

	// Build maze graph

}
