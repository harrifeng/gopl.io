package main

import "fmt"

var p = f()

func f() *int {
	v := 1
	return &v
}

func main() {
	fmt.Println(p)
	fmt.Println(f() == f())
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 0x82024e220									  //
// false										  //
////////////////////////////////////////////////////
