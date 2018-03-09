package maze

import (
	"log"
)

type Node struct {
	Point

	Name  string
	Links []*NodeLink
}

func NewNode(linkCount int) *Node {
	return &Node{
		Links: make([]*NodeLink, 0, linkCount),
	}
}

func (n *Node) String() string {
	if n.Name != "" {
		return n.Name
	}
	return n.Point.String()
}

type NodeLink struct {
	Nodes    []*Node
	Distance int
}

func NewLink(a, b *Node) *NodeLink {
	var d int
	if a.X == b.X {
		if a.Y > b.Y {
			d = a.Y - b.Y
		} else {
			d = b.Y - a.Y
		}
	} else if a.Y == b.Y {
		if a.X > b.X {
			d = a.X - b.X
		} else {
			d = b.X - a.X
		}
	} else {
		panic("There cannot be a link between Nodes that are not horizontally or vertically aligned")
	}
	log.Printf("New link between %s and %s, distance %d", a, b, d)

	l := &NodeLink{
		Distance: d,
		Nodes:    []*Node{a, b},
	}

	return l
}
