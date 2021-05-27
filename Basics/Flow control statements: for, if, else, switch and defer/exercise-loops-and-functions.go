// https://go-tour-jp.appspot.com/flowcontrol/8
// 初課題
package main

import (
	"fmt"
	"math"
)

// 引数の平方根に近づける
func Sqrt(x float64) float64 {
	z := 1.0

	// 10回繰り返すことが要件となっている
	for i:=1; i<=10; i++ {

		// 記載通りの計算式
		z -= (z*z - x) / (2*z)

		count := i  
		countString := "回目の結果"
		result := z
		fmt.Println(count , countString , result)

		// 実際の平方根と同じになれば抜ける
		if math.Sqrt(x) == z {
			fmt.Println("clear!")
			break;
		}
	}

	return z
}

func main() {
	var num float64 = 2 // 別の数値を試したい時はここを変更する
	result := Sqrt(num)	
	fmt.Println(result)
	fmt.Println(math.Sqrt(num))
}