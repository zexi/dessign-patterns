package observer

type Observer[T any] interface {
	GetId() string
	OnChange(T)
}

type observerImpl[T any] struct {
	id       string
	onChange func(data T)
}

func NewObserver[T any](id string, onChange func(data T)) Observer[T] {
	return &observerImpl[T]{
		id:       id,
		onChange: onChange,
	}
}

func (o observerImpl[T]) GetId() string {
	return o.id
}

func (o observerImpl[T]) OnChange(t T) {
	o.onChange(t)
}
