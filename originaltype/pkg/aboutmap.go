package pkg

import "fmt"

// 	map[key_type]value_type
//	map[string]string // key与value元素的类型相同
//	map[int]string    // key与value元素的类型不同

//value无限制，但key必须有唯一性，即支持 == 和 != 运算

func Mapmain() {
	var m map[string]int
	//空的slice可以直接append，但是空的map是不可操作的，必须初始化
	fmt.Println(" ", m == nil)
}

func initialmap() {
	//可省略字面值类型，因为Go已经获知了map类型中的元素类型
	type Position struct {
		x float64
		y float64
	}

	m0 := map[Position]string{
		{29.935523, 52.568915}:  "school",
		{25.352594, 113.304361}: "shopping-mall",
		{73.224455, 111.804306}: "hospital",
	}
	fmt.Println(len(m0))

	//通过make初始化
	m1 := make(map[int]string)    // 未指定初始容量
	m2 := make(map[int]string, 8) // 指定初始容量为8
	fmt.Println(m1 == nil, m2 == nil)

	//查找
	m := make(map[string]int)
	v, ok := m["key1"]
	if !ok {
		// "key1"不在map中
	}
	fmt.Println(v) // "key1"在map中，v将被赋予"key1"键对应的value
	/*
			错误的查找方式：
		m := make(map[string]int)
		v := m["key1"]
		即使key1不在该map中，v也会被赋予一个value的类型零值
	*/

	//delete
	mmap := map[string]int{
		"key1": 1,
		"key2": 2,
	}

	fmt.Println(mmap)    // map[key1:1 key2:2]
	delete(mmap, "key2") // 删除"key2"
	fmt.Println(mmap)    // map[key1:1]

}

// 程序逻辑不可依赖map的遍历顺序，因为它每次遍历得到的元素顺序不一样
func Traverse() {
	m := map[int]int{
		1: 11,
		2: 12,
		3: 13,
	}

	fmt.Printf("{ ")
	for k, v := range m {
		fmt.Printf("[%d, %d] ", k, v)
	}
	fmt.Printf("}\n")
	/*
				for k, _ := range m {
				  	// 使用k
				}

				for k := range m {
			  		// 使用k
				}

				for _, v := range m {
		  			// 使用v
				}
	*/
}
