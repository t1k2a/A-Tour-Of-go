// https://go-tour-jp.appspot.com/flowcontrol/2
package main

import "fmt"

func main() {
	sum := 1
	// 初期化と後処理ステートメントは記述しなくてもOK
	// whileっぽい？
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}