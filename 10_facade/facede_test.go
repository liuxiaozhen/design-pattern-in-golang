package facade

import "testing"

func Test_facde(t *testing.T) {
	pc := NewComputer()
	pc.boot()
}
