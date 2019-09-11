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

func LongestSubString(first string, second string) (int, string) {
	// status init
	status := make([][]int, len(first))
	for i := 0; i < len(first); i++ {
		status[i] = make([]int, len(second))
	}

	// first row and column
	for i := 0; i < len(second); i++ {
		if first[0] == second[i] {
			status[0][i] = 1
		}
	}

	for i := 0; i < len(first); i++ {
		if second[0] == first[i] {
			status[i][0] = 1
		}
	}

	maxStringLen := 0
	index := 0

	for i := 1; i < len(first); i++ {
		for j := 1; j < len(second); j++ {
			if first[i] != second[j] {
				status[i][j] = 0
			} else {
				status[i][j] = status[i-1][j-1] + 1
				if status[i][j] > maxStringLen {
					maxStringLen = status[i][j]
					index = i
				}
			}
		}
	}

	s := string([]byte(first)[index+1-maxStringLen : index+1])
	return maxStringLen, s
}

func LongestSubSequence(first string, second string) (int, string) {
	// status init
	status := make([][]int, len(first))
	for i := 0; i < len(first); i++ {
		status[i] = make([]int, len(second))
	}

	// first row and column
	matched := false
	for i := 0; i < len(second); i++ {
		if first[0] == second[i] {
			status[0][i] = 1
			matched = true
		}
		if matched {
			status[0][i] = 1
		}
	}

	matched = false
	for i := 0; i < len(first); i++ {
		if second[0] == first[i] {
			status[i][0] = 1
			matched = true
		}
		if matched {
			status[i][0] = 1
		}
	}

	maxStringLen := 0
	index := 0

	// dp equation
	for i := 1; i < len(first); i++ {
		for j := 1; j < len(second); j++ {
			if first[i] == second[j] {
				status[i][j] = status[i-1][j-1] + 1
			} else {
				status[i][j] = GetMax(status[i-1][j], status[i][j-1])
			}

			if status[i][j] > maxStringLen {
				maxStringLen = status[i][j]
				index = i
			}
		}
	}

	result := make([]byte, 0)
	index = 0
	for i := 0; i < len(first); i++ {
		if status[i][len(second)-1] > index {
			result = append(result, first[i])
			index++
		}
	}

	return maxStringLen, string(result)
}

func main() {
	fmt.Println("Welcome to the playground!")

	firstString := "bdefgih"
	secondString := "abcefgh"

	length, substring := LongestSubSequence(firstString, secondString)
	fmt.Println("Search longest sub string for", firstString, "and", secondString, ",result:", substring, "with length:", length)

	length, substring = LongestSubString(firstString, secondString)
	fmt.Println("The better search longest sub string for", firstString, "and", secondString, ",result:", substring, "with length:", length)
}

// Welcome to the playground!
// Search longest sub string for bdefgih and abcefgh ,result: befgh with length: 5
// The better search longest sub string for bdefgih and abcefgh ,result: efg with length: 3
