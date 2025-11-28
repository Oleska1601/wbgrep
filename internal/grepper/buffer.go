package grepper

type buffer struct {
	values []string
	size   int
}

func newBuffer(size int) *buffer {
	return &buffer{
		values: make([]string, 0, size),
		size:   size,
	}
}

// добавить элемент
func (b *buffer) add(item string) {
	if len(b.values) == b.size {
		b.values = b.values[1:]
	}
	b.values = append(b.values, item)

}

func (b *buffer) getAll() []string {
	return b.values
}

func (b *buffer) clear() {
	b.values = make([]string, 0, b.size)
}
