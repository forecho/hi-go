package main

import "fmt"

import "math/rand"

func main() {
	var degrees = 0
	for {
		fmt.Println(degrees)

		degrees++
		if degrees >= 360 {
			degrees = 0
			if rand.Intn(2) == 0 {
				break
			}
		}
	}
}

// 我们还可以通过不为 for 语句设置任何条件来产生无限循环，然后在有需要的时候通过在循环体内使用 break 语句来跳出循环。
