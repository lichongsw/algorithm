package main

import (
	"fmt"
)

func binarySearch(arr []int, val int) int {
	first := 0
	last := len(arr) - 1

	for first <= last {
		mid := first + (last-first)>>1
		switch {
		case arr[mid] == val:
			return mid
		case arr[mid] < val:
			first = mid + 1
		case arr[mid] > val:
			last = mid - 1
		}
	}

	// not exist
	return -1
}

func binarySearchByRecursion(arr []int, val int, start int) int {
	first := 0
	last := len(arr) - 1

	if first <= last {
		mid := first + (last-first)>>1
		switch {
		case arr[mid] == val:
			return mid + start
		case arr[mid] < val:
			return binarySearchByRecursion(arr[mid+1:], val, mid+1)
		case arr[mid] > val:
			return binarySearchByRecursion(arr[:mid], val, start)
		}
	}

	// not exist
	return -1
}

func main() {
	fmt.Println("Welcome to the playground!")

	inputSlice := []int{1, 3, 5, 6, 7, 9, 11, 12, 14, 15}
	searchItem := 9
	searchNotExistItem := 8
	fmt.Println("Result for item position:", binarySearch(inputSlice, searchNotExistItem))
	fmt.Println("Result for item position:", binarySearch(inputSlice, searchItem))

	fmt.Println("Result for item position:", binarySearchByRecursion(inputSlice, searchNotExistItem, 0))
	fmt.Println("Result for item position:", binarySearchByRecursion(inputSlice, searchItem, 0))
}
