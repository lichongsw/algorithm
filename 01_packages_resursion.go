package main

import (
	"fmt"
)

const (
	productTotalCnt  = 5  // 物品总数
	weightLimitation = 12 // 背包承受的最大重量
)

var weights = []int{2, 2, 4, 6, 3} // 物品的重量
var values = []int{3, 4, 8, 9, 6}  // 物品的价值

var maxValue int = 0
var maxValueWeight int = 0
var resultPackages = []int{}
var recursionCnt = 0
var unpackFirst = false // 控制递归是优先装物品还是优先不装物品（true优先不装，代码更简洁，但流程理解烧脑一点。优先装上更符合人的思维展开方式）

func updatePackage(packages []int) {
	resultPackages = resultPackages[:0]
	resultPackages = append(resultPackages, packages...)
	fmt.Println("Update packages:", packages, "with value:", maxValue, "with weight:", maxValueWeight)
}

func Package(packages []int, productIndex int, currentWeight int, currentValue int) {
	recursionCnt++
	if productIndex >= productTotalCnt {
		return
	}

	// 这段临时属性为了让使用者更好的理解代码递归过程
	initWeight := currentWeight
	initValue := currentValue
	initPackage := []int{}
	initPackage = append(initPackage, packages...)

	// 1. 不装当前物品，再处理下一个物品
	// 1.1 不装不改变任何属性，可以省掉临时属性的开销（例如变种情况可能会有选择放弃不装当前物品的次数限制）
	// 1.2 递归处理下一个物品
	if unpackFirst {
		Package(packages, productIndex+1, currentWeight, currentValue)
		// 使用临时属性也ok
		// Package(initPackage, productIndex+1, initWeight, initValue)
	}

	// 2. 尝试装上当前物品，再处理下一个物品
	if currentWeight+weights[productIndex] <= weightLimitation {
		// 2.1 装上物品后最大价值下最有重量的属性更新
		packages = append(packages, productIndex)
		fmt.Println("Do pack success:", packages, "with new index:", productIndex, "having toatal weight:", currentWeight+weights[productIndex])
		currentWeight += weights[productIndex]
		currentValue += values[productIndex]
		if currentValue == maxValue {
			if currentWeight < maxValueWeight {
				// 同等价值下重量轻
				maxValueWeight = currentWeight
				updatePackage(packages)
			}
		} else if currentValue > maxValue {
			// 更高价值（没有超重）
			maxValue = currentValue
			maxValueWeight = currentWeight
			updatePackage(packages)
		}

		// 2.2. 递归处理下一个物品
		Package(packages, productIndex+1, currentWeight, currentValue)
	} else {
		fmt.Println("Do pack failed due to weight limitation", packages, "with index:", productIndex, "current total weight:", currentWeight, "pack weight:", weights[productIndex])
	}

	// 不装与装的顺序是没有先后之分的，是两个同等的分支。就是要再选择之前保存一下状态，避免属性干扰，所以这个例子里面优先不装的代码更简洁
	if !unpackFirst {
		Package(initPackage, productIndex+1, initWeight, initValue)
	}
}

func main() {
	Package([]int{}, 0, 0, 0)
	fmt.Println("Most valueable items with value:", maxValue, "and weight:", maxValueWeight, "and index:", resultPackages)
	fmt.Println("Recursion count:", recursionCnt)
}

// Current packages: [4] with value: 6 with weight: 3
// Current packages: [3] with value: 9 with weight: 6
// Current packages: [3 4] with value: 15 with weight: 9
// Current packages: [2 3] with value: 17 with weight: 10
// Current packages: [1 3 4] with value: 19 with weight: 11
// Current packages: [1 2 3] with value: 21 with weight: 12
// Current packages: [0 1 2 4] with value: 21 with weight: 11
// Most valueable items with value: 21 and index: [0 1 2 4]
