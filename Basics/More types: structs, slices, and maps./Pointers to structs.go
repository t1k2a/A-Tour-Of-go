// https://go-tour-jp.appspot.com/moretypes/4
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// structのポインタを通してアクセスすることもできる
func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)	
}