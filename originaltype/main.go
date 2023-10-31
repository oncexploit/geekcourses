package main

import (
	"fmt"
	"math"
	"originaltype/pkg"
)

func main() {
	//int,int8,int32,int64 (uint)
	//uintptr  大到可存储任意指针
	overchar()
}
func overchar() {
	presentvar()
	showfloat()
	pkg.Toseeistobelieve()
	pkg.Digintostring()
	pkg.Aboutrune()
	pkg.Main()
	pkg.Iteratechar()
	pkg.Addstring()
	pkg.Comparestring()
	pkg.ConvertString()
}

func presentvar() {
	var a int8 = 59
	fmt.Printf("%b\n", a) //输出二进制：111011
	fmt.Printf("%d\n", a) //输出十进制：59
	fmt.Printf("%o\n", a) //输出八进制：73
	fmt.Printf("%O\n", a) //输出八进制(带0o前缀)：0o73
	fmt.Printf("%x\n", a) //输出十六进制(小写)：3b
	fmt.Printf("%X\n", a) //输出十六进制(大写)：3B
}

func showfloat() {
	var f float32 = 139.8125
	bits := math.Float32bits(f)
	fmt.Printf("%b\n", bits)
}

func _present() {
	/*	a := 53         // 十进制
		b := 0700        // 八进制，以"0"为前缀
		c1 := 0xaabbcc   // 十六进制，以"0x"为前缀
		c2 := 0xddeeff   // 十六进制，以"0X"为前缀
		d1 := 0b10000001 // 二进制，以"0b"为前缀
		d2 := 0b10000001 // 二进制，以"0B"为前缀
		e1 := 0o700      // 八进制，以"0o"为前缀
		e2 := 0o700      // 八进制，以"0O"为前缀
	*/
}
