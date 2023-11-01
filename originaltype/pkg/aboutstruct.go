package pkg

import "unsafe"

//零值可用，字面值初始化，内存对齐

type T struct {
	b byte
	i int64
	u uint16
}

type S struct {
	b byte
	u uint16
	i int64
}

func main() {
	var t T
	println(unsafe.Sizeof(t)) // 24
	var s S
	println(unsafe.Sizeof(s)) // 16
}
