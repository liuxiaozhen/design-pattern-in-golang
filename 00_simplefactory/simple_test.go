package simplefactory

import (
	"testing"
)

// test simpleFactory Rectangle
func Test_simpleFactory1(t *testing.T) {
	//test Rectangle
	d, err := NewShape("Rectangle")
	if err != nil {
		t.Fatal("Rectangle test failed")
	}
	if s := d.draw(); s != "draw a Rectangle" {
		t.Fatal("Rectangle test failed")
	}
}

// test simpleFactory Square
func Test_simpleFactory2(t *testing.T) {
	//test Square
	d, err := NewShape("Square")
	if err != nil {
		t.Fatal("Square test failed")
	}
	if s := d.draw(); s != "draw a Square" {
		t.Fatal("Square test failed")
	}
}

// test simpleFactory Circle
func Test_simpleFactory3(t *testing.T) {
	//test Circle
	d1, err := NewShape("Circle")
	if err != nil {
		t.Fatal("Circle test failed")
	}
	if s := d1.draw(); s != "draw a Circle" {
		t.Fatal("Circle test failed")
	}
}
