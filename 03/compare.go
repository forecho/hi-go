package main

import "fmt"

func main() {

	fmt.Println("There is a sign near the entrance that reads 'No Minors'.")

	var age = 41
	var minor = age < 18

	fmt.Printf("At age %v, am I a minor? %v\n", age, minor)
}

// Go 只提供了一个相等运算符， 并且它不允许直接比较文本和数字
