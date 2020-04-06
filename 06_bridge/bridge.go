package bridge

import "fmt"

// 抽象接口 Abstraction
type Road interface {
	Run()
}

// Implementor
type Car interface {
	GetCar() string
}

// Implementor
type Speed interface {
	GetSpeed() string
}

// RefinedAbstraction
type FreeWay struct {
	implCar   Car
	implSpeed Speed
}

func (h *FreeWay) Run() {
	fmt.Printf("a %v is running at a freeway  of %v/h \n", h.implCar.GetCar(), h.implSpeed.GetSpeed())
}

func NewFreeWay(car Car, speed Speed) Road {
	return &FreeWay{
		car,
		speed,
	}
}

// RefinedAbstraction
type Street struct {
	implCar   Car
	implSpeed Speed
}

func (h *Street) Run() {
	fmt.Printf("a %v is running at a street  of %v/h \n", h.implCar.GetCar(), h.implSpeed.GetSpeed())
}

func NewStreet(car Car, speed Speed) Road {
	return &Street{
		car,
		speed,
	}
}

// ConcreteImplementor
type Bus struct {
}

func (*Bus) GetCar() string {
	return "Bus"
}

// ConcreteImplementor
type Jeep struct {
}

func (*Jeep) GetCar() string {
	return "Jeep"
}

type HighSpeed struct {
}

func (*HighSpeed) GetSpeed() string {
	return "100km"
}

type LowSpeed struct {
}

func (*LowSpeed) GetSpeed() string {
	return "30km"
}
