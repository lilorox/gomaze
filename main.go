package main

import (
	"log"
	"os"

	"github.com/lilorox/gomaze/maze"
)

var imageFile = "examples/tiny.png"

func main() {
	f, err := os.Open(imageFile)
	if err != nil {
		log.Fatalf("Cannot open file %s", imageFile)
	}

	m := maze.New()
	err = m.LoadFromImage(f)
	if err != nil {
		log.Fatal("Could not load maze from image")
	}
}
