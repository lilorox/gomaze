package maze

type Node struct {
	Point

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

func (n *Node) String() string {
	return n.Point.String()
}
