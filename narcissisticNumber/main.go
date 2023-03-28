package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Narcissistic ", nasrcissistic(1634))
}

func nasrcissistic(n int) bool {
	toStr := fmt.Sprintf("%d", n)

	total := 0

	for _, i := range strings.Split(toStr, "") {
		a, _ := strconv.Atoi(i)
		loop := a
		for i := 1; i < len(toStr); i++ {
			loop = loop * a
		}

		total = total + loop
	}

	nascissistic := false
	if total == n {
		nascissistic = true
	}

	return nascissistic
}
