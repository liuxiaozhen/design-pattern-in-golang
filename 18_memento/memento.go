package memento

type Originator struct {
	state int
}

func (o *Originator) createMemento() *Memento {
	return &Memento{
		state: o.state,
	}
}

func (o *Originator) getState() int {
	return o.state
}

func (o *Originator) setState(state int) {
	o.state = state
}
func (o *Originator) setMemento(m *Memento) {
	o.state = m.Load()
}

func NewOriginator(state int) *Originator {
	return &Originator{
		state,
	}
}

type Memento struct {
	state int
}

func (m *Memento) Save(state int) {
	m.state = state
}

func (m *Memento) Load() int {
	return m.state
}

type Caretaker struct {
	m *Memento
}

func (c *Caretaker) getMemento() *Memento {
	return c.m
}

func NewCaretaker(m *Memento) *Caretaker {
	return &Caretaker{
		m,
	}
}
