package main

func Sum(n []int) int {
	var s int

	for _, v := range n {
		s += v
	}

	return s
}