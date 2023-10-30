package pkg

import (
	"fmt"
	"unicode/utf8"
)

//原生类型，有类型检查

//string生命周期内不可变，避免同步安全问题

/*
var s string = "hello"
s[0] = 'k'   // 错误：字符串的内容是不可改变的
s = "gopher" // ok

string类型虽然是不能更改的，但是可以被替换，因为stringStruct中的str指针是可以改变的，
只是指针指向的内容是不可以改变的，也就说每一次更改字符串，就需要重新分配一次内存，之前分配的空间会被gc回收
*/
func Toseeistobelieve() {
	var s string = `         ,_---~~~~~----._
    _,,_,*^____      _____*g*\"*,--,
   / __/ /'     ^.  /      \ ^@q   f
  [  @f | @))    |  | @))   l  0 _/
   \/   \~____ / __ \_____/     \
    |           _l__l_           I
    }          [______]           I
    ]            | | |            |
    ]             ~ ~             |
    |                            |
     |                           |`
	fmt.Println(s)
}

func Digintostring() {
	var s = "中国人"
	fmt.Printf("the length of s = %d\n", len(s)) // 9

	for i := 0; i < len(s); i++ {
		fmt.Printf("0x%x ", s[i]) // 0xe4 0xb8 0xad 0xe5 0x9b 0xbd 0xe4 0xba 0xba
	}
	fmt.Printf("\n")

	//meaningful way
	fmt.Println("the character count in s is", utf8.RuneCountInString(s)) // 3

	for _, c := range s {
		fmt.Printf("0x%x ", c) // 0x4e2d 0x56fd 0x4eba
	}
	fmt.Printf("\n")
}

func Aboutrune() {
	// rune is another way of int32
	//type rune = int32 (in $GOROOT/src/builtin.go)
	var a = 'A'
	var b = '\u4e2d'
	var c = '\x27'
	fmt.Println(a, b, c)
}

func Addstring() {
	s := "Rob Pike, "
	s = s + "Robert Griesemer, "
	s += " Ken Thompson"

	fmt.Println(s) // Rob Pike, Robert Griesemer, Ken Thompson
}

func Comparestring() {
	// ==
	s1 := "世界和平"
	s2 := "世界" + "和平"
	fmt.Println(s1 == s2) // true

	// !=
	s1 = "Go"
	s2 = "C"
	fmt.Println(s1 != s2) // true

	// < and <=
	s1 = "12345"
	s2 = "23456"
	fmt.Println(s1 < s2)  // true
	fmt.Println(s1 <= s2) // true

	// > and >=
	s1 = "12345"
	s2 = "123"
	fmt.Println(s1 > s2)  // true
	fmt.Println(s1 >= s2) // true
}

func ConvertString() {
	var s string = "中国人"

	// string -> []rune
	rs := []rune(s)
	fmt.Printf("%x\n", rs) // [4e2d 56fd 4eba]

	// string -> []byte
	bs := []byte(s)
	fmt.Printf("%x\n", bs) // e4b8ade59bbde4baba

	// []rune -> string
	s1 := string(rs)
	fmt.Println(s1) // 中国人

	// []byte -> string
	s2 := string(bs)
	fmt.Println(s2) // 中国人
}

// rune -> []byte
func encodeRune() {
	var r rune = 0x4E2D                            //unicode码点值
	fmt.Printf("the unicode charactor is %c\n", r) // 中
	buf := make([]byte, 3)
	_ = utf8.EncodeRune(buf, r)                       // 对rune进行utf-8编码
	fmt.Printf("utf-8 representation is 0x%X\n", buf) // 0xE4B8AD
}

// []byte -> rune
func decodeRune() {
	var buf = []byte{0xE4, 0xB8, 0xAD}
	r, _ := utf8.DecodeRune(buf)                                                             // 对buf进行utf-8解码
	fmt.Printf("the unicode charactor after decoding [0xE4, 0xB8, 0xAD] is %s\n", string(r)) // 中
}
