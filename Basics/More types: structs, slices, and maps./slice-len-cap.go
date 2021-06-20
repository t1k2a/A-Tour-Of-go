// https://go-tour-jp.appspot.com/moretypes/11
package main

import "fmt"

// スライスは長さ（length）と容量(capacity)の両方を持っている
func main() {
	s := []int{2, 3, 4, 7, 11, 13}
	printSlice(s)

	// スライスに0個の長さを与える
	s = s[:0]
	printSlice(s)

	// 長さを拡張
	s = s[:4]
	printSlice(s)

	// 2個目までの要素を削除
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	// 長さと容量と値を表示
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}