package command

import (
	"fmt"
)

// 命令接口
type Command interface {
	Execute()
}

// Receiver接口
type Receiver interface {
	On()
	Off()
}

type TurnonCommand struct {
	receiver Receiver
}

func (c *TurnonCommand) Execute() {
	c.receiver.On()
}

func NewTurnonCommand(r Receiver) *TurnonCommand {
	return &TurnonCommand{
		receiver: r,
	}
}

type TurnoffCommand struct {
	receiver Receiver
}

func (c *TurnoffCommand) Execute() {
	c.receiver.Off()
}
func NewTurnoffCommand(r Receiver) *TurnoffCommand {
	return &TurnoffCommand{
		receiver: r,
	}
}

type LightReveiver struct {
}

func (light *LightReveiver) On() {
	fmt.Println("light turn on")
}

func (light *LightReveiver) Off() {
	fmt.Println("light turn off")
}

type TVReveiver struct {
}

func (tv *TVReveiver) On() {
	fmt.Println("tv turn on")
}

func (tv *TVReveiver) Off() {
	fmt.Println("tv turn off")
}

type Invoker struct {
	status bool
	on     Command
	off    Command
}

func (i *Invoker) setCommand(on Command, off Command) {
	i.on = on
	i.off = off
}

func (i *Invoker) Press() {
	if i.status == false {
		i.on.Execute()
		i.status = true
	} else {
		i.off.Execute()
		i.status = false
	}
}

func NewInvoker() *Invoker {
	return &Invoker{
		status: false,
		on:     nil,
		off:    nil,
	}
}
