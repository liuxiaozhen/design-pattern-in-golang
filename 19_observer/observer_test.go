package observer

import "testing"

func Test_observer(t *testing.T) {
	subject := NewConcreteSubject()
	subject.attach(&ObserverA{})
	subject.attach(&ObserverB{})
	// update state
	subject.UpdateState(3)
	// update state
	subject.UpdateState(5)
}
