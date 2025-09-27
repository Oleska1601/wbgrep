package grepper

type queue struct {
	values []ResultLine
	size   int
}

func newQueue(size int) *queue {
	return &queue{
		values: make([]ResultLine, 0, size),
		size:   size,
	}
}

// добавить элемент
func (q *queue) enqueue(item ResultLine) {
	if len(q.values) == q.size {
		q.values = q.values[1:]
	}
	q.values = append(q.values, item)

}

func (q *queue) getAll() []ResultLine {
	return q.values
}

func (q *queue) clear() {
	q.values = make([]ResultLine, 0, q.size)
}
