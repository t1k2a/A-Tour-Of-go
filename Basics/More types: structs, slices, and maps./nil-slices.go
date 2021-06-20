// https://go-tour-jp.appspot.com/moretypes/12
package main

import "fmt"

// スライスのゼロ値はnil
func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}