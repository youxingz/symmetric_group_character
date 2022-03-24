package safe_struct

import (
	"sync"
	b "math/big"
)

type SafeTable struct {
    sync.RWMutex
    Map map[int32]map[int32]b.Int
}

func NewSafeTable(size int) *SafeTable {
    sm := new(SafeTable)
    sm.Map = make(map[int32]map[int32]b.Int, size)
    return sm
}

func (sm *SafeTable) ReadTable(row int32, col int32) (b.Int, bool) {
    sm.RLock()
		val, ok := sm.Map[row][col]
    sm.RUnlock()
    return val, ok
}

func (sm *SafeTable) WriteTable(row int32, col int32, value b.Int) {
    sm.Lock()
		if rows, ok := sm.Map[row]; ok {
			rows[col] = value // ref?
		} else {
			rows := make(map[int32]b.Int)
			rows[col] = value
			sm.Map[row] = rows
		}
    sm.Unlock()
}
