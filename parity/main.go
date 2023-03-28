package main

import (
	"fmt"
	"math"
)

func main() {
	outlier := findOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36})
	fmt.Println(outlier)
}

func findOutlier(integers []int) int {
	length := len(integers)
	var evens []int
	var odds []int

	for i := 0; i < length; i++ {
		if integers[i]%2 == 0 {
			evens = append(evens, integers[i])
		}

		a := integers[i] % 2
		if math.Abs(float64(a)) == 1 {
			odds = append(odds, integers[i])
		}
	}
	if len(evens) > len(odds) {
		return odds[0]
	} else {
		return evens[0]
	}
}
