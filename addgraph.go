package gograph

import "errors"

func (g *Graph) AddFirst(n *Node) {
	n.next = g.firstNode
	g.firstNode = n
}

func (g *Graph) AddLast(n *Node) {
	m := g.firstNode
	for m.next != nil {
		m = m.next
	}

	m.next = n
}

func (g *Graph) Add(v interface{}) {
	n := CreateNode(v)
	g.totalNode++
	if g.firstNode == nil {
		g.firstNode = n
		return
	}
	g.AddLast(n)
}

func (n *Node) Connect(m *Node, value int) {
	e := CreateEdge(m, value)
	e.next = n.firstEdge
	n.firstEdge = e
}

func (g *Graph) ConnectNodes(parent interface{}, child interface{}, value int) error {
	n := g.firstNode
	for n != nil {
		if n.value == parent {
			break
		}
		n = n.next
	}
	if n == nil {
		return errors.New("parent not found")
	}

	m := g.firstNode
	for m != nil {
		if m.value == parent {
			break
		}
		m = m.next
	}
	n.Connect(m, value)

	return nil
}
