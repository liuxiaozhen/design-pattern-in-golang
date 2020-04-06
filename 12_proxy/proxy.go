package proxy

import "fmt"

type subject interface {
	doAction()
}

type realSubject struct {
}

func (r *realSubject) doAction() {
	fmt.Println("real do something")
}

type subjectProxy struct {
	real *realSubject
}

func (r *subjectProxy) doAction() {
	fmt.Println("proxy do control access")
	r.real.doAction()
}

// 只暴露Proxy构造接口
func NewSubjectProxy() *subjectProxy {
	return &subjectProxy{
		real: &realSubject{},
	}
}
