package gograph

import "errors"

type Node struct {
	value     interface{}
	next      *Node
	firstEdge *Edge
}

type Edge struct {
	next   *Edge
	toNode *Node
	value  int
}

type Graph struct {
	firstNode *Node
	totalNode int
}

func (g *Graph) CreateGraph(v interface{}) {
	n := CreateNode(v)
	g.firstNode = n
	g.totalNode = 1
}

func (g *Graph) AddNewChild(parent interface{}, child interface{}, value int) error {

	g.Add(child)
	m := g.firstNode
	for m.next != nil {
		m = m.next
	}

	return g.ConnectNodes(parent, child, value)
}

func (g *Graph) RemoveNode(v interface{}) error {
	if g.totalNode == 0 {
		return errors.New("empty graph")
	}

	g.totalNode--

	for n := g.firstNode; n != nil; n = n.next {
		if n.value == v {
			g.removeOnlyNode(n)
			g.RemoveEdge(n)
			break
		}
	}

	//DELETE UNACCESSABLE NODE
	g.RemoveUnaccessable()

	return nil
}
