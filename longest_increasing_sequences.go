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

func LongestNumberSequence(numbers []int) (int) {
	if len(numbers) == 0 {
		return 0
	}

	status := make([]int, len(numbers))
	for i := 0; i < len(numbers); i++ {
		status[i] = 1
	}

	maxLength := 1
	for i := 1; i < len(numbers); i++ {
		for j := i - 1; j >= 0; j-- {
			if numbers[j] < numbers[i] {
				if (status[j] == maxLength || numbers[j] == numbers[i] - 1) {
					status[i] = status[j] + 1
					break
				}
				
				status[i] = GetMax(status[j] + 1, status[i])
			}
		}
		fmt.Println("Debug i", i, "status[i]", status[i])	
		if status[i] > maxLength {
			maxLength = status[i]
		}
	}

	return maxLength
}

func main() {
	fmt.Println("Welcome to the playground!")

	numbers := []int{10,9,2,5,3,7,101,18}

	length := LongestNumberSequence(numbers)
	fmt.Println("The longest sequence for", numbers, "with length:", length)
}

// Welcome to the playground!
// Debug i 1 status[i] 1
// Debug i 2 status[i] 1
// Debug i 3 status[i] 2
// Debug i 4 status[i] 2
// Debug i 5 status[i] 3
// Debug i 6 status[i] 4
// Debug i 7 status[i] 4
// The longest sequence for [10 9 2 5 3 7 101 18] with length: 4