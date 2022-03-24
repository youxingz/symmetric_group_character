package main

var primes = []int64{2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,51,57,59,61,67,71,73,79,83,89,97,101,103,107,109,113,127,131,137,139,149,151,157,163,167,173,179,181,191,193,197,199,211,223,227,229}

func K(partition []int8) int64 {
	return (prod(partition) << 6) + sum(partition)
}

func sum(numbers []int8) int64 {
	var s = int64(0)
	for _, v := range numbers {
		s += primes[v] // int64(v + 1)
	}
	return s
}
func prod(numbers []int8) int64 {
	var s = int64(1)
	for _, v := range numbers {
		s *= primes[v] // int64(v + 1)
		s = s % (1 << 62)
	}
	return s
}