package singleton

import (
	"sync"
	"sync/atomic"
)

var initialized uint32
var mu sync.Mutex

// instance2 单例
var instance2 *Singleton

// instance2 单例方法
func GetInstance2() *Singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance2
	}
	mu.Lock()
	defer mu.Unlock()
	if initialized == 0 {
		instance2 = &Singleton{}
		atomic.StoreUint32(&initialized, 1)
	}
	return instance2
}
