package main

import "fmt"

func main() {
	fmt.Printf("%d\n", new(int))
	fmt.Printf("%d\n", new(int))
	fmt.Printf("%d\n", new([0]int))
	fmt.Printf("%d\n", new([0]int))
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 34898576208									  //
// 34898576240									  //
// &[]											  //
// &[]											  //
////////////////////////////////////////////////////