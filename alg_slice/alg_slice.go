package alg_slice

import (
	pt "symmetric_group/partition"
	ut "symmetric_group/alg_slice/util"
	"os"
	"fmt"
	"sync"
	"time"
	b "math/big"
	"bufio"
	"strconv"
)

const THREAD_NUMBER = 4

func CharacterTableNonRecursiveMultipleThread(n int8, dir string) {
	var wgp sync.WaitGroup
	fmt.Println("Calculating...")

	partitions := pt.PartitionsOf(n)
	p_size := len(partitions)
	part_size := 2000 // p_size / THREAD_NUMBER

	fmt.Printf("Partition Size: %d\n", p_size)

	partition_filename := fmt.Sprintf("%s/character_table (S_%d)_partitions.txt", dir, n)
	fmt.Printf("File: %s\n", partition_filename)
	PrintPartitions(partitions, partition_filename)

	start_time := time.Now().UnixNano()

	for start := 0; start < p_size; start += part_size {
		end := start + part_size
		if end > p_size {
			end = p_size
		}
		wgp.Add(1)		
		filename := fmt.Sprintf("%s/character_table (S_%d)_part(%d-%d).txt", dir, n, start, end)
		fmt.Printf("File: %s\n", filename)

		go func(partitions *[][]int8, filename string, start int, end int, wg *sync.WaitGroup) {
			defer wg.Done()
			CharacterTableSlice(*partitions, start, end, filename)
			fmt.Printf("File Writing Completed, Time Spent: %d ms \t[%s]\n", (time.Now().UnixNano() - start_time)/1000000, filename)
		}(&partitions, filename, start, end, &wgp)
	}
	wgp.Wait()
	fmt.Printf("Total Time Spent: %d ms\n", (time.Now().UnixNano() - start_time)/1000000)
	
}

func PrintPartitions(partitions [][]int8, filename string) {
	file, fileErr := os.Create(filename)
	if fileErr != nil {
		panic(fileErr)
	}
	w := bufio.NewWriter(file)

	row_spliter := []byte("\t")
	line_spliter := []byte("\n")

	for index, partition := range partitions {
		w.Write([]byte(strconv.Itoa(index)))
		w.Write(row_spliter)
		// w.Write(I2B(partition))
		w.Write(ut.PartitionToBytes(partition))
		// w.WriteString(B2S(partition))
		w.Write(line_spliter)
	}
	// save into storage
	if err := w.Flush(); err != nil {
		panic(err)
	}
}

// if n >= 20
func CharacterTableSlice(partitions [][]int8, start int, end int, filename string) {
	// file & buffer
	file, fileErr := os.Create(filename)
	if fileErr != nil {
		panic(fileErr)
	}
	w := bufio.NewWriter(file)
	// buf := make([]byte, 1024)

	part := partitions[start:end]

	tag := fmt.Sprintf("%dX%d\t", len(partitions), len(part))
	w.WriteString(tag)

	row_spliter := []byte("\t")
	line_spliter := []byte("\n")
	left_bracket := []byte("[")
	right_bracket := []byte("]")
	
	for index, _ := range partitions {
		// w.WriteString(B2S(row))
		w.Write(left_bracket)
		w.Write([]byte(strconv.Itoa(index)))
		w.Write(right_bracket)
		w.Write(row_spliter)
	}
	w.Write(line_spliter)
	for col_index, col := range part {
		comb := powerSumToSchur(col)
		// w.WriteString(B2S(col))
		w.Write(left_bracket)
		w.Write([]byte(strconv.Itoa(col_index+start)))
		w.Write(right_bracket)
		w.Write(row_spliter)
		for _, row := range partitions {
			value, _ := comb[pt.PKey(row)]
			w.WriteString(value.String())
			w.Write(row_spliter)
		}
		w.Write(line_spliter)
	}
	// save into storage
	if err := w.Flush(); err != nil {
		panic(err)
	}
}

var BIG_0 = b.NewInt(0)

func powerSumToSchur(partition []int8) map[int32]b.Int {
	var acc = map[int32]b.Int{}
	acc[pt.PKey([]int8{})] = *b.NewInt(1)
	var i = len(partition) - 1
	for ; i >= 0; i-- {
		// fmt.Println(acc)
		var prod = map[int32]b.Int{}
		for part, coeff := range support(acc) {
			// fmt.Println("PV==>", part, Value(part))
			iterateBorderStrip(pt.PValue(part), partition[i], func (shape []int8, height int8) {
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

func support(comb map[int32]b.Int) map[int32]b.Int {
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

type saveResult func(shape []int8, height int8)

func iterateBorderStrip(shape []int8, length int8, save saveResult) {
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
			var outer_shape = removeZeros(outer)
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
func removeZeros(array []int8) []int8 {
	var result = []int8{}
	for _, i := range array {
		if i != 0 {
			result = append(result, i)
		}
	}
	return result
}
