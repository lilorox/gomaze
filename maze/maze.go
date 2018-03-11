package maze

import (
	"errors"
	"image"
	_ "image/png"
	"io"
	"log"
)

type Maze struct {
	Start, End *Node
	Size       image.Rectangle

	Nodes []*Node
}

func NewMaze() (m *Maze) {
	m = &Maze{
		Start: &Node{},
		End:   &Node{},
	}

	return
}

func (m *Maze) LoadFromImage(f io.Reader) (err error) {
	im, _, err := image.Decode(f)
	if err != nil {
		return
	}
	m.Size = im.Bounds()

	width := m.Size.Max.X
	height := m.Size.Max.Y
	log.Printf("Width: %d, Height: %d", width, height)

	// Maintains a list of Nodes that can be linked to horizontally and vertically
	horiNodes := make([]*Node, width)
	vertNodes := make([]*Node, width)

	// Find start and end
	m.End.Y = height - 1
	for x := 1; x < width-1; x++ {
		if IsPath(im, x, 0) {
			m.Start.X = x
			vertNodes[x] = m.Start

			m.Nodes = append(m.Nodes, m.Start)
		}
		if IsPath(im, x, height-1) {
			m.End.X = x

			m.Nodes = append(m.Nodes, m.End)
		}
	}
	log.Printf("Start: %s, End: %s", m.Start, m.End)

	// Build maze graph line by line
	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			// If the current tile of the maze is a wall, remove the possible
			// horiNode and vertNode connections and move on
			if !IsPath(im, x, y) {
				vertNodes[x] = nil
				horiNodes[y] = nil
				continue
			}

			// Check if the current point is a node in the graph: if it has
			// only neighbors on up+down or left+right, it is not a node;
			// otherwise it's a corner or a T and is part of the graph
			neighbors, _ := Neighbors(im, x, y)
			if neighbors == N_UP+N_DOWN || neighbors == N_RIGHT+N_LEFT {
				continue
			}

			n := &Node{X: x, Y: y}
			m.Nodes = append(m.Nodes, n)
			log.Printf("New node: %s", n)

			// Add horizontal link to the left to the previous Node on the row
			if horiNodes[y] != nil {
				log.Printf("%s has %s on the left", n, horiNodes[y])
				n.Neighbors[1] = horiNodes[y]
				horiNodes[y].Neighbors[3] = n
			}
			horiNodes[y] = n

			// Add vertical link upward to the previous Node on the column
			if vertNodes[x] != nil {
				log.Printf("%s has %s above", n, vertNodes[x])
				n.Neighbors[2] = vertNodes[x]
				vertNodes[x].Neighbors[0] = n
			}
			vertNodes[x] = n
		}
	}

	// Link the end
	if vertNodes[m.End.X] == nil {
		return errors.New("The end of the maze is not connected to the rest")
	}

	m.End.Neighbors[0] = vertNodes[m.End.X]
	vertNodes[m.End.X].Neighbors[2] = m.End

	return
}
