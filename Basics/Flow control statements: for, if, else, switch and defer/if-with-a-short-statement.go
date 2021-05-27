// https://go-tour-jp.appspot.com/flowcontrol/7
package main

import (
	"fmt"
	"math"
)

// if内でforのように初期化できるらしい
// スコープ範囲がif内になるのでif外では初期化された変数は使えないぜ
func pow(x, n, lim float64) float64 {
	// mathクラスのpow関数でべき乗する
	if v := math.Pow(x, n); v <lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 2, 20),
		pow(3, 2, 8), // 第３引数が10未満なら第３引数を返す
	)
}