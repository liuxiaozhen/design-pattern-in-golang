package iterator

import "fmt"

type iterator interface {
	hasNext() bool
	next()
}

type Container interface {
	getIterator() iterator
}

type iteratorImpl struct {
	index int
	len   int
	data  []interface{}
}

func (i *iteratorImpl) hasNext() bool {
	if i.index < i.len {
		return true
	}
	return false
}

func (i *iteratorImpl) next() {
	fmt.Println(i.data[i.index])
	i.index++
}

func NewIteratorImpl(data []interface{}) iterator {
	return &iteratorImpl{
		index: 0,
		len:   len(data),
		data:  data,
	}
}

type container struct {
	data []interface{}
}

func (c *container) getIterator() iterator {
	return NewIteratorImpl(c.data)
}

func NewContainer() *container {
	data := []interface{}{1, 2, "ibm", 4, "apple", 6, 7}
	return &container{
		data,
	}
}
