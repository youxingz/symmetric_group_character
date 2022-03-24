package main

import ("fmt")

func _main() {
	// fmt.Println(_R(5,0))
	// fmt.Println(_R(5,1))
	// fmt.Println(_R(5,2))
	// fmt.Println(_R(5,3))
	// fmt.Println(_R(5,4))
	fmt.Println(_R(64,64))
	// fmt.Println(_R(4,1))
	// fmt.Println(_R(2,1))

	// const n = 3
	// for k := 0; k <= n; k++ {
	// 	fmt.Println([]int{n,k}, _R(n,k))
	// }
}

func _P(n int, k int) int {
	if (n == 0 && k == 0) {
		return 1
	}
	if (n >= 1 && k == 0) {
		return 0
	}
	if (k > n) {
		return 0
	}
	// if (k <= n) {
		return _R(n-k, k)
}

func _R(n int, k int) int {
	var sum = 0
	for i := 0; i <= k; i++ {
		sum += _P(n, i)
	}
	return sum
}
