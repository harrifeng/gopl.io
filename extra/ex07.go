package main

import (
	"fmt"
)

func main() {
	var x interface{}
	x = "abc"
	var v string
	var k bool
	v, k = x.(string)
	if k {
		fmt.Println(v, k)
	} else {
		fmt.Println("first x is Not a String")
	}

	x = 12
	v, k = x.(string)
	if k {
		fmt.Println(v, k)
	} else {
		fmt.Println("second x is Not a String")
	}
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// abc true										  //
// second x is Not a String						  //
////////////////////////////////////////////////////
