package main

import (
	"flag"
	"fmt"
)

// func romanToInt(s string) int {
// 	result := 0
// 	for len(s) > 1 {
// 		//fmt.Println(len(s))
// 		if value, ok := symbol[s[:2]]; ok {
// 			result += value
// 			s = s[2:]
// 		} else {
// 			result += symbol[s[:1]]
// 			s = s[1:]
// 		}
// 	}

// 	if len(s)  > 0 {
// 		result += symbol[s[:1]]
// 	}
	
// 	return result
// }

func romanToInt(s string) int {
	result := 0
    chracters := 0
    for chracters < len(s) {
		if chracters < len(s) - 1 {
			if value, ok := symbol[s[chracters:chracters+2]]; ok {
				result += value
				chracters += 2
				continue
			}
		} 

		result += symbol[s[chracters:chracters+1]]
		chracters += 1
	}
	
	return result
}


var s string
var symbol map[string]int

func init() {
	flag.StringVar(&s, "s", "", "input string")
	symbol = make(map[string]int, 16)
	symbol["I"] = 1
	symbol["IV"] = 4
	symbol["V"] = 5
	symbol["IX"] = 9
	symbol["X"] = 10
	symbol["XL"] = 40	
	symbol["L"] = 50
	symbol["XC"] = 90	
	symbol["C"] = 100
	symbol["CD"] = 400	
	symbol["D"] = 500
	symbol["CM"] = 900	
	symbol["M"] = 1000
}

func main() {
	flag.Parse()
	fmt.Println("Welcome to the playground!")
	fmt.Println(s)


	fmt.Println(symbol)

	fmt.Println("Result:", romanToInt(s))
}