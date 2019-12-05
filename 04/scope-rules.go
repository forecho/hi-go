package main

import "math/rand"

import "fmt"

var era = "AD" // era 变量在整个包都是可用的。

func main() {
	year := 2018
	// switch month := rand.Intn(12) + 1; month {
	// case 2:
	// 	day := rand.Intn(28) + 1
	// 	fmt.Println(era, year, month, day)
	// case 4, 6, 9, 11:
	// 	day := rand.Intn(30) + 1
	// 	fmt.Println(era, year, month, day)
	// default:
	// 	day := rand.Intn(31) + 1
	// 	fmt.Println(era, year, month, day)
	// }

	month := rand.Intn(12) + 1
	daysMonth := 31

	switch month {
	case 2:
		daysMonth = 28
	case 4, 6, 9, 11:
		daysMonth = 30
	}

	day := rand.Intn(daysMonth) + 1
	fmt.Println(era, year, month, day)
}
