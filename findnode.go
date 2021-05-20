package gograph

func (g *Graph) BFS(v interface{}) *Node {
	q := createQueue()
	q.enqueue(g.firstNode)
	for q.pointer != nil {
		n := q.pointer.toNode
		if n.value == v {
			return n
		}
		e := n.firstEdge
		for e != nil {
			if !q.inQueue(e.toNode) {
				q.enqueue(e.toNode)
			}
			e = e.next
		}
		q.pointer = q.pointer.next
	}
	return nil
}

func (g *Graph) DFS(v interface{}) *Node {
	q := createQueue()
	return dfs(g.firstNode, &q, v)
}

func dfs(n *Node, q *Queue, v interface{}) *Node {

	e := n.firstEdge
	if q.head == nil {
		q.enqueue(n)
	}
	if n.value == v {
		return n
	}

	for e != nil {
		if !q.inQueue(e.toNode) {
			q.enqueue(e.toNode)
			m := dfs(e.toNode, q, v)
			if m != nil {
				return m
			}
		}
		e = e.next
	}

	return nil
}
