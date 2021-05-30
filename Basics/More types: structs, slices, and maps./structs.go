// https://go-tour-jp.appspot.com/moretypes/2
package main
import "fmt"

// sturuct（構造体）は、フィールド（field）の集まり
// メンバ変数的なものの集まり
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}