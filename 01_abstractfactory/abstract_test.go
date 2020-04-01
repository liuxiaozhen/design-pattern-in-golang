package abstractfactory

import (
	"fmt"
	"testing"
)

func makePC(pc PCFactory) {
	fmt.Println(pc.CreateCPU().CPU())
	fmt.Println(pc.CreateMotherBoard().MotherBoard())
	fmt.Println(pc.CreateHardDisk().HardDisk())
}

func Test_Intel(t *testing.T) {
	var pc PCFactory
	pc = &IntelFactory{}
	makePC(pc)
	//output
	//Intel CPU
	//Intel MotherBoard
	//Intel HardDisk
}

func Test_AMD(t *testing.T) {
	var pc PCFactory
	pc = &AMDFactory{}
	makePC(pc)
	//output
	//AMD CPU
	//AMD MotherBoard
	//AMD HardDisk
}
