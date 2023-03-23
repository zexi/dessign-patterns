package observer

import "fmt"

type Subject[T any] interface {
	RegisterObserver(ob Observer[T])
	RemoveObserver(ob Observer[T])
	NotifyObservers(data T)
}

type subjectImpl[T any] struct {
	obs map[string]Observer[T]
}

func NewSubject[T any]() Subject[T] {
	return &subjectImpl[T]{
		obs: make(map[string]Observer[T], 0),
	}
}
func (s *subjectImpl[T]) RegisterObserver(ob Observer[T]) {
	s.obs[ob.GetId()] = ob
}

func (s *subjectImpl[T]) RemoveObserver(ob Observer[T]) {
	delete(s.obs, ob.GetId())
}

func (s *subjectImpl[T]) NotifyObservers(data T) {
	for id, ob := range s.obs {
		fmt.Printf("notify observer %q\n", id)
		ob.OnChange(data)
	}
}
