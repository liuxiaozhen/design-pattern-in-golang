package interpreter

import (
	"testing"
)

func TestInterpreter(t *testing.T) {
	p := &Parser{}
	p.Parse("0 + 1 - 2 + 3 - 4 + 5")
	res := p.Result().Interpret()
	expect := 3
	if res != expect {
		t.Fatalf("expect %d got %d", expect, res)
	}
}
