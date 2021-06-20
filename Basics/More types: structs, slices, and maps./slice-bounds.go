// https://go-tour-jp.appspot.com/moretypes/10
package main

import "fmt"

// スライスするときは↓のような書き方で上限と下限を省略できる
func main() {
	s := []int{2, 3, 4, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}