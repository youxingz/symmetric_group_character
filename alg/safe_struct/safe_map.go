package safe_struct

import "sync"

type SafeMap struct {
    sync.RWMutex
    Map map[int32]map[int32]int64
}

func NewSafeMap(size int) *SafeMap {
    sm := new(SafeMap)
    sm.Map = make(map[int32]map[int32]int64, size)
    return sm
}

func (sm *SafeMap) ReadMap(key int32) map[int32]int64 {
    sm.RLock()
    value := sm.Map[key]
    sm.RUnlock()
    return value
}

func (sm *SafeMap) WriteMap(key int32, value map[int32]int64) {
    sm.Lock()
    sm.Map[key] = value
    sm.Unlock()
}
