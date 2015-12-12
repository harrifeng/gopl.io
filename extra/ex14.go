package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 13											  //
// 9											  //
////////////////////////////////////////////////////
