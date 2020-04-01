package adapter

import (
	"testing"
)

func Test_Adapter(t *testing.T) {
	expect := "Specific Request"
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	res := target.Request()
	if res != expect {
		t.Fatalf("expect: %s, actual: %s", expect, res)
	}
}
