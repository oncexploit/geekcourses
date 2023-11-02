package foringo

import "fmt"

func learnmore() {
	var sum int
	var sl = []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(sl); i++ {
		if sl[i]%2 == 0 {
			// 忽略切片中值为偶数的元素
			continue
		}
		sum += sl[i]
	}
	println(sum) // 9
}

func main() {
	var sl = [][]int{
		{1, 34, 26, 35, 78},
		{3, 45, 13, 24, 99},
		{101, 13, 38, 7, 127},
		{54, 27, 40, 83, 81},
	}

outerloop: //continue用于从此处继续循环
	for i := 0; i < len(sl); i++ {
		for j := 0; j < len(sl[i]); j++ {
			if sl[i][j] == 13 {
				fmt.Printf("found 13 at [%d, %d]\n", i, j)
				continue outerloop //用于跳出嵌套循环中的内部循环
			}
		}
	}
}

func anothercontrol() {
	var sl = []int{5, 19, 6, 3, 8, 12}
	var firstEven = -1

	// 找出整型切片sl中的第一个偶数
	for i := 0; i < len(sl); i++ {
		if sl[i]%2 == 0 {
			firstEven = sl[i]
			break
		}
	}

	println(firstEven) // 6

	var gold = 38

	var sl2 = [][]int{
		{1, 34, 26, 35, 78},
		{3, 45, 13, 24, 99},
		{101, 13, 38, 7, 127},
		{54, 27, 40, 83, 81},
	}

outerloop: //break用于跳出此处的循环
	for i := 0; i < len(sl); i++ {
		for j := 0; j < len(sl2[i]); j++ {
			if sl2[i][j] == gold {
				fmt.Printf("found gold at [%d, %d]\n", i, j)
				break outerloop
			}
		}
	}

}
