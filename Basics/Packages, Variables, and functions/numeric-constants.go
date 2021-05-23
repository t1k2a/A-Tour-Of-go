package main

import "fmt"

const(
	// 1の後に100個の0が続く
	Big = 1 << 100
	// もう一度右に９９桁シフトすると、1<<1つまり2になる
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1}
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// fmt.Println(needInt(Big)) // intの大きさを超えてエラーになる
}