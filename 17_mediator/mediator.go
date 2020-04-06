package mediator

import (
	"fmt"
	"sync"
	"time"
)

type iMediator interface {
	sendMessage(*User, string)
}

// 聊天室作为中介者,负责收发消息
type chatRoom struct {
}

func (c *chatRoom) sendMessage(user *User, msg string) {
	fmt.Printf("%v||[%s]说:%s\n", time.Now().Format("2006-01-02T 15:04:05"), user.GetName(), msg)
}

var chatImpl *chatRoom
var once sync.Once

func GetChatRoomInstance() *chatRoom {
	once.Do(func() {
		chatImpl = &chatRoom{}
	})
	return chatImpl
}

type User struct {
	name string
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) sendMessage(msg string) {
	GetChatRoomInstance().sendMessage(u, msg)
}

func NewUser(name string) *User {
	return &User{name: name}
}
