package maze

import "fmt"

type Node struct {
	X, Y int

	Visited   bool
	Neighbors [4]*Node // up, right, down, left
}

func (n *Node) NeighborsCount() (c uint8) {
	for i := 0; i < 4; i++ {
		if n.Neighbors[i] != nil {
			c++
		}
	}
	return
}

func (n *Node) DistanceTo(other *Node) (d int) {
	if n.X == other.X {
		if n.Y > other.Y {
			d = n.Y - other.Y
		} else {
			d = other.Y - n.Y
		}
	} else if n.Y == other.Y {
		if n.X > other.X {
			d = n.X - other.X
		} else {
			d = other.X - n.X
		}
	} else {
		panic("There cannot be a link between Nodes that are not horizontally or vertically aligned")
	}

	return
}

func (n *Node) String() string {
	return fmt.Sprintf("(%d, %d)", n.X, n.Y)
}
