package pkg

import (
	"fmt"
	"unsafe"
)

const arraylen = 10

var arr1 [arraylen]int
var arr2 [5]string

var mArr [2][3][4]int

func Printarr() {
	var arr = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("数组长度：", len(arr))           // 6
	fmt.Println("数组大小：", unsafe.Sizeof(arr)) // 48
}

func _initialarray() {
	var arr1 [6]int // [0 0 0 0 0 0]
	var arr2 = [6]int{
		11, 12, 13, 14, 15, 16,
	} // [11 12 13 14 15 16]
	fmt.Println(arr1)
	fmt.Println(arr2)

	var arr3 = [...]int{
		21, 22, 23,
	} // [21 22 23]
	fmt.Printf("%T\n", arr3) // [3]int

	var arr4 = [...]int{
		99: 39, // 将第100个元素(下标值为99)的值赋值为39，其余元素值均为0
	}
	fmt.Printf("%T\n", arr4) // [100]int
}
