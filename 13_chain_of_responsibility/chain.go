package chain

import (
	"errors"
	"fmt"
)

const (
	RequestLevel_0 = iota
	RequestLevel_1
	RequestLevel_2
)

var (
	ErrorNextIsNil = errors.New("need next hander")
	ErrorLevel     = errors.New("level error")
)

type Handler interface {
	setNext(handler Handler)
	handlerMessage(level int) error //request 0:请假 1:休假 2:辞职
}

type HeadMan struct {
	nextHander Handler
	level      int
}

func (h *HeadMan) setNext(next Handler) {
	h.nextHander = next
}

func (h *HeadMan) handlerMessage(reqLevel int) error {
	if reqLevel == h.level {
		return h.response()
	} else if reqLevel > h.level {
		if h.nextHander == nil {
			return ErrorNextIsNil
		}
		return h.nextHander.handlerMessage(reqLevel)
	}
	return ErrorLevel
}

func (h *HeadMan) response() error {
	fmt.Printf("HeadMan confirm request Level[%v]\n", h.level)
	return nil
}

func NewHeadMan() *HeadMan {
	return &HeadMan{
		level: RequestLevel_0,
	}
}

type Director struct {
	nextHander Handler
	level      int
}

func (h *Director) setNext(next Handler) {
	h.nextHander = next
}

func (h *Director) handlerMessage(reqLevel int) error {
	if reqLevel == h.level {
		return h.response()
	} else if reqLevel > h.level {
		if h.nextHander == nil {
			return ErrorNextIsNil
		}
		return h.nextHander.handlerMessage(reqLevel)
	}
	return ErrorLevel
}

func (h *Director) response() error {
	fmt.Printf("Director confirm request Level[%v]\n", h.level)
	return nil
}

func NewDirector() *Director {
	return &Director{
		level: RequestLevel_1,
	}
}

type Manager struct {
	nextHander Handler
	level      int
}

func (h *Manager) setNext(next Handler) {
	h.nextHander = next
}

func (h *Manager) handlerMessage(reqLevel int) error {
	if reqLevel == h.level {
		return h.response()
	} else if reqLevel > h.level {
		if h.nextHander == nil {
			return ErrorNextIsNil
		}
		return h.nextHander.handlerMessage(reqLevel)
	}
	return ErrorLevel
}

func (h *Manager) response() error {
	fmt.Printf("Manager confirm request Level[%v]\n", h.level)
	return nil
}

func NewManager() *Manager {
	return &Manager{
		level: RequestLevel_2,
	}
}
