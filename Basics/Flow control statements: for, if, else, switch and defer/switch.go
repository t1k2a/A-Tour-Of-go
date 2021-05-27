// https://go-tour-jp.appspot.com/flowcontrol/9
package main

import (
	"fmt"
	"runtime"
)

// Goのswitch文は自動的にbreakが入る
// caseは定数じゃなくても整数じゃなくてもOK
func main() {
	fmt.Print("Go run on ")
	// 使用PCのOSを取得 webのチュートリアルで実行するとLinuxになる
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X. ")
	case "linux":
		fmt.Println("Linux. ")
	default:
		// freebsd, openbsd
		// plam9, windows...
		fmt.Printf("%s. \n", os)
	}
}