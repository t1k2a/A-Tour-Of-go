// https://go-tour-jp.appspot.com/flowcontrol/3
package main

import "fmt"

func main() {
	sum := 1
	// セミコロンを省略できるらしい
	// 思った通りwhileのような振る舞いをするが、
	// Goではforだけを使うらしい
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}