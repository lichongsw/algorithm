package main

import (
	"fmt"
)

func min(x int, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func GetPosition(search string) map[byte]int {
	result := make(map[byte]int)
	searchLen := len(search)
	for i := 0; i < searchLen; i++ {
		if oldPosition, ok := result[search[i]]; ok {
			result[search[i]] = min(oldPosition, i)
		} else {
			result[search[i]] = i
		}
	}

	return result
}

func SundaySearch(source string, search string) (int, bool) {
	sourceIndex, searchIndex := 0, 0
	sourceLen := len(source)
	searchLen := len(search)
	if searchLen > sourceLen {
		return -1, false
	}

	searchPosition := GetPosition(search)
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
			endIndex := sourceIndex + searchLen - searchIndex
			if endIndex < sourceLen {
				if position, ok := searchPosition[source[endIndex]]; ok {
					sourceIndex = endIndex - position
				} else {
					sourceIndex = endIndex + 1
				}
			} else {
				break
			}
			searchIndex = 0
		}
	}

	return -1, false
}

func main() {
	fmt.Println("Welcome to the playground!")
	sourceString := "abcdeabcabcdefgabcd"
	searchString := "abcabc"
	searchString2 := "abcdefgh"

	index, ok := SundaySearch(sourceString, searchString)
	fmt.Println("Find sub string", ok, "with index:", index)

	index, ok = SundaySearch(sourceString, searchString2)
	fmt.Println("Find sub string", ok, "with index:", index)
}

// Welcome to the playground!
// Debug match index: 0 0
// Debug match index: 1 1
// Debug match index: 2 2
// Debug mismatch index: 3 3
// Debug match index: 5 0
// Debug match index: 6 1
// Debug match index: 7 2
// Debug match index: 8 3
// Debug match index: 9 4
// Debug match index: 10 5
// Find sub string true with index: 5
// Debug match index: 0 0
// Debug match index: 1 1
// Debug match index: 2 2
// Debug match index: 3 3
// Debug match index: 4 4
// Debug mismatch index: 5 5
// Debug match index: 8 0
// Debug match index: 9 1
// Debug match index: 10 2
// Debug match index: 11 3
// Debug match index: 12 4
// Debug match index: 13 5
// Debug match index: 14 6
// Debug mismatch index: 15 7
// Debug match index: 15 0
// Debug match index: 16 1
// Debug match index: 17 2
// Debug match index: 18 3
// Find sub string false with index: -1

