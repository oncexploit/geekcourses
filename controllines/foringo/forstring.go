package foringo

import "fmt"

func showusages() {
	var s = "中国人"
	for i, v := range s { //这种用法中的v是 rune 类型
		fmt.Printf("%d %s 0x%x\n", i, string(v), v)
	}
}
