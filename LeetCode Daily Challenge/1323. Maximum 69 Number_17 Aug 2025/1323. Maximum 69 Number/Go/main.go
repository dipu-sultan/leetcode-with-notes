package main

import (
	"fmt"
	"math"
)

func maximum69Number(num int) int {
	cur := 0
	numC := num
	indFirstSix := -1

	for numC != 0 {
		if numC%10 == 6 {
			indFirstSix = cur
		}
		numC /= 10
		cur++
	}

	if indFirstSix == -1 {
		return num
	}
	return num + 3*int(math.Pow(10, float64(indFirstSix)))
}

func main() {
	// Test cases
	test1 := 9669
	fmt.Printf("Input: %d\n", test1)
	fmt.Printf("Output: %d\n", maximum69Number(test1))
}
