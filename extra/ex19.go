package main

import "fmt"

func main() {
	const (
		First uint = 1 << iota
		Second
		Third
		Fourth
		Fifth
	)

	fmt.Println(First, Second, Third, Fourth, Fifth)
	fmt.Println("Binary:")
	fmt.Printf("%b, %b, %b, %b, %b", First, Second, Third, Fourth, Fifth)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 1 2 4 8 16									  //
// Binary:										  //
// 1, 10, 100, 1000, 10000						  //
////////////////////////////////////////////////////
