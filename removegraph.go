package gograph

func (g *Graph) removeOnlyNode(n *Node) {
	g.totalNode--
	if n == g.firstNode {
		g.firstNode = n.next
	} else {
		m := g.firstNode
		for m.next != n {
			m = m.next
		}
		m.next = n.next
	}
}

func (g *Graph) RemoveUnaccessable() {

	for n := g.firstNode; n != nil; n = n.next {
		if g.BFS(n.value) == nil {
			g.removeOnlyNode(n)
		}
	}
}

func (g *Graph) RemoveEdge(n *Node) {
	m := g.firstNode
	for m != nil {
		e := m.firstEdge
		for e != nil {
			if e.toNode == n {
				if e == m.firstEdge {
					m.firstEdge = m.firstEdge.next
				} else {
					f := m.firstEdge
					for f.next != e {
						f = f.next
					}
					f.next = e.next
				}
			}
			e = e.next
		}
		m = m.next
	}
}
