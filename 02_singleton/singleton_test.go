package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

func TestSingleton2(t *testing.T) {
	ins1 := GetInstance2()
	ins2 := GetInstance2()
	if ins1 != ins2 {
		t.Fatal("instance2 is not equal")
	}
}
