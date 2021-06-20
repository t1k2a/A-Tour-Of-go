// https://go-tour-jp.appspot.com/moretypes/9
package main

import "fmt"

// スライスのリテラルは長さの無い配列リテラルのようなもの
// [3]bool{true, true, false}←配列リテラル
// []bool{true, true, false}←スライスのリテラル
// リテラル＝人間後で書いたプログラムのソースコードの中に直接ベタ書きした文字とか数字とかのこと
func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{	
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}