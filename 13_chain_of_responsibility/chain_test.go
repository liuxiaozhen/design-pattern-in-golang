package chain

import "testing"

func Test_chain(t *testing.T) {
	// 责任处理实体
	headman := NewHeadMan()
	director := NewDirector()
	manager := NewManager()

	// 组装责任链
	headman.setNext(director)
	director.setNext(manager)

	// 传递消息
	headman.handlerMessage(RequestLevel_0)
	headman.handlerMessage(RequestLevel_1)
	headman.handlerMessage(RequestLevel_2)

	//output
	//HeadMan confirm request Level[0]
	//Director confirm request Level[1]
	//Manager confirm request Level[2]
}
