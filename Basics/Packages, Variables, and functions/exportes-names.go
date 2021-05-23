package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
	fmt.Println(math.pi) // piの頭文字が小文字のため、mathパッケージからexportされたものではない
}