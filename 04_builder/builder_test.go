package builder

import (
	"testing"
)

var expertA = "sprite|chicken wing|big Mac"
var expertB = "cola|chips|McChicken"

func Test_builderA(t *testing.T) {
	builder := &ConcreteBuilderA{}
	director := NewDirector(builder)
	director.Construct()
	product := builder.GetProduct()
	if product != expertA {
		t.Fatalf("error||expert=%v||valule=%v", expertA, product)
	}
	//output
	//sprite|chicken wing|big Mac

}

func Test_builderB(t *testing.T) {
	builder := &ConcreteBuilderB{}
	director := NewDirector(builder)
	director.Construct()
	product := builder.GetProduct()
	if product != expertB {
		t.Fatalf("error||expert=%v||valule=%v", expertB, product)
	}
	//output
	//cola|chips|McChicken
}
