package pkg

import "fmt"

var aa = 11

func foo(n int) {
	aa := 1
	aa += n
}

func Shadow() {
	fmt.Println(aa)
	foo(5)
	fmt.Println("after calling foo,aa = ", aa)
}

func Bar() {
	if a := 1; false { //等价于 a := 1
	} else if b := 2; false { // if false {...} else {...}
	} else if c := 3; false {
	} else {
		println(a, b, c)
	}
}
