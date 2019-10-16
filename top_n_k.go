package main

import (
	"fmt"
)

func n_k_item(a []int, n int, k int) (int, bool) {
	if n < k || n <= 0 {
		return 0, false
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
		return a[left], true
	} else if k < n-left {
		return n_k_item(a[left+1:], len(a[left+1:]), k)
	} else {
		return n_k_item(a[:left], len(a[:left]), k-(n-left))
	}
}

func n_k_scope(a []int, n int, k int) ([]int, bool) {
	if n < k || n <= 0 {
		return nil, false
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
		return a[left:], true
	} else if k < n-left {
		return n_k_scope(a[left+1:], len(a[left+1:]), k)
	} else {
		result := make([]int, len(a[left:]))
		copy(result, a[left:])
		additional, _ := n_k_scope(a[:left], len(a[:left]), k-(n-left))
		return append(result, additional...), true
	}
}

func main() {
	fmt.Println("Welcome to the playground!")

	inputSlice := []int{1, 3, 5, 3, 2, 4, 6, 8, 4, 5, 7}
	k, _ := n_k_item(inputSlice, len(inputSlice), 4)
	fmt.Println("Result for item:", k)

	kscope, _ := n_k_scope(inputSlice, len(inputSlice), 4)
	fmt.Println("Result for scope:", kscope)
}

// Welcome to the playground!
// Result for item: 5
// Result for scope: [8 7 6 5]