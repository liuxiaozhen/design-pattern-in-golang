package decorator

import "fmt"

type Component interface {
	Operation()
}

type ConcreteComponent struct {
}

func (c *ConcreteComponent) Operation() {
	fmt.Println("component operation")
}

func NewConcreteComponent() *ConcreteComponent {
	return &ConcreteComponent{}
}

type Decorator interface {
	Component
	beforeOperation()
	afterOperation()
}

// 具体的装饰器类
type ConcreteDecorator struct {
	Component
}

func NewConcreteDecorator(c Component) Component {
	return &ConcreteDecorator{
		c,
	}
}

func (c *ConcreteDecorator) beforeOperation() {
	fmt.Println("before Operation")
}

func (c *ConcreteDecorator) afterOperation() {
	fmt.Println("after Operation")
}

func (c *ConcreteDecorator) Operation() {
	c.beforeOperation()
	c.Component.Operation()
	c.afterOperation()
}
