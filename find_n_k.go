package main

import (
	"fmt"
)

func n_k(a []int, n int, k int) int {
	if n < 2 && n > 0 && n >= k {
		return a[0]
	}

	left, right := 0, n-1

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	if k == n-left {
		return a[left]
	} else if k < n-left {
		return n_k(a[left+1:], len(a[left+1:]), k)
	} else {
		return n_k(a[:left], len(a[:left]), k-(n-left))
	}
}

func main() {
	fmt.Println("Welcome to the playground!")

	inputSlice := []int{1, 3, 5, 3, 2, 4, 6, 8, 4, 5, 6}

	fmt.Println("Result:", n_k(inputSlice, len(inputSlice), 8))
}
