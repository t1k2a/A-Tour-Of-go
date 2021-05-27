//https://go-tour-jp.appspot.com/flowcontrol/10
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	// 曜日取得
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today+1:
		fmt.Println("Tommorow.")
	case today+2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away")
	}
}
