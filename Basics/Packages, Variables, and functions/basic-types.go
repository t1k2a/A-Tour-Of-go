package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe	bool		= false
	MaxInt	uint64/* 64ビット符号なし整数 */		= 1<<64 -1
	z		complex128/* 複素数128ビット */	= cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}