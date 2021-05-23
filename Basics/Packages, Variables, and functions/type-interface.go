package main

import "fmt"

func main() {
	v := 42 // 数値を変えてみる
	// v := 43.312
	v := 1.234+9.87i
	fmt.Printf("v is of type %T\n", v)
}