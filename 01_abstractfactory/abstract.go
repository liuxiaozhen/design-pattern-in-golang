package abstractfactory

// 产品等级
type CPU interface {
	CPU() string
}

// 产品等级
type MotherBoard interface {
	MotherBoard() string
}

// 产品等级
type HardDisk interface {
	HardDisk() string
}

// 抽象工厂方法
type PCFactory interface {
	CreateCPU() CPU
	CreateMotherBoard() MotherBoard
	CreateHardDisk() HardDisk
}

// AMD
type AMDCPU struct{}

func (*AMDCPU) CPU() string {
	return "AMD CPU"
}

type AMDMotherBoard struct{}

func (*AMDMotherBoard) MotherBoard() string {
	return "AMD MotherBoard"
}

type AMDHardDisk struct{}

func (*AMDHardDisk) HardDisk() string {
	return "AMD HardDisk"
}

type AMDFactory struct{}

func (*AMDFactory) CreateCPU() CPU {
	return &AMDCPU{}
}
func (*AMDFactory) CreateMotherBoard() MotherBoard {
	return &AMDMotherBoard{}
}
func (*AMDFactory) CreateHardDisk() HardDisk {
	return &AMDHardDisk{}
}

// Intel
type IntelCPU struct{}

func (*IntelCPU) CPU() string {
	return "Intel CPU"
}

type IntelMotherBoard struct{}

func (*IntelMotherBoard) MotherBoard() string {
	return "Intel MotherBoard"
}

type IntelHardDisk struct{}

func (*IntelHardDisk) HardDisk() string {
	return "Intel HardDisk"
}

type IntelFactory struct{}

func (*IntelFactory) CreateCPU() CPU {
	return &IntelCPU{}
}
func (*IntelFactory) CreateMotherBoard() MotherBoard {
	return &IntelMotherBoard{}
}
func (*IntelFactory) CreateHardDisk() HardDisk {
	return &IntelHardDisk{}
}
