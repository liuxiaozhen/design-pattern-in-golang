package mediator

import "testing"

func Test_mediator(t *testing.T) {
	Joe := NewUser("Joe")
	Will := NewUser("Will")
	Joe.sendMessage("你好！")
	Will.sendMessage("你也好！")
}
