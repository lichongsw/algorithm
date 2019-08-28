package main

import (
	"fmt"
)

const (
	countLimitation  = 5  // 物品个数
	weightLimitation = 12 // 背包承受的最大重量
)

var weights = []int{2, 2, 4, 6, 3} // 物品的重量
var values = []int{3, 4, 8, 9, 6}  // 物品的价值
var maxValue int = 0
var maxValueWeight int = 0
var resultPackages = []int{}

type ValueItemRecord struct {
	value int
	items []int
}

func Package() {
	states := [countLimitation][weightLimitation + 1]ValueItemRecord{}
	// special value
	states[0][0] = ValueItemRecord{0, nil}

	// simple processing for the first item
	if weights[0] <= weightLimitation {
		states[0][weights[0]] = ValueItemRecord{values[0], []int{0}}
	}

	for i := 1; i < countLimitation; i++ {
		// do not choose the item i
		for j := 0; j <= weightLimitation; j++ {
			if states[i-1][j].value >= 0 {
				states[i][j] = states[i-1][j]
			}
		}

		// do choose the item i
		for j := 0; j <= weightLimitation-weights[i]; j++ {
			if states[i-1][j].value >= 0 {
				v := states[i-1][j].value + values[i]
				if v > states[i][j+weights[i]].value {
					states[i][j+weights[i]].value = v
					states[i][j+weights[i]].items = append(states[i-1][j].items, i)
					states[i-1][j].items = nil
					fmt.Println("Current packages:", states[i][j+weights[i]].items, "with value:", v)
				}
			}
		}
	}

	for j := 0; j <= weightLimitation; j++ {
		if states[countLimitation-1][j].value > maxValue {
			maxValue = states[countLimitation-1][j].value
			resultPackages = states[countLimitation-1][j].items
		}
	}
}

func main() {
	fmt.Println("Welcome to the playground!")
	Package()
	fmt.Println("Most valueable items with value:", maxValue, "and index:", resultPackages)
}

// Welcome to the playground!
// Current packages: [1] with value: 4
// Current packages: [1] with value: 4
// Current packages: [0 1] with value: 7
// Current packages: [1] with value: 4
// Current packages: [1] with value: 4
// Current packages: [1] with value: 4
// Current packages: [1] with value: 4
// Current packages: [1] with value: 4
// Current packages: [1] with value: 4
// Current packages: [1] with value: 4
// Current packages: [1] with value: 4
// Current packages: [2] with value: 8
// Current packages: [2] with value: 8
// Current packages: [1 2] with value: 12
// Current packages: [1 2] with value: 12
// Current packages: [0 1 2] with value: 15
// Current packages: [1 2] with value: 12
// Current packages: [1 2] with value: 12
// Current packages: [1 2] with value: 12
// Current packages: [1 2] with value: 12
// Current packages: [1 3] with value: 13
// Current packages: [2 3] with value: 17
// Current packages: [2 3] with value: 17
// Current packages: [1 2 3] with value: 21
// Current packages: [4] with value: 6
// Current packages: [1 4] with value: 10
// Current packages: [2 4] with value: 14
// Current packages: [1 4 4] with value: 18
// Current packages: [1 3 4] with value: 18
// Current packages: [0 1 2 4] with value: 21
// Most valueable items with value: 21 and index: [0 1 2 4]