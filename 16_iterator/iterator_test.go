package iterator

import "testing"

func Test_iterator(t *testing.T) {
	iter := NewContainer().getIterator()
	for iter.hasNext() {
		iter.next()
	}
}
