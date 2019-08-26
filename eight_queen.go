package main

import (
	"fmt"
)

func IsPositionOK(result [][]int, xPos int, yPos int) bool {
	for i := 0; i < xPos; i++ {
		y := GetYPosition(result, i)
		if y == yPos || i+y == xPos+yPos || i+yPos == y+xPos {
			return false
		}
	}
	return true
}

func GetYPosition(result [][]int, xPos int) (y int) {
	for yPos := 0; yPos < 8; yPos++ {
		if result[xPos][yPos] == 1 {
			return yPos
		}
	}
	return -1
}

func EightQueen() ([][]int, bool) {
	foundResult := true
	result := make([][]int, 8)
	for i := 0; i < 8; i++ {
		result[i] = make([]int, 8)
	}

	for x := 0; x < 8; {
		yPos := GetYPosition(result, x)
		if yPos != -1 {
			result[x][yPos] = 0
			yPos += 1
		} else {
			yPos = 0
		}

		if yPos == 8 {
			x--
		} else {
			for y := yPos; y < 8; y++ {
				if IsPositionOK(result, x, y) {
					result[x][y] = 1
					x++
					break
				} else {
					if y == 7 {
						x--
					}
				}
			}
		}

		if x < 0 {
			foundResult = false
			break
		}
	}

	return result, foundResult
}

func Recursion(result [][]int, x int) {
	if x < 0 || x > 7 {
		return
	}

	yPos := GetYPosition(result, x)
	if yPos != -1 {
		result[x][yPos] = 0
		yPos += 1
	} else {
		yPos = 0
	}

	if yPos == 8 {
		Recursion(result, x-1)
	}

	for y := yPos; y < 8; y++ {
		if IsPositionOK(result, x, y) {
			result[x][y] = 1
			Recursion(result, x+1)
			break
		} else {
			if y == 7 {
				Recursion(result, x-1)
			}
		}
	}
}

func EightQueenRecursion() ([][]int, bool) {
	foundResult := true
	result := make([][]int, 8)
	for i := 0; i < 8; i++ {
		result[i] = make([]int, 8)
	}

	Recursion(result, 0)

	return result, foundResult
}

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("Eight queen normal loop solution result:")
	if result, ok := EightQueen(); ok {
		for i := 0; i < len(result); i++ {
			fmt.Println(result[i])
		}
	}

	fmt.Println("Eight queen recursion solution result:")
	if result2, ok := EightQueenRecursion(); ok {
		for i := 0; i < len(result2); i++ {
			fmt.Println(result2[i])
		}
	}
}

// Welcome to the playground!
// Eight queen normal loop solution result:
// [1 0 0 0 0 0 0 0]
// [0 0 0 0 1 0 0 0]
// [0 0 0 0 0 0 0 1]
// [0 0 0 0 0 1 0 0]
// [0 0 1 0 0 0 0 0]
// [0 0 0 0 0 0 1 0]
// [0 1 0 0 0 0 0 0]
// [0 0 0 1 0 0 0 0]
// Eight queen recursion solution result:
// [1 0 0 0 0 0 0 0]
// [0 0 0 0 1 0 0 0]
// [0 0 0 0 0 0 0 1]
// [0 0 0 0 0 1 0 0]
// [0 0 1 0 0 0 0 0]
// [0 0 0 0 0 0 1 0]
// [0 1 0 0 0 0 0 0]
// [0 0 0 1 0 0 0 0]