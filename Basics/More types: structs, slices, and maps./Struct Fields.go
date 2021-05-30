package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// structのフィールドは、ドット（.）を用いてアクセスする
func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}