package adapter

// 目标抽象类定义客户所需目标接口
type Target interface {
	Request() string
}

// 是被适配的目标接口
type Adaptee interface {
	SpecificRequest() string
}

// 定义一个Adaptee具体实现类
type adapteeImpl struct{}

func (adaptee *adapteeImpl) SpecificRequest() string {
	return "Specific Request"
}

// Adaptee具体类的构造方法
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// Adapter 是转换Adaptee为Target接口的适配器
type Adapter struct {
	Adaptee
}

// 实现Target接口
func (a *Adapter) Request() string {
	return a.SpecificRequest()
}

func NewAdapter(adaptee Adaptee) Target {
	return &Adapter{
		adaptee,
	}
}
