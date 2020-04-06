package decorator

import "testing"

func Test_decorator(t *testing.T) {
	c := NewConcreteComponent()
	d := NewConcreteDecorator(c)
	d.Operation()
}
