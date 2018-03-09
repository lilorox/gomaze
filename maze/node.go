package maze

import (
	"log"
)

type Node struct {
	Point

	Links []*NodeLink
}

func NewNode(linkCount int) *Node {
	return &Node{
		Links: make([]*NodeLink, 0, linkCount),
	}
}

type NodeLink struct {
	Nodes    [2]*Node
	Distance int
}

func NewLink(a, b *Node) *NodeLink {
	log.Printf("New link between %s and %s", a, b)

	var d int
	if a.X == b.X {
		if a.X > b.X {
			d = a.X - b.X
		} else {
			d = b.X - a.X
		}
	} else if a.Y == b.Y {
		if a.Y > b.Y {
			d = a.Y - b.Y
		} else {
			d = b.Y - a.Y
		}
	} else {
		panic("There cannot be a link between Nodes that are not horizontally or vertically aligned")
	}

	l := &NodeLink{
		Distance: d,
	}
	l.Nodes[0] = a
	l.Nodes[1] = b

	return l
}
