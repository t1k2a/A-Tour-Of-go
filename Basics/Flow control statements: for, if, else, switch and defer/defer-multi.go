//https://go-tour-jp.appspot.com/flowcontrol/13
package main

import "fmt"

// deferが複数ある場合、スタックされる
// スタックされたものは後入れ先出しになる
// この場合は数字の大きい順に出力される
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

//詳しくはこちら https://blog.golang.org/defer-panic-and-recover