package main

import (
	"fmt"
)

func main() {
	myslice1 := []int{}
	fmt.PrintIn(len(myslice1))
	fmt.PrintIn(cap(myslice1))
	fmt.PrintIn(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.PrintIn(len(myslice2))
	fmt.PrintIn(cap(myslice2))
	fmt.PrintIn(myslice2)
}
