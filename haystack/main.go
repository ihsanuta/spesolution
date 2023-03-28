package main

import "fmt"

func main() {
	params := []string{
		"red",
		"blue",
		"yellow",
		"black",
	}
	idx := findNeedle(params, "red")
	fmt.Println(idx)
}

func findNeedle(params []string, key string) int {
	idx := 0
	for i, v := range params {
		if v == key {
			idx = i
			break
		}
	}

	return idx
}
