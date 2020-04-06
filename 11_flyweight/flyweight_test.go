package flyweight

import (
	"fmt"
	"testing"
)

func Test_flyweight(t *testing.T) {
	fcFactory := NewFileFlyweightFactory()
	f := fcFactory.Get("abc reading")
	fmt.Println(f.Data())

	f2 := fcFactory.Get("need equal")
	f3 := fcFactory.Get("need equal")
	if f2 != f3 {
		t.Fatal("flyweight test error")
	}
}
