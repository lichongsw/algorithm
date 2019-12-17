package main

import (
	"bytes"
	"flag"
	"fmt"
)

func getCountFromNum(inputNumber int) int {
	num := inputNumber - '0'
	if num >= 2 && num <= 6 {
		return 3
	}

	if num == 7 {
		return 4
	} else if num == 8 {
		return 3
	} else if num == 9 {
		return 4
	} else {
		return 0
	}
}

func getStringFromNum(inputNumber int) []string {
	num := inputNumber - '0'
	if num >= 2 && num <= 6 {
		return []string{string('a' + 3*(num-2)),
			string('b' + 3*(num-2)),
			string('c' + 3*(num-2))}
	}

	if num == 7 {
		return []string{string('p'), string('q'), string('r'), string('s')}
	} else if num == 8 {
		return []string{string('t'), string('u'), string('v')}
	} else if num == 9 {
		return []string{string('w'), string('x'), string('y'), string('z')}
	} else {
		return nil
	}
}

func combination(a []string, b []string) []string {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}

	result := make([]string, 0, len(a)*len(b))

	var buffer bytes.Buffer
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			buffer.Reset()
			buffer.WriteString(a[i])
			buffer.WriteString(b[j])
			result = append(result, buffer.String())
		}
	}

	return result
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	} else if len(digits) == 1 {
		return getStringFromNum(int(digits[0]))
	} else {
		return combination(getStringFromNum(int(digits[0])), letterCombinations(digits[1:]))
	}
}

var s string

func init() {
	flag.StringVar(&s, "s", "", "input numbers as a string")
}

func main() {
	flag.Parse()
	fmt.Println("Welcome to the playground!")
	fmt.Println(s)
	//inputString := "22"

	fmt.Println("Result:", letterCombinations(s))
}
