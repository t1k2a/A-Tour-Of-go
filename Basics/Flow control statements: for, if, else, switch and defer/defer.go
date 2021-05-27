//https://go-tour-jp.appspot.com/flowcontrol/12
package main

import "fmt"

// deferステートメントで呼び出し元の関数の終わりまで遅延させる
// deferへ渡した関数の引数は、すぐに評価されるが、
// その関数自体は呼び出し元の関数がreturnするまで実行されない
// この場合はmainが終わるまでdeferは実行されない
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}