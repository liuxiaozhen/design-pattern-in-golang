package singleton

import (
	"sync"
)

// Singleton 是单例模式类
type Singleton struct{}

// instance 单例
var instance *Singleton
var once sync.Once

//GetInstance 用于获取单例模式对象
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})

	return instance
}
