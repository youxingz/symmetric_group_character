package alg_bigint

import (
	pt "symmetric_group/partition"
	sf "symmetric_group/alg_bigint/safe_struct"
	"os"
	"fmt"
	"sync"
	"time"
	b "math/big"
)

func CharacterTableRecursive(n int8, file *os.File) {
	start := time.Now().UnixNano()

	partitions := pt.PartitionsOf(n)
	p_size := len(partitions)
	fmt.Println("Partition Length:", len(partitions))
	fmt.Println("Calculating...")

	table := sf.NewSafeTable(p_size)

	// body
	if false { // we do not support multiple threads
		var wgp sync.WaitGroup
		for i:=p_size-1; i >= 0; i-- {
			row := pt.PKey(partitions[i])
			for j:=p_size-1; j >= 0; j-- {
				col := pt.PKey(partitions[j])
				wgp.Add(1)
				go func(prow []int8, pcol []int8, wg *sync.WaitGroup) {
					defer wg.Done()
					var value = Calculate(prow, pcol)
					table.WriteTable(row, col, value)
				}(partitions[i], partitions[j], &wgp)
			}
		}
		wgp.Wait()
	} else {
		for i:=p_size-1; i >= 0; i-- {
			row := pt.PKey(partitions[i])
			for j:=p_size-1; j >= 0; j-- {
				var value = Calculate(partitions[i], partitions[j])
				if file != nil {
					col := pt.PKey(partitions[j])
					table.WriteTable(row, col, value)
				}
			}
		}
	}


	timeAfterCalculating := time.Now().UnixNano()
	fmt.Println("Time Spent:      ", (timeAfterCalculating - start)/1000000, "ms")

	// File Writing...
	if file != nil {
		fmt.Println("File Writing...")
		fmt.Fprintf(file, "Character table of Sysmmetric Group (n=%d)\n", n)
		fmt.Fprintf(file, "%dX%d\t", p_size, p_size)

		// header
		for j:=p_size-1; j >= 0; j-- {
			fmt.Fprint(file, partitions[j], "\t")
		}
		fmt.Fprintln(file)
			// body
		for i:=p_size-1; i >= 0; i-- {
			fmt.Fprint(file, partitions[i], "\t")
			row := pt.PKey(partitions[i])
			for j:=p_size-1; j >= 0; j-- {
				col := pt.PKey(partitions[j])
				value, _ := table.ReadTable(row, col)
				// var value = Calculate(partitions[i], partitions[j])
				fmt.Fprint(file, value.String(), "\t")
			}
			fmt.Fprintln(file)
		}
	}

	fmt.Println("Time Spent:      ", (time.Now().UnixNano() - timeAfterCalculating)/1000000, "ms")
	fmt.Println("Total Time Spent:", (time.Now().UnixNano() - start)/1000000, "ms")
}

var CACHE = sf.NewSafeTable(20000) // default size

func Calculate(lambda []int8, rho []int8) b.Int {
	// fmt.Println(lambda, rho)
	if (len(lambda) == 0 && len(rho) == 0) {
		return *b.NewInt(1)
	}
	// cache
	var K_lambda = pt.PKey(lambda)
	var K_rho = pt.PKey(rho)
	// if val, ok := CACHE[K_lambda][K_rho]; ok {
	if val, ok := CACHE.ReadTable(K_lambda, K_rho); ok {
			return val
	}
	var subShape = rho[1:]
	var borderStripLength = rho[0];

	// var sum = int64(0);
	var sum = b.NewInt(0)

	var T_row = int8(0)
	var T_col = lambda[0] - int8(1)
	var B_row = T_row
	var B_col = T_col

	for true {
			var distance = (B_row - T_row) + (T_col - B_col) + int8(1)
			if distance == borderStripLength {
					if !(T_col + 1 < lambda[T_row]) { // make sure T is at right start col.
							var heightCoeff = int64(0)
							if (B_row - T_row) % 2 == 0 {
								heightCoeff = 1
							} else {
								heightCoeff = -1
							}
							// cut border strip
							var subLambda = make([]int8, len(lambda))
							copy(subLambda, lambda)
							if T_row == B_row {
									subLambda[B_row] -= borderStripLength
							} else {
								var i = int8(0)
									for i = T_row; i < B_row; i++ {
											subLambda[i] = (subLambda[i + 1] - 1)
									}
									subLambda[B_row] = (B_col)
							}
							if (IsWeeklyDecreasing(subLambda)) {
								// sum += Calculate(_RemoveZeros(subLambda), subShape) * heightCoeff
								var subSum = Calculate(_RemoveZeros(subLambda), subShape)
								if (heightCoeff == -1) {
									sum.Sub(sum, &subSum)
								} else {
									sum.Add(sum, &subSum)
								}
							}
					}
					distance--;  // B++
			}
			if distance < borderStripLength { // B++
					if B_row == int8(len(lambda) - 1) { // touch bottom!
							if (B_col > 0) { B_col-- } else { break }
					} else {
							if B_col + 1 > lambda[B_row + 1] {
									B_col--
							} else {
									B_row++
							}
					}
			}
			if distance > borderStripLength { // T++
					if T_row == int8(len(lambda) - 1) { // touch bottom!
							if (T_col > 0) { T_col-- } else { break }
					} else {
							if T_col + 1 > lambda[T_row + 1] {
									T_col--
							} else {
									T_row++
							}
					}
			}
	}
	// if _, ok := CACHE[K_lambda]; !ok {
	// 	CACHE[K_lambda] = make(map[int32]int64)
	// }
	// CACHE[K_lambda][K_rho] = sum
	CACHE.WriteTable(K_lambda, K_rho, *sum)
	return *sum
}

func IsWeeklyDecreasing(numbers []int8) bool {
	var i = 0
	var max = len(numbers) - 1
	for ; i < max; i++ {
		if numbers[i] < numbers[i+1] {
			return false
		}
	}
	return true
}

func _RemoveZeros(array []int8) []int8 {
	var result = []int8{}
	for _, i := range array {
		if i != 0 {
			result = append(result, i)
		}
	}
	return result
}
