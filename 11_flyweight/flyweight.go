package flyweight

import (
	"sync"
)

// 创建一个文件享元,readOnly
type FileFlyweight struct {
	Text string
}

func (f *FileFlyweight) Data() string {
	return f.Text
}

func (f *FileFlyweight) LoadText(name string) {
	f.Text = name + " with a big Text"
}

func NewFileFlyweight(fileName string) *FileFlyweight {
	f := &FileFlyweight{
		fileName,
	}
	f.LoadText(fileName)
	return f
}

// 享元工厂结构
type FileFlyweightFactory struct {
	pool map[string]*FileFlyweight
}

func (fc *FileFlyweightFactory) Get(file string) *FileFlyweight {
	if f, ok := fc.pool[file]; ok {
		return f
	} else {
		fc.pool[file] = NewFileFlyweight(file)
		return fc.pool[file]
	}

}

// 单例模式构造flyweightFactory
var flyweightFactory *FileFlyweightFactory
var once sync.Once

func NewFileFlyweightFactory() *FileFlyweightFactory {
	once.Do(
		func() {
			flyweightFactory = &FileFlyweightFactory{
				make(map[string]*FileFlyweight),
			}
		})
	return flyweightFactory
}
