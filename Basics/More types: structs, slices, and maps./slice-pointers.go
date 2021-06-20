// https://go-tour-jp.appspot.com/moretypes/8
package main

import "fmt"

func main() {
	// スライスを変更しようとするとその元となる配列の対応する要素が
	// 変更される
	// スライスは可変配列
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a,b)
	fmt.Println(names)
}
