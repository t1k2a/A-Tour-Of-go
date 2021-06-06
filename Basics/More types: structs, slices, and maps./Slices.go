// https://go-tour-jp.appspot.com/moretypes/7
package main

import "fmt"

// 変数名[n:n]でスライス配列　配列とは異なり、可変長となる
// スライス配列は割と一般的らしい
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}