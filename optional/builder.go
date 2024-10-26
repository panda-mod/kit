package optional

type Builder[T any] struct {
	target T
}

func New[T any](target T) *Builder[T] {
	return &Builder[T]{
		target: target,
	}
}

func (b *Builder[T]) With(opts ...Option[T]) *Builder[T] {
	for _, fn := range opts {
		fn(&b.target)
	}
	return b
}

func (b *Builder[T]) Value() T {
	return b.target
}

func (b *Builder[T]) Pointer() *T {
	return &b.target
}
