package main

import (
	"fmt"
	"learn-go-with-test/iteration"
)

func main()  {
	c := "d"
	r := 5
	fmt.Printf("\n%v will be repeated %v times\n\tresults: %v\n\n", c, r, iteration.Repeat(c, r))
}