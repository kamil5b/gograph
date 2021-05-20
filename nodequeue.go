package gograph

type QueueNode struct {
	toNode *Node
	next   *QueueNode
}

type Queue struct {
	head, pointer, tail *QueueNode
}

func createQueue() Queue {
	var q Queue
	q.head = nil
	q.pointer = nil
	q.tail = nil
	return q
}

func createQueueNode(n *Node) *QueueNode {
	q := new(QueueNode)
	q.toNode = n
	q.next = nil
	return q
}

func (q *Queue) enqueue(n *Node) {
	nq := createQueueNode(n)
	if q.tail == nil {
		q.tail = nq
		q.head = nq
		q.pointer = nq
	} else {
		q.tail.next = nq
		q.tail = nq
	}
}

func (q *Queue) inQueue(n *Node) bool {
	for qn := q.head; qn != nil; qn = qn.next {
		if qn.toNode == n {
			return true
		}
	}

	return false
}
