package facade

import "fmt"

type Boot interface {
	boot()
}

type Computer struct {
	CPU  Boot
	Mem  Boot
	Disk Boot
}

func (c *Computer) boot() {
	c.CPU.boot()
	c.Mem.boot()
	c.Disk.boot()
}

type CPU struct {
}

func (c *CPU) boot() {
	fmt.Println("cpu boot up")
}

type Memory struct {
}

func (m *Memory) boot() {
	fmt.Println("memory boot up")
}

type Disk struct {
}

func (d *Disk) boot() {
	fmt.Println("disk boot up")
}

func NewComputer() Boot {
	return &Computer{
		&CPU{},
		&Memory{},
		&Disk{},
	}
}
