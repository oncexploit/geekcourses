package foringo

//for range is the only way

func aboutforinmap() {
	var m = map[string]int{
		"Rob":  67,
		"Russ": 39,
		"John": 29,
	}

	for k, v := range m {
		println(k, v)
	}
}
