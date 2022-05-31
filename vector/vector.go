package vector

type Vector[T any] struct {
	root   []*Element[T]
	length int
}

func New[T any]() *Vector[T] {
	return &Vector[T]{}
}

func (v *Vector[T]) Push(values ...T) {
	for _, value := range values {
		v.root = append(v.root, newElement(value, v.length))
		v.length = len(v.root)
	}
}

type Element[T any] struct {
	value T
	index int
}

func newElement[T any](value T, index int) *Element[T] {
	return &Element[T]{
		value: value,
		index: index,
	}
}
