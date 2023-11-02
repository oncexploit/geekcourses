package foringo

import "fmt"

func Formain() {
	//normal form
	var sl = []int{1, 2, 3, 4, 5}
	for i := 0; i < len(sl); i++ {
		fmt.Printf("sl[%d] = %d\n", i, sl[i])
	}

	//for range way
	for i, v := range sl {
		fmt.Printf("sl[%d] = %d\n", i, v)
	}
}

//different ways of for range

func aboutrange() {
	var sl = []int{1, 2, 3, 4, 5}
	//ignore elements
	for i := range sl {
		// ...
		fmt.Println(i)
	}

	//ignore index
	for _, v := range sl {
		// ...
		fmt.Println(v)
	}

	//ignore both
	for range sl {
		// ...
	}
}
