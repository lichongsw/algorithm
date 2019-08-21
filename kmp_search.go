package main

import (
	"fmt"
)

func GetNext(source string) []int {
	sourceLen := len(source)
	next := make([]int, sourceLen)

	for i := 0; i < sourceLen; i++ {
		if i == 0 {
			next[i] = 0
		} else {
			if source[i] == source[next[i-1]] {
				next[i] = next[i-1] + 1
			} else {
				next[i] = 0
			}
		}
	}

	return next
}

func KMPSearch(source string, search string) (int, bool) {
	sourceIndex, searchIndex := 0, 0
	sourceLen := len(source)
	searchLen := len(search)
	if searchLen > sourceLen {
		return -1, false
	}

	next := GetNext(source)
	fmt.Println("Assistant array:", next)
	for sourceIndex < sourceLen {
		if source[sourceIndex] == search[searchIndex] {
			fmt.Println("Debug match index:", sourceIndex, searchIndex)
			if searchIndex == searchLen-1 {
				return sourceIndex + 1 - searchLen, true
			}
			sourceIndex++
			searchIndex++
		} else {
			fmt.Println("Debug mismatch index:", sourceIndex, searchIndex)
			if searchIndex > next[sourceIndex-1] {
				searchIndex = next[sourceIndex-1]
			} else {
				searchIndex = next[sourceIndex]
				sourceIndex++
			}
		}
	}

	return -1, false
}

func main() {
	fmt.Println("Welcome to the playground!")
	sourceString := "abcdeabcdefgabcd"
	searchString := "abcdefg"
	searchString2 := "abcdefgh"

	index, ok := KMPSearch(sourceString, searchString)
	fmt.Println("Find sub string", ok, "with index:", index)

	index, ok = KMPSearch(sourceString, searchString2)
	fmt.Println("Find sub string", ok, "with index:", index)
}

// Welcome to the playground!
// Assistant array: [0 0 0 0 0 1 2 3 4 5 0 0 1 2 3 4]
// Debug match index: 0 0
// Debug match index: 1 1
// Debug match index: 2 2
// Debug match index: 3 3
// Debug match index: 4 4
// Debug mismatch index: 5 5
// Debug match index: 5 0
// Debug match index: 6 1
// Debug match index: 7 2
// Debug match index: 8 3
// Debug match index: 9 4
// Debug match index: 10 5
// Debug match index: 11 6
// Find sub string true with index: 5
// Assistant array: [0 0 0 0 0 1 2 3 4 5 0 0 1 2 3 4]
// Debug match index: 0 0
// Debug match index: 1 1
// Debug match index: 2 2
// Debug match index: 3 3
// Debug match index: 4 4
// Debug mismatch index: 5 5
// Debug match index: 5 0
// Debug match index: 6 1
// Debug match index: 7 2
// Debug match index: 8 3
// Debug match index: 9 4
// Debug match index: 10 5
// Debug match index: 11 6
// Debug mismatch index: 12 7
// Debug match index: 12 0
// Debug match index: 13 1
// Debug match index: 14 2
// Debug match index: 15 3
// Find sub string false with index: -1