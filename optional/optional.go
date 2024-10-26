package optional

// Option 通用选项接口
type Option[T any] func(*T)

// Build 通用选项构造器
func Build[T any](target T, opts []Option[T]) *Builder[T] {
	return New(target).With(opts...)
}
