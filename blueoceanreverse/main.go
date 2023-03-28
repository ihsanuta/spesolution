package main

import "fmt"

func main() {
	c := check([]int{1, 2, 3, 4, 5, 5, 5}, []int{3, 1})
	fmt.Println(c)
}

func check(params []int, p []int) []int {
	res := []int{}
	for _, v := range params {
		if !contains(p, v) {
			res = append(res, v)
		}
	}

	return res
}

func contains(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
