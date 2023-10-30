package main

import (
	"fmt"
	"gobasic/pkg"
)

func main() {
	//do nothing for now
	vardeclare()
	pkg.Shadow()
	pkg.Bar()
	pkg.Main()
}

func vardeclare() {
	var (
		a       = 8
		b       = int32(13)
		c       = string("abc")
		e, f, g = 'C', 'D', 'E'
	)

	aa, bb, cc := 12, 'C', "asd"

	fmt.Println(a, b, c, e, f, g, aa, bb, cc)
}
