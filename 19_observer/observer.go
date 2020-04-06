package observer

import "fmt"

type Observer interface {
	Update(state int)
}
type Subject interface {
	attach(Observer)
	notify()
}

type ConcreteSubject struct {
	state     int
	observers []Observer
}

func (s *ConcreteSubject) attach(ob Observer) {
	s.observers = append(s.observers, ob)
}

func (s *ConcreteSubject) notify() {
	for _, o := range s.observers {
		o.Update(s.state)
	}
}

func (s *ConcreteSubject) UpdateState(state int) {
	s.state = state
	s.notify()
}

func NewConcreteSubject() *ConcreteSubject {
	return &ConcreteSubject{
		observers: make([]Observer, 0),
	}
}

type ObserverA struct {
}

func (*ObserverA) Update(state int) {
	fmt.Printf("ObserverA receive subject state=%v\n", state)
}

type ObserverB struct {
}

func (*ObserverB) Update(state int) {
	fmt.Printf("ObserverB receive subject state= %v\n", state)
}
