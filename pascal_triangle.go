package main

import (
	"fmt"
)

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func LevelSearch(matrix [][]int) int {
	levels := len(matrix)
	for level := levels - 1; level > 0; level-- {
		for item := 0; item < len(matrix[level])-1; item++ {
			matrix[level-1][item] += min(matrix[level][item], matrix[level][item+1])
		}
	}

	return matrix[0][0]
}

func main() {
	fmt.Println("Welcome to the playground!")
	matrix := [][]int{
		{5},
		{7, 8},
		{2, 3, 4},
		{4, 9, 6, 1},
		{2, 7, 9, 4, 5}}
	fmt.Println("Search the shortest path via level, get result:", LevelSearch(matrix))
}

// Welcome to the playground!
// Debug level: 4 item: 0
// Debug level: 4 item: 1
// Debug level: 4 item: 2
// Debug level: 4 item: 3
// Debug level: 3 item: 0
// Debug level: 3 item: 1
// Debug level: 3 item: 2
// Debug level: 2 item: 0
// Debug level: 2 item: 1
// Debug level: 1 item: 0
// Search the shortest path via level, get result: 20