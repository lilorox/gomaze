package maze

import (
	"errors"
	"fmt"
	"image"
	_ "image/png"
	"io"
	"log"
)

type Maze struct {
	Start, End *Node
	Size       image.Rectangle

	NodeGrid [][]*Node
	Links    []*NodeLink
}

func NewMaze() *Maze {
	m := &Maze{
		Start: NewNode(1),
		End:   NewNode(1),
	}
	m.Start.Name = "start"
	m.End.Name = "end"

	return m
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
	m.NodeGrid = make([][]*Node, width)
	for i := 0; i < width; i++ {
		m.NodeGrid[i] = make([]*Node, height)
	}

	// Maintains a list of Nodes that can be linked to horizontally and vertically
	horiNodes := make([]*Node, width)
	vertNodes := make([]*Node, width)

	// Find start and end
	m.End.Y = height - 1
	for x := 1; x < width-1; x++ {
		if IsPath(im, x, 0) {
			m.Start.X = x
			vertNodes[x] = m.Start
			m.NodeGrid[x][0] = m.Start
		}
		if IsPath(im, x, height-1) {
			m.End.X = x
			m.NodeGrid[x][height-1] = m.End
		}
	}
	log.Printf("Start: %s, End: %s", m.Start, m.End)

	// Build maze graph line by line
	p := Point{}
	for p.Y = 1; p.Y < height-1; p.Y++ {
		for p.X = 1; p.X < width-1; p.X++ {
			// If the current tile of the maze is a wall, remove the possible
			// horiNode and vertNode connections and move on
			if !IsPath(im, p.X, p.Y) {
				vertNodes[p.X] = nil
				horiNodes[p.Y] = nil
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
			m.NodeGrid[p.X][p.Y] = n

			// Add horizontal link to the left to the previous Node on the row
			if horiNodes[p.Y] != nil {
				l := NewLink(horiNodes[p.Y], n)
				horiNodes[p.Y].Links = append(horiNodes[p.Y].Links, l)
				n.Links = append(n.Links, l)
				m.Links = append(m.Links, l)
			}
			horiNodes[p.Y] = n

			// Add vertical link upward to the previous Node on the column
			if vertNodes[p.X] != nil {
				l := NewLink(vertNodes[p.X], n)
				vertNodes[p.X].Links = append(vertNodes[p.X].Links, l)
				n.Links = append(n.Links, l)
				m.Links = append(m.Links, l)
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
	m.Links = append(m.Links, l)

	return
}

func (m *Maze) ToDotFile(f io.Writer) (err error) {
	dot := `digraph G {
	center=1
	rank=same
	rankdir=LR
	ration=auto
	splines=line
	edge [dir=none]
`

	// Add node in clusters
	for i := range m.NodeGrid[0] {
		subgraph := fmt.Sprintf("\tsubgraph cluster_%d {\n\t\tstyle=invis\n", i)
		for j := range m.NodeGrid {
			n := m.NodeGrid[j][i]
			if n != nil {
				subgraph += fmt.Sprintf(
					"\t\t\"%s\" [pos=\"%d,%d!\"];\n",
					n,
					n.Point.X,
					m.Size.Max.Y-n.Point.Y,
				)
			}
		}
		dot += subgraph + "\t}\n"
	}

	// Add links globally
	for i := range m.Links {
		dot += fmt.Sprintf(
			"\t\"%s\" -> \"%s\" [label=\"%d\"]\n",
			m.Links[i].Nodes[0],
			m.Links[i].Nodes[1],
			m.Links[i].Distance,
		)
	}

	dot += "}"
	_, err = f.Write([]byte(dot))
	return
}
