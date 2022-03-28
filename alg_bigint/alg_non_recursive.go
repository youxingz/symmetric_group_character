package alg_bigint

import (
	sf "symmetric_group/alg_bigint/safe_struct"
	pt "symmetric_group/partition"
	"os"
	"fmt"
	"sync"
	"time"
	b "math/big"
)

func CharacterTableSchur(n int8, file *os.File, multipleThread bool) { // just print it here
	start := time.Now().UnixNano()

	partitions := pt.PartitionsOf(n)
	p_size := len(partitions)
	fmt.Println("Partition Length:", len(partitions))
	if multipleThread {
		fmt.Println("mode: Multiple Threads")
	}
	fmt.Println("Calculating...")

	table := sf.NewSafeMap(p_size)

	if multipleThread {
		var wgp sync.WaitGroup
		for _, col := range partitions {
			wgp.Add(1)
			go func(safeMap *sf.SafeMap, partition []int8, wg *sync.WaitGroup) {
				defer wg.Done()
				if file == nil {
					PowerSumToSchur(partition)
				} else {
					// table[PKey(partition)] = PowerSumToSchur(partition)
					safeMap.WriteMap(pt.PKey(partition), PowerSumToSchur(partition))
				}
			}(table, col, &wgp)
		}
		wgp.Wait()
	} else {
		for _, col := range partitions {
			// table[PKey(col)] = PowerSumToSchur(col)
			table.WriteMap(pt.PKey(col), PowerSumToSchur(col))
		}
	}

	timeAfterCalculating := time.Now().UnixNano()
	fmt.Println("Time Spent:      ", (timeAfterCalculating - start)/1000000, "ms")

	// File writing...
	if file != nil {
		fmt.Println("File Writing...")
		fmt.Fprintf(file, "Character table of Sysmmetric Group (n=%d)\n", n)
		fmt.Fprintf(file, "%dX%d\t", p_size, p_size)
		for _, row := range partitions {
			fmt.Fprint(file, row, "\t")
		}
		fmt.Fprintln(file)
		for _, col := range partitions {
			// var comb = table[PKey(col)]
			var comb = table.ReadMap(pt.PKey(col))
			fmt.Fprint(file, col, "\t")
			for _, row := range partitions {
				value, _ := comb[pt.PKey(row)]
				fmt.Fprint(file, value.String(), "\t")
			}
			fmt.Fprintln(file)
		}
	}

	fmt.Println("Time Spent:      ", (time.Now().UnixNano() - timeAfterCalculating)/1000000, "ms")
	fmt.Println("Total Time Spent:", (time.Now().UnixNano() - start)/1000000, "ms")
}

var BIG_0 = b.NewInt(0)

func PowerSumToSchur(partition []int8) map[int32]b.Int {
	var acc = map[int32]b.Int{}
	acc[pt.PKey([]int8{})] = *b.NewInt(1)
	var i = len(partition) - 1
	for ; i >= 0; i-- {
		// fmt.Println(acc)
		var prod = map[int32]b.Int{}
		for part, coeff := range Support(acc) {
			// fmt.Println("PV==>", part, Value(part))
			IterateBorderStrip(pt.PValue(part), partition[i], func (shape []int8, height int8) {
				// fmt.Println(shape)
				var key = pt.PKey(shape)
				var origin_value = prod[key]
				// fmt.Println("KO==>", key, origin_value)
				if height % 2 == 0 {
					// prod[key] = origin_value + coeff
					prod[key] = *(origin_value.Add(&origin_value, &coeff))
				} else {
					// prod[key] = origin_value - coeff
					prod[key] = *(origin_value.Sub(&origin_value, &coeff))
				}
			})
		}
		// fmt.Println(prod)
		acc = prod
	}
	// fmt.Println(acc)
	return acc
}

func Support(comb map[int32]b.Int) map[int32]b.Int {
	var result = map[int32]b.Int{}
	for k, v := range comb {
		// if v != 0 {
		// 	result[k] = v
		// }
		if v.Cmp(BIG_0) != 0 {
			result[k] = v
		}
	}
	return result
}

type SaveResult func(shape []int8, height int8)

func IterateBorderStrip(shape []int8, length int8, save SaveResult) {
	// fmt.Println("args: ", shape, length)
	var length_of_shape = int8(len(shape))
	var inner = make([]int8, length_of_shape + length)
	copy(inner, shape)

	var length_of_inner = int8(len(inner))
	var outer = make([]int8, length_of_inner)
	// fmt.Println(">>>> outer: ", outer)
	copy(outer, inner)
	// fmt.Println(">>>> outer: ", outer)
	if len(inner) == 0 {
		fmt.Println("args: ", shape, length)
	}

	var B_row = int8(0)
	var B_col = inner[0]
  var T_col = inner[0] + length - 1
	var T_row = int8(0)

	// fmt.Println("inner", inner)
	// fmt.Println("outer", outer)

	for true {
		var L1dist = (B_row - T_row) + (T_col - B_col) + 1
		// fmt.Println(">> While: ", L1dist, B_row, T_row)
		if L1dist == length {
			outer[T_row] = T_col + 1
			var r = T_row + 1
			for ; r <= B_row; r++ {
				outer[r] += inner[r - 1] - inner[r] + 1
			}
			// fmt.Println("outer:", outer)
			var height = B_row - T_row
			var outer_shape = RemoveZeros(outer)
			// fmt.Println("call: ", height, outer_shape)
			save(outer_shape, height)
			// fmt.Println(outer_shape)
			for r = T_row; r <= B_row; r++ {
				outer[r] = inner[r]
			}
		}
		if L1dist <= length {
			if B_row + 1 == length_of_inner {
				return
			}
			B_row += 1
			B_col = inner[B_row]
		}
		if L1dist > length {
			if T_col == 0 {
				return
			}
			T_col -= 1
			for inner[T_row] > T_col {
				T_row += 1
			}
		}
	}
}
func RemoveZeros(array []int8) []int8 {
	var result = []int8{}
	for _, i := range array {
		if i != 0 {
			result = append(result, i)
		}
	}
	return result
}
