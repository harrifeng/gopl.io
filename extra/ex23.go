package main

import "fmt"

func main() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c)
	d := [3]int{1, 2}
	_ = d
	// fmt.Println(a == d) 		// invalid operation: a == d (mismatched types [2]int and [3]int)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// true false false								  //
////////////////////////////////////////////////////
