package maze

import (
	"fmt"
	"io"
	"log"
)

type DotGraph struct {
	Maze     *Maze
	NodeGrid [][]*Node
	Edges    [][2]*Node
}

func NewDotGraph(m *Maze) (d *DotGraph) {
	log.Printf("Building dot graph")
	d = &DotGraph{
		Maze:     m,
		NodeGrid: make([][]*Node, m.Size.Max.X),
	}

	for i := 0; i < m.Size.Max.X; i++ {
		d.NodeGrid[i] = make([]*Node, m.Size.Max.Y)
	}

	for i := range m.Nodes {
		n := m.Nodes[i]
		d.NodeGrid[n.X][n.Y] = n
		// Only add edges from right and down neighbors to avoid duplicates
		for _, j := range []int{1, 2} {
			if n.Neighbors[j] != nil {
				d.Edges = append(d.Edges, [2]*Node{n, n.Neighbors[j]})
			}
		}
	}
	log.Printf("Dot graph ready with %d nodes and %d edges", len(m.Nodes), len(d.Edges))

	return
}

func (d *DotGraph) Save(f io.Writer) (err error) {
	dot := `digraph G {
	center=1
	rank=same
	rankdir=LR
	ration=auto
	splines=line
	edge [dir=none]
`

	// Add node in clusters
	for i := range d.NodeGrid[0] {
		subgraph := fmt.Sprintf("\tsubgraph cluster_%d {\n\t\tstyle=invis\n", i)
		for j := range d.NodeGrid {
			n := d.NodeGrid[j][i]
			if n != nil {
				subgraph += fmt.Sprintf(
					"\t\t\"%s\" [pos=\"%d,%d!\"];\n",
					n,
					n.X,
					d.Maze.Size.Max.Y-n.Y,
				)
			}
		}
		dot += subgraph + "\t}\n"
	}

	// Add links globally
	for i := range d.Edges {
		dot += fmt.Sprintf(
			"\t\"%s\" -> \"%s\" [label=\"%.1f\"]\n",
			d.Edges[i][0],
			d.Edges[i][1],
			d.Edges[i][0].DistanceTo(d.Edges[i][1]),
		)
	}

	dot += "}"
	_, err = f.Write([]byte(dot))

	return
}
