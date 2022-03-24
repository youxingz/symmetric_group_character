package safe_struct

import (
	"sync"
	b "math/big"
)

type SafeMap struct {
    sync.RWMutex
    Map map[int32]map[int32]b.Int
}

func NewSafeMap(size int) *SafeMap {
    sm := new(SafeMap)
    sm.Map = make(map[int32]map[int32]b.Int, size)
    return sm
}

func (sm *SafeMap) ReadMap(key int32) map[int32]b.Int {
    sm.RLock()
    value := sm.Map[key]
    sm.RUnlock()
    return value
}

func (sm *SafeMap) WriteMap(key int32, value map[int32]b.Int) {
    sm.Lock()
    sm.Map[key] = value
    sm.Unlock()
}
