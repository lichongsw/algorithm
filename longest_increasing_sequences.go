package main

import (
	"fmt"
)

func GetMax(first int, second int) int {
	max := first
	if max < second {
		max = second
	}

	return max
}

func LongestIncreasingSequence(numbers []int) (int, []int) {
	if len(numbers) == 0 {
		return 0,nil
	}

	status := make([]int, len(numbers))
	for i := 0; i < len(numbers); i++ {
		status[i] = 1
	}

	maxLength := 1
	for i := 0; i < len(numbers); i++ {
		for j := i - 1; j >= 0; j-- {
			if numbers[j] < numbers[i] {
				status[i] = GetMax(status[j]+1, status[i])
				if status[j] == maxLength || numbers[j] == numbers[i]-1 {
					break
				}
			}
		}
		fmt.Println("Debug i", i, "status[i]", status[i])
		if status[i] > maxLength {
			maxLength = status[i]
		}
	}

	result := make([]int, maxLength)
	recordLength := maxLength
	for i := len(numbers) - 1; i >= 0; i-- {
		if status[i] == recordLength {
			result[recordLength - 1] = numbers[i]
			recordLength--
		}
	}
	fmt.Println("Debug result", result)

	return maxLength, result
}

func main() {
	fmt.Println("Welcome to the playground!")

	numbers := []int{5, 6, 1, 3, 8, 9, 6, 7, 9, 4, 10, 5, 6}

	length, result := LongestIncreasingSequence(numbers)
	fmt.Println("The longest sequence for", numbers, "with length:", length, "and result:", result)
}

// Welcome to the playground!
// Debug i 0 status[i] 1
// Debug i 1 status[i] 2
// Debug i 2 status[i] 1
// Debug i 3 status[i] 2
// Debug i 4 status[i] 3
// Debug i 5 status[i] 4
// Debug i 6 status[i] 3
// Debug i 7 status[i] 4
// Debug i 8 status[i] 5
// Debug i 9 status[i] 3
// Debug i 10 status[i] 6
// Debug i 11 status[i] 4
// Debug i 12 status[i] 5
// Debug result [1 3 6 7 9 10]
// The longest sequence for [5 6 1 3 8 9 6 7 9 4 10 5 6] with length: 6 and result: [1 3 6 7 9 10]