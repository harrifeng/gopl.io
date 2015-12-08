package main

import "fmt"

func main() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Println(s[0:5])
	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 12											  //
// 104 119										  //
// hello										  //
// hello										  //
// world										  //
// hello, world									  //
////////////////////////////////////////////////////
