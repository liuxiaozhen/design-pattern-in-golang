package builder

import (
	"testing"
)

var expectA = "sprite|chicken wing|big Mac"
var expectB = "cola|chips|McChicken"

func Test_builderA(t *testing.T) {
	builder := &ConcreteBuilderA{}
	director := NewDirector(builder)
	director.Construct()
	product := builder.GetProduct()
	if product != expectA {
		t.Fatalf("error||expert=%v||valule=%v", expectA, product)
	}
	//output
	//sprite|chicken wing|big Mac

}

func Test_builderB(t *testing.T) {
	builder := &ConcreteBuilderB{}
	director := NewDirector(builder)
	director.Construct()
	product := builder.GetProduct()
	if product != expectB {
		t.Fatalf("error||expert=%v||valule=%v", expectB, product)
	}
	//output
	//cola|chips|McChicken
}
