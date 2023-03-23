package observer

import (
	"fmt"
	"testing"
)

func TestObserverPattern(t *testing.T) {
	s := NewSubject[string]()
	o1 := NewObserver("o1", func(data string) {
		fmt.Printf("%q updated\n", data)
	})
	o2 := NewObserver("o2", func(data string) {
		fmt.Printf("display %q\n", data)
	})
	s.RegisterObserver(o1)
	s.RegisterObserver(o2)
	s.NotifyObservers("testdata1")
}
