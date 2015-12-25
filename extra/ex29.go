package main

import "fmt"

func main() {
	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	fmt.Println(x)
	x = append(x, x...)
	fmt.Println(x)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// [1 2 3 4 5 6]								  //
// [1 2 3 4 5 6 1 2 3 4 5 6]					  //
////////////////////////////////////////////////////
