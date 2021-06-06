// https://go-tour-jp.appspot.com/moretypes/6
package main

import "fmt"

func main() {
	// 配列の長さも含めて型　よって後から変えられない
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}