package main

import (
	"fmt"
)

const (
	float32_precision = 0.0000001
)

func square(num uint32, precision uint32) float32 {
	// get the intger
	var i float32 = 0

	for i*i <= float32(num) {
		if i*i == float32(num) {
			return i
		}
		i = i + 1
	}
	i = float32(i - 1)

	// calc the precision
	var multi float32 = 1
	for precision > 0 {
		precision--
		multi *= 10
		unit := 1 / multi

		for i*i <= float32(num) {
			if i*i >= float32(num) {
				return i
			}
			i = i + unit
		}
		i = i - unit

		if unit <= float32_precision {
			break
		}
	}

	return i
}

func main() {
	fmt.Println("Welcome to the playground!")

	num := uint32(2)

	fmt.Println("Result for item position:", square(num, uint32(7)))
}
