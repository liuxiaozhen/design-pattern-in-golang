package command

import "testing"

func Test_TVInvoker(t *testing.T) {
	tv := &TVReveiver{}
	on := NewTurnonCommand(tv)
	off := NewTurnoffCommand(tv)
	panel := NewInvoker()
	panel.setCommand(on, off)
	panel.Press()
	panel.Press()
}

func Test_LightInvoker(t *testing.T) {
	light := &LightReveiver{}
	on := NewTurnonCommand(light)
	off := NewTurnoffCommand(light)
	panel := NewInvoker()
	panel.setCommand(on, off)
	panel.Press()
	panel.Press()
}
