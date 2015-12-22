package main

import "fmt"

func main() {
	type Weekday int

	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 0 1 2 3 4 5 6								  //
////////////////////////////////////////////////////
