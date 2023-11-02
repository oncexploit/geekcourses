package ifingo

/* formal expression
if boolean_expression {
    // 新分支	(true)
}
(false)
// 原分支

if boolean_expression1 {
  // 分支1
} else if boolean_expression2 {
  // 分支2

... ...

} else if boolean_expressionN {
  // 分支N
} else {
  // 分支N+1
}
*/

func calcprev() {
	a, b := false, true
	//使用小括号表达你希望的执行顺序更可取
	if a && b != true { // "!="的优先级比"&&"高
		println("(a && b) != true")
		return
	}
	println("a && (b != true) == false")
}

// 声明if语句的自用变量
func main() {
	if a, c := 1, 2; a > 0 {
		println(a)
	} else if b := 3; b > 0 {
		println(a, b)
	} else {
		println(a, b, c)
	}
}

/*
if的happy path原则：
仅使用单分支控制结构；
当布尔表达式求值为 false 时，也就是出现错误时，在单分支中快速返回；
正常逻辑在代码布局上始终“靠左”，这样读者可以从上到下一眼看到该函数正常逻辑的全貌；
函数执行到最后一行代表一种成功状态
*/
