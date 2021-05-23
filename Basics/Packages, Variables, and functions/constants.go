package main

import "fmt"

// 定数は:=を使って宣言できない
const Pi = 3.14

func main() {
	const world = "世界"
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}