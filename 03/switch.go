package main

import "fmt"

func main() {
	var room = "lake"

	switch {
	case room == "cave":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case room == "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough
	case room == "underwater":
		fmt.Println("The water is freezing cold.")
	}
}

// 在 C、Java、JavaScript 等语言中， 下降是 switch 语句各个分支的默认行为。 Go 对此采取了更为谨慎的做法， 用户需要显式地使用 fallthrough 关键字才会引发下降。
