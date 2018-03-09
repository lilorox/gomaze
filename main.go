package main

import (
	"flag"
	"log"
	"os"

	"github.com/lilorox/gomaze/maze"
)

func main() {
	dotVar := flag.String("dot", "", "Dot file for graphviz.")
	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatal("No maze image specified on input")
	}
	imageFile := flag.Arg(0)

	f, err := os.Open(imageFile)
	if err != nil {
		log.Fatalf("Cannot open file %s", imageFile)
	}

	m := maze.NewMaze()
	err = m.LoadFromImage(f)
	if err != nil {
		log.Fatal("Could not load maze from image")
	}

	if dotVar != nil {
		dot, err := os.OpenFile(*dotVar, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatalf("Cannot open dot file %s with write access", *dotVar)
		}

		if err := m.ToDotFile(dot); err != nil {
			log.Fatal(err.Error())
		}
		f.Close()
	}
}
