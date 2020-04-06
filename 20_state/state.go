package state

import "fmt"

type Context struct {
	state State
}

func (c *Context) setState(s State) {
	c.state = s
}

func (c *Context) getState() State {
	return c.state
}

func (c *Context) request() {
	c.state.handler(c)
}

func NewContext(s State) *Context {
	return &Context{
		state: s,
	}
}

type State interface {
	handler(c *Context)
}

type StateA struct {
}

func (s *StateA) handler(c *Context) {
	fmt.Println("this is state A.")
	c.setState(NewStateB())
}

func NewStateA() *StateA {
	return &StateA{}
}

type StateB struct {
}

func (s *StateB) handler(c *Context) {
	fmt.Println("this is state B.")
	c.setState(NewStateC())
}

func NewStateB() *StateB {
	return &StateB{}
}

type StateC struct {
}

func (s *StateC) handler(c *Context) {
	fmt.Println("this is state C.")
	c.setState(NewStateA())
}
func NewStateC() *StateC {
	return &StateC{}
}
