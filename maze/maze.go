package maze

import (
	"errors"
	//"fmt"
	"image"
	_ "image/png"
	"io"
	"log"
)

/*
func printNodeGrid(ng [][]*Node) {
	for i := range ng[0] {
		line := ""
		for j := range ng {
			if ng[j][i] == nil {
				line += "."
			} else {
				line += "x"
			}
		}
		fmt.Println(line)
	}
}
*/

type Maze struct {
	Start, End *Node
	Size       image.Rectangle

	//NodeGrid [][]*Node
	Links []*NodeLink
}

func New() *Maze {
	return &Maze{
		Start: NewNode(1),
		End:   NewNode(1),
	}
}

func (m *Maze) LoadFromImage(f io.Reader) (err error) {
	im, format, err := image.Decode(f)
	if err != nil {
		return
	}
	log.Printf("Image format: %s", format)
	m.Size = im.Bounds()

	width := m.Size.Max.X
	height := m.Size.Max.Y
	log.Printf("Width: %d, Height: %d", width, height)

	// Initialize the node grid
	/*
		m.NodeGrid = make([][]*Node, width)
		for i := 0; i < width; i++ {
			m.NodeGrid[i] = make([]*Node, height)
		}
	*/
	// Maintains a list of Nodes that can be used for vertical links
	vertNodes := make([]*Node, width)

	// Find start and end
	m.End.Y = height - 1
	for x := 1; x < width-1; x++ {
		if IsPath(im, x, 0) {
			m.Start.X = x
			vertNodes[x] = m.Start
			//m.NodeGrid[x][0] = m.Start
		}
		if IsPath(im, x, height-1) {
			m.End.X = x
			//m.NodeGrid[x][height-1] = m.End
		}
	}
	log.Printf("Start: %s, End: %s", m.Start, m.End)

	// Build maze graph line by line
	p := Point{}
	for p.Y = 1; p.Y < height-1; p.Y++ {
		var prevNode *Node // previous node created on this row

		for p.X = 1; p.X < width-1; p.X++ {
			// If the current tile of the maze is a wall, remove the possible
			// vertNode connection and move on
			if !IsPath(im, p.X, p.Y) {
				vertNodes[p.X] = nil
				continue
			}

			// Check if the current point is a node in the graph
			u, r, d, l, count := p.Neighbors(im)
			//log.Printf("%s: %t %t %t %t", p, u, r, d, l)
			if (!u && !d && r && l) || (u && d && !r && !l) {
				//log.Printf("Not graph node: %s", p)
				continue
			}

			log.Printf("Graph node: %s", p)
			n := NewNode(count)
			n.Point = p
			//m.NodeGrid[p.X][p.Y] = n

			// Add horizontal link to the left to the previous Node on the row
			if prevNode != nil {
				l := NewLink(prevNode, n)
				prevNode.Links = append(prevNode.Links, l)
				n.Links = append(n.Links, l)
			}
			prevNode = n

			// Add vertical link upward to the previous Node on the column
			if vertNodes[p.X] != nil {
				l := NewLink(vertNodes[p.X], n)
				vertNodes[p.X].Links = append(vertNodes[p.X].Links, l)
				n.Links = append(n.Links, l)
			}
			vertNodes[p.X] = n
		}
	}

	// Link the end
	if vertNodes[m.End.X] == nil {
		return errors.New("The end of the maze is not connected to the rest")
	}
	l := NewLink(vertNodes[m.End.X], m.End)
	vertNodes[m.End.X].Links = append(vertNodes[m.End.X].Links, l)
	m.End.Links = append(m.End.Links, l)

	return
}

func (m *Maze) ToDotFile(f io.Writer) (err error) {
	//f.Write([]byte{"graph G {\n"})

	//f.Write([]byte{"}"})
	return
}
