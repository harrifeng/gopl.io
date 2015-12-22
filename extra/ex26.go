package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	// rotate s by 2
	s := []int{0, 1, 2, 3, 4, 5} // Important! slice is []int type
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// [2 3 4 5 0 1]								  //
////////////////////////////////////////////////////
