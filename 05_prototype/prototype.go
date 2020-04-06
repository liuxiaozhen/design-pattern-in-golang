package prototype

//Cloneable 是原型对象需要实现的接口
type Cloneable interface {
	Clone() Cloneable
}

// 原型管理器,通过map来保存一个原型对象
type PrototypeManager struct {
	pMap map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		make(map[string]Cloneable),
	}
}

func (m *PrototypeManager) Set(proName string, v Cloneable) {
	m.pMap[proName] = v
}

func (m *PrototypeManager) Get(proName sring) Cloneable {
	return m.pMap[proName].Clone()
}
