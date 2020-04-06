package proxy

import "testing"

func Test_proxy(t *testing.T) {
	p := NewSubjectProxy()
	p.doAction()
}
