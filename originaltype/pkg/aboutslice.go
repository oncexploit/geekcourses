package pkg

import "fmt"

/*
type slice struct {
    array unsafe.Pointer	array: 是指向底层数组的指针；
    len   int				len: 是切片的长度，即切片中当前元素的个数;
    cap   int				cap: 是底层数组的长度，也是切片的最大容量，cap 值永远大于等于 len 值。
}
*/

// 这是一个切片，与数组相比，它仅仅少了长度属性
var nums = []int{1, 2, 3, 4, 5, 6}

func main() {
	nums = append(nums, 7)

	//直接创建切片
	sl1 := make([]byte, 6, 10) // 其中10为cap值，即底层数组长度，6为切片的初始长度
	sl2 := make([]byte, 6)     // cap = len = 6
	fmt.Println(sl1, sl2)

	//基于已存在数组创建切片
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sl := arr[3:7:9]
	sl[0] += 10
	fmt.Println("arr[3] = ", arr[3])
}
