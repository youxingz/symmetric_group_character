package main

import (
	"os"
	"fmt"
	"strconv"
	alg "symmetric_group/alg_bigint"
	alg_slice "symmetric_group/alg_slice"
)

func main() {
	N, _ := strconv.Atoi(os.Args[1])
	R := false
	MultipleThread := false
	SliceMode := false
	if len(os.Args) > 2 {
		if os.Args[2] == "-r" {
			R = true
		}
		if os.Args[2] == "-m" {
			MultipleThread = true
		}
		if os.Args[2] == "-s" {
			SliceMode = true
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

	if SliceMode {
		dir := "output_large"
		os.Mkdir(dir, os.ModePerm)
		alg_slice.CharacterTableNonRecursiveMultipleThread(int8(N), dir)
		return
	}

	dir := "output"
	os.Mkdir(dir, os.ModePerm)

	filename := fmt.Sprintf("%s/character_table (S_%d).txt", dir, N)
	file, fileErr := os.Create(filename)
	if fileErr != nil {
		panic(fileErr)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
				panic(err)
		}
	}()

	// file = nil // test without print into file.

	if R {
		fmt.Printf("[Character table of Sysmmetric Group (n=%d)] Recursive Version\n", N)
		alg.CharacterTableRecursive(int8(N), file)
	} else {
		fmt.Printf("[Character table of Sysmmetric Group (n=%d)] Non-Recursive Version\n", N)
		alg.CharacterTableSchur(int8(N), file, MultipleThread)
	}
	fmt.Printf("Output file: %s/%s\n", dir, filename)
	fmt.Println("Done!")
}
