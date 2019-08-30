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
	states := [weightLimitation + 1]ValueItemRecord{}
	for i := 0; i < len(states); i++ {
		states[i].value = -1
	}

	// special value
	states[0] = ValueItemRecord{0, nil}

	// simple processing for the first item
	if weights[0] <= weightLimitation {
		states[weights[0]] = ValueItemRecord{values[0], []int{0}}
	}

	for i := 1; i < countLimitation; i++ {
		for j := weightLimitation - weights[i]; j >= 0; j-- {
			if states[j].value >= 0 {
				v := states[j].value + values[i]
				if v >= states[j+weights[i]].value {
					if v > maxValue {
						maxValue = v
					}
					states[j+weights[i]].value = v
					states[j+weights[i]].items = append(states[j].items, i)
					fmt.Println("Current packages:", states[j+weights[i]].items, ", with value:", v, ", with weight", j+weights[i])
				}
			}
		}
	}

	for j := weightLimitation; j >= 0; j-- {
		if states[j].value == maxValue {
			fmt.Println("Debug max value:", maxValue, "with packages:", states[j].items, "with weight:", j)
			resultPackages = states[j].items
		}
	}
}

func main() {
	fmt.Println("Welcome to the playground!")
	Package()
	fmt.Println("Most valueable items with value:", maxValue, "and package index:", resultPackages)
}

// Welcome to the playground!
// Current packages: [0 1] , with value: 7 , with weight 4
// Current packages: [1] , with value: 4 , with weight 2
// Current packages: [0 1 2] , with value: 15 , with weight 8
// Current packages: [1 2] , with value: 12 , with weight 6
// Current packages: [2] , with value: 8 , with weight 4
// Current packages: [1 2 3] , with value: 21 , with weight 12
// Current packages: [2 3] , with value: 17 , with weight 10
// Current packages: [0 1 2 4] , with value: 21 , with weight 11
// Current packages: [1 2 4] , with value: 18 , with weight 9
// Current packages: [2 4] , with value: 14 , with weight 7
// Current packages: [1 4] , with value: 10 , with weight 5
// Current packages: [4] , with value: 6 , with weight 3
// Debug max value: 21 with packages: [1 2 3] with weight: 12
// Debug max value: 21 with packages: [0 1 2 4] with weight: 11
// Most valueable items with value: 21 and package index: [0 1 2 4]