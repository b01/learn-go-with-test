package main

func Sum(n []int) int {
	var s int

	for _, v := range n {
		s += v
	}

	return s
}

func SumAll(numbers ...[]int) ([]int) {
	var sums []int

	for _, v := range numbers {
		sums = append(sums, Sum(v))
	}

	return sums
}