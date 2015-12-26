package main

import "fmt"

func main() {
	var ages map[string]int
	fmt.Println(ages == nil)
	fmt.Println(len(ages) == 0)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// true											  //
// true											  //
////////////////////////////////////////////////////
