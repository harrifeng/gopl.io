package main

import "fmt"

func main() {
	var runes []rune

	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// ['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']		  //
////////////////////////////////////////////////////
