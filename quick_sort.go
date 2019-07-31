package main

import (
	"fmt"
)

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}


func main() {
	fmt.Println("Welcome to the playground!")

	inputSlice := []int{1, 3, 5, 3, 2, 4, 6, 8, 4, 5, 6}

	fmt.Println("Result:", n_k(inputSlice, len(inputSlice), 8))
}
