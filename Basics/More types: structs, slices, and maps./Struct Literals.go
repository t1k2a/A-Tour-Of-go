// https://go-tour-jp.appspot.com/moretypes/5
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2} // Vertexに値を入れる
	v2 = Vertex{X: 1} // 暗黙的にY：0を表現 goのintは初期値０
	v3 = Vertex{} // X:とY:が０
	p = &Vertex{1, 2} // Vertexのポインタに値を入れる
)

func main() {
	fmt.Println(v1, p, v2, v3)
}