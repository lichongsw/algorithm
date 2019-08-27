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

func Package(packages []int, count int, currentWeight int, currentValue int) {
	if currentWeight == weightLimitation || count == countLimitation {
		if currentValue == maxValue {
			if currentWeight < maxValueWeight {
				maxValueWeight = currentWeight
				resultPackages = nil
				resultPackages = append(resultPackages, packages...)
				fmt.Println("Current packages:", packages, "with value:", maxValue)
			}
		} else if currentValue > maxValue {
			maxValue = currentValue
			maxValueWeight = currentWeight
			resultPackages = nil
			resultPackages = append(resultPackages, packages...)
			fmt.Println("Current packages:", packages, "with value:", maxValue)
		}

		return
	}

	Package(packages, count+1, currentWeight, currentValue)
	if currentWeight+weights[count] <= weightLimitation {
		packages = append(packages, count)
		Package(packages, count+1, currentWeight+weights[count], currentValue+values[count])
	}
}

func main() {
	fmt.Println("Welcome to the playground!")
	Package([]int{}, 0, 0, 0)
	fmt.Println("Most valueable items with value:", maxValue, "and index:", resultPackages)
}

// Welcome to the playground!
// Current packages: [4] with value: 6
// Current packages: [3] with value: 9
// Current packages: [3 4] with value: 15
// Current packages: [2 3] with value: 17
// Current packages: [1 3 4] with value: 19
// Current packages: [1 2 3] with value: 21
// Current packages: [0 1 2 4] with value: 21
// Most valueable items with value: 21 and index: [0 1 2 4]
