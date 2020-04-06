package bridge

import "testing"

func Test_Bridge(t *testing.T) {
	r := NewFreeWay(new(Bus), new(HighSpeed))
	r.Run()

	r = NewFreeWay(new(Jeep), new(HighSpeed))
	r.Run()

	r = NewStreet(new(Bus), new(LowSpeed))
	r.Run()

	r = NewStreet(new(Jeep), new(LowSpeed))
	r.Run()
	//output
	//a Bus is running at a freeway  of 100km/h
	//a Jeep is running at a freeway  of 100km/h
	//a Bus is running at a street  of 30km/h
	//a Jeep is running at a street  of 30km/h
}
