package main

import "fmt"

func main() {
	count := 11
	for count = 10; count > 0; count-- {
		fmt.Println(count)
	}
	// 因为声明 count 变量是在循环体外，所以最后 count 还有效果
	fmt.Println(count)
}
