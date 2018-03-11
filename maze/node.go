package maze

import (
	"fmt"
	"math"
)

type Node struct {
	X, Y      int
	Neighbors [4]*Node // up, right, down, left
}

func (n *Node) DistanceTo(other *Node) float64 {
	dX := float64(other.X - n.X)
	if dX == 0 {
		if n.Y > other.Y {
			return float64(n.Y - other.Y)
		}
		return float64(other.Y - n.Y)
	}

	dY := float64(other.Y - n.Y)
	if dY == 0 {
		if n.X > other.X {
			return float64(n.X - other.X)
		}
		return float64(other.X - n.X)
	}

	return math.Sqrt(dX*dX + dY*dY)
}

func (n *Node) String() string {
	return fmt.Sprintf("(%d, %d)", n.X, n.Y)
}
