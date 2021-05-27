// https://go-tour-jp.appspot.com/flowcontrol/5
package main

import (
	"fmt"
	"math"
)

// ifも括弧は不要で中括弧は必要
// Goではおなじみらしい
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}