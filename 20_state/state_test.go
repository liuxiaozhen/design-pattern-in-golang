package state

import "testing"

func Test_state(t *testing.T) {
	c := NewContext(NewStateA())
	c.request()
	c.request()
	c.request()
	c.request()
	//output
	//this is state A.
	//this is state B.
	//this is state C.
	//this is state A.
}
