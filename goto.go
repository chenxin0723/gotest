package main

import (
	"fmt"
)

func myFunc() {
	i := 0
XinChen: //这行的第一个词,以冒号结束作为标签
	println(i)
	i++
	goto XinChen //跳转到 Here 去
}

type say func(name string)

func main() {

	gg := func(name string) {
		fmt.Println(name)
	}

	// s := func(lover string) string {
	// 	return "Xin Chen " + lover
	// }("wei dan dan")

	// println(s)
	// fmt.Printf("my name is %s, and your?", s)

}
