// https://go-tour-jp.appspot.com/moretypes/1
package main

import "fmt"

// Goではポインタを使うらしい。ポインタは値のメモリアドレスを示すとのこと
// C言語とは異なり、ポインタ演算はない
func main() {
	i, j := 42, 2701

	p := &i				// ポインタ引き出し
	fmt.Println(*p) 	// ポインタpを通してiから値を読み出す
	*p = 21				// ポインタpを通してiへ値を代入する
	fmt.Println(i)		// ポインタpを通してiを読み出す
	
	p = &j				// ポインタ引き出し
	*p = *p / 37		// ポインタpを通してjの計算をする
	fmt.Println(j)		// ポインタpを通してjを読み出す
}