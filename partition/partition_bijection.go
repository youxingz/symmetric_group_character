package partition

import (
	// "fmt"
)

var N = int32(64)
// We only use the R array (the only trick is filling it correctly.)
var RData = make([]int32, N*N)

func init() {
	// Base case: P(0, 0) = R(0, 0) = 1.
	RData[0*N + 0] = 1

	// Recursive case: we already have R(n, 0) = 0, so for the other points we use
	// R(n, k) = R(n, k - 1) + P(n, k)
	var n, k = int32(0), int32(1)
	for ; n < N; n++ {
		for k = int32(1); k < N; k++ {
			var v = int32(0)
			if k <= n {
				v = RData[N*(n - k) + k]
			}
			RData[N*n + k] = RData[N*n + k - 1] + v
		}
	}
	// fmt.Println(RData)
}

func R(n int32, k int32) int32 {
	if n >= 0 && k >= 0 {
		return RData[N*n + k]
	}
	return 0
}

func PKey(cycle_type []int8) int32 {
	var size = int32(0)
	var i = 0
	for ; i < len(cycle_type); i++ {
		size += int32(cycle_type[i])
	}
	if size >= N {
		return -1
	}
	var index = int32(0)
	i = 0
	var sz = int32(size)
	for ; i < len(cycle_type) && cycle_type[i] > 0; i++ {
		var cy = int32(cycle_type[i])
		index += R(sz, cy - 1)
		sz -= cy
	}
	return (index << 6) + size
}

func PValue(index int32) []int8 {
	var parts = []int8{}
	var size = index & 0x3f
	var idx = int32(uint32(index) >> 6)

	// Here we should really be doing a binary search...
	for size > 0 {
			// The indices for partitions with first part k are [R(n, k-1), R(n, k)).
			var k = int32(1)
			for k <= size && !(idx < R(size, k)) {
				k++
			}

			if (k > size) {
				return nil
			}
			parts = append(parts, int8(k))
			idx -= R(size, k - 1)
			size -= k
	}

	return parts
}
