package main

import (
	"os"
	"fmt"
	"strconv"
	alg "symmetric_group/alg_bigint"
)

func main() {
	N, _ := strconv.Atoi(os.Args[1])
	R := false
	MultipleThread := false
	if len(os.Args) > 2 {
		if os.Args[2] == "-r" {
			R = true
		}
		if os.Args[2] == "-m" {
			MultipleThread = true
		}
	}

	if N <= 0 {
		fmt.Println("Please input the correct number (n > 0)")
		return
	}
	if N > 64 {
		fmt.Printf("Can not calculate character table of S_%d\n (n <= 64)", N)
		return
	}

	os.Mkdir("output", os.ModePerm)

	filename := fmt.Sprintf("output/character_table (S_%d).txt", N)
	file, fileErr := os.Create(filename)
	if fileErr != nil {
			fmt.Println(fileErr)
			return
	}

	// file = nil // test without print into file.

	if R {
		fmt.Printf("[Character table of Sysmmetric Group (n=%d)] Recursive Version\n", N)
		alg.CharacterTableRecursive(int8(N), file)
	} else {
		fmt.Printf("[Character table of Sysmmetric Group (n=%d)] Non-Recursive Version\n", N)
		alg.CharacterTableSchur(int8(N), file, MultipleThread)
	}
	fmt.Printf("Output file: output/%s\n", filename)
	fmt.Println("Done!")
}
