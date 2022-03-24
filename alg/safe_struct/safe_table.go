package safe_struct

import "sync"

type SafeTable struct {
    sync.RWMutex
    Map map[int32]map[int32]int64
}

func NewSafeTable(size int) *SafeTable {
    sm := new(SafeTable)
    sm.Map = make(map[int32]map[int32]int64, size)
    return sm
}

func (sm *SafeTable) ReadTable(row int32, col int32) (int64, bool) {
    sm.RLock()
		val, ok := sm.Map[row][col]
    sm.RUnlock()
    return val, ok
}

func (sm *SafeTable) WriteTable(row int32, col int32, value int64) {
    sm.Lock()
		if rows, ok := sm.Map[row]; ok {
			rows[col] = value // ref?
		} else {
			rows := make(map[int32]int64)
			rows[col] = value
			sm.Map[row] = rows
		}
    sm.Unlock()
}
