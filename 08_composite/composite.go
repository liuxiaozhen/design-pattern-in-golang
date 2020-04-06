package composite

import (
	"fmt"
	"strings"
)

type Component interface {
	Name() string
	Add(Component)
	Remove(Component)
	Display(int)
}

type baseComponent struct {
	name string
}

func (c *baseComponent) Name() string {
	return c.name
}

func (c *baseComponent) Add(new Component) {
	return
}

func (c *baseComponent) Remove(new Component) {
	return
}

func (c *baseComponent) Display(int) {
}

type Composite struct {
	baseComponent
	childs map[string]Component
}

func NewComposite(name string) *Composite {
	return &Composite{
		baseComponent: baseComponent{name},
		childs:        make(map[string]Component),
	}
}

func (c *Composite) Add(child Component) {
	name := child.Name()
	if _, ok := c.childs[name]; !ok {
		c.childs[name] = child
	}
}

func (c *Composite) Remove(child Component) {
	name := child.Name()
	if _, ok := c.childs[name]; ok {
		delete(c.childs, name)
	}
}

func (c *Composite) Display(pretty int) {
	fmt.Printf("*%v%v\n", strings.Repeat("*", pretty), c.Name())
	for _, v := range c.childs {
		v.Display(pretty + 1)
	}
}
