package partition

import (
	// "fmt"
	// "bytes"
	// "strings"
	// "strconv"
)

// type Partition struct {
// 	cycle_type []int8
// }

// func Key(cycle_type []int8) string {
// 	var groups = GroupPartition(cycle_type)
// 	// fmt.Println(groups)
// 	var buffer bytes.Buffer
// 	// buffer.WriteString("")
// 	var i = 0
// 	for ; i < len(groups); i++ {
// 		var c = groups[i]
// 		if len(c) != 0 {
// 			buffer.WriteString(strconv.Itoa(int(c[0])))
// 			buffer.WriteString("^")
// 			buffer.WriteString(strconv.Itoa(int(c[1])))
// 			if i != len(groups) - 1 {
// 				buffer.WriteString("|")
// 			}
// 		}
// 	}
// 	// for _, c := range groups {
// 	// 	if len(c) != 0 {
// 	// 		buffer.WriteString(strconv.Itoa(int(c[0])))
// 	// 		buffer.WriteString("^")
// 	// 		buffer.WriteString(strconv.Itoa(int(c[1])))
// 	// 		buffer.WriteString("|")
// 	// 	}
// 	// }
// 	// fmt.Println(buffer.String())
// 	return buffer.String()
// }

// // func Key(partition Partition) string {
// // 	return Key(partition.cycle_type)
// // }

// func Value(cycle_type string) []int8 {
// 	// fmt.Println("V start")
// 	var result = []int8{}
// 	for _, item := range strings.Split(cycle_type, "|") {
// 		// fmt.Println(item)
// 		if item != "" {
// 			var sub = strings.Split(item, "^")
// 			// fmt.Println(sub)
// 			le, _ := strconv.Atoi(sub[0])
// 			re, _ := strconv.Atoi(sub[1])
// 			// fmt.Println(le, re)
// 			var i = 0
// 			for ; i < re; i++ {
// 				result = append(result, int8(le))
// 			}
// 		}
// 	}
// 	// fmt.Println("V end")
// 	return result
// }

// func GroupPartition(p []int8) [][]int8 {
// 	if len(p) == 0 {
// 		return [][]int8{[]int8{}}
// 	}
// 	var i, j = 0, 0
// 	var groups = [][]int8{}
// 	for i < len(p) {
// 		for j = i; j < len(p); j++ {
// 			if p[i] != p[j] {
// 				break
// 			}
// 		}
// 		groups = append(groups, []int8{p[i],int8(j-i)})
// 		i=j
// 	}
// 	return groups
// }

func PartitionsOf(n int8) [][]int8 {
	var partitions = Create(n, n);
	var result = [][]int8{}
	var i = 0;
	for ; i < len(partitions); i++ {
		var partition = partitions[i]
		result = append(result, partition)
	}
	return result
}

func Create(n int8, w int8) [][]int8{
	if n == 0 {
		return [][]int8{[]int8{}}
	}
	if w <= 0 {
		return [][]int8{}
	}
	var result = [][]int8{}
	var row = Min(n, w)
	for ; row >= 1; row-- {
		var smaller = Create(n - row, row)
		for i := range smaller {
			smaller[i] = append([]int8{row}, smaller[i]...) // append to first
		}
		result = append(result, smaller...)
	}
	return result
}

func Min(a int8, b int8) int8 {
	if a < b {
		return a
	}
	return b
}
