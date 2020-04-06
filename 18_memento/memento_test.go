package memento

import (
	"fmt"
	"testing"
)

func Test_memento(t *testing.T) {
	o := NewOriginator(3)
	//创建存档
	c := NewCaretaker(o.createMemento())
	o.setState(5)
	fmt.Println(o.getState())
	//从存档加载
	o.setMemento(c.getMemento())
	fmt.Println(o.getState())
}
