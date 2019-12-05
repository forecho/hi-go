package main

import "math/rand"

import "fmt"

func main() {
	var count = 0
	// 等同于 count :=0
	for count < 10 {
		var num = rand.Intn(10) + 1
		fmt.Println(num)

		count++
	}

}
