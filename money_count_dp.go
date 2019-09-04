package main

import (
	"fmt"
)

func MoneyCount(options []int, total int) []int {
	totalRecords := make(map[int]bool, 0)
	itemRecords := make(map[int][]int, 0)
	optionLen := len(options)

	// init the records with only one item
	for i := 0; i < optionLen; i++ {
		itemRecords[options[i]] = append(itemRecords[options[i]], options[i])
		totalRecords[options[i]] = true
		if options[i] == total {
			return itemRecords[options[i]]
		}
	}

	calcCount := 0
	for true {
		tmpRecords := make(map[int]bool, 0)
		isEnd := true
		for current, _ := range totalRecords {
			for i := 0; i < optionLen; i++ {
				calcCount++
				newTotal := current + options[i]

				if newTotal > total {
					break
				}

				if _, ok := totalRecords[newTotal]; ok {
				} else {
					isEnd = false
					tmpRecords[newTotal] = true
					itemRecords[newTotal] = make([]int, len(itemRecords[current]))
					copy(itemRecords[newTotal], itemRecords[current])
					itemRecords[newTotal] = append(itemRecords[newTotal], options[i])
				}

				if newTotal == total {
					isEnd = true
					break
				}
			}
		}

		// merge two records
		for k, v := range tmpRecords {
			totalRecords[k] = v
		}

		if isEnd {
			break
		}
	}

	fmt.Println("Debug calc count:", calcCount)
	return itemRecords[total]
}

func GetSliceMin(status []int) int {
	min := status[0]
	for _, v := range status {
		if v < min {
			min = v
		}
	}

	return min
}

func MoneyCountDp(options []int, total int) []int {
	status := make([]int, total+1)
	for i := 0; i < total+1; i++ {
		status[i] = -1
	}

	for _, option := range options {
		if option < total+1 {
			status[option] = 1
		}
	}

	calcCount := 0
	for i := 0; i <= total; i++ {
		if status[i] != -1 {
			continue
		} else {
			tmpStatus := make([]int, 0)
			for _, option := range options {
				calcCount++
				if i < option {
					if len(tmpStatus) == 0 {
						status[i] = 0
					}
					break
				} else if i == option {
					status[i] = 1
					break
				} else {
					if status[i-option] > 0 {
						tmpStatus = append(tmpStatus, status[i-option])
					}
				}
			}

			if len(tmpStatus) > 0 {
				status[i] = 1 + GetSliceMin(tmpStatus)
			} else {
				status[i] = 0
			}
		}
	}

	// inverse for detail result
	if status[total] > 0 {
		result := make([]int, 0)
		for i, j := total, total; j >= 0; j-- {
			if status[i] == status[j]+1 {
				result = append(result, i-j)
				i = j
			}
			if status[j] == 1 {
				result = append(result, j)
				break
			}
		}

		fmt.Println("Debug calc count:", calcCount)
		return result
	}

	fmt.Println("Debug calc count:", calcCount)
	return nil
}

func main() {
	fmt.Println("Welcome to the playground!")

	options := []int{1, 2, 5, 10, 20, 50}
	total := 99

	options2 := []int{3, 7}
	total2 := 5

	fmt.Println("Search the money count for total:", total, "get result:", MoneyCount(options, total))
	fmt.Println("Search the money count for total2:", total2, "get result:", MoneyCount(options2, total2))

	fmt.Println("Search the money count for total:", total, "get result:", MoneyCountDp(options, total))
	fmt.Println("Search the money count for total2:", total2, "get result:", MoneyCountDp(options2, total2))
}

// Welcome to the playground!
// Debug calc count: 2029
// Search the money count for total: 99 get result: [2 2 20 20 5 50]
// Debug calc count: 2
// Search the money count for total2: 5 get result: []
// Debug calc count: 536
// Search the money count for total: 99 get result: [2 2 5 20 20 50]
// Debug calc count: 7
// Search the money count for total2: 5 get result: []
