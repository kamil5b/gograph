package gograph

func CreateNode(v interface{}) *Node {
	n := new(Node)
	n.value = n
	n.next = nil
	n.firstEdge = nil
	return n
}

func CreateEdge(n *Node, v int) *Edge {
	e := new(Edge)
	e.toNode = n
	e.next = nil
	e.value = v
	return e
}
