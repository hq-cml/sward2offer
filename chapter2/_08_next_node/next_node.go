package _08_next_node

type Node struct {
	I      int
	Left   *Node
	Right  *Node
	Parant *Node
}

func NextNode(find *Node) *Node {
	if find == nil {
		return nil
	}

	p := find
	if p.Right != nil {
		p = p.Right
		for p.Left != nil {
			p = p.Left
		}
		return p
	}
	if p.Parant != nil {
		if p.Parant.Left == p {
			return p.Parant
		} else {
			for p.Parant != nil && p.Parant.Right == p {
				p = p.Parant
			}
			return p.Parant
		}
	}
	return nil
}
