package main

import "fmt"

func main() {
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%x\n", i, r, r)

	}
}

///////////////////////////////////////////////////////
// <===================OUTPUT===================>	 //
// 0	'H'	48										 //
// 1	'e'	65										 //
// 2	'l'	6c										 //
// 3	'l'	6c										 //
// 4	'o'	6f										 //
// 5	','	2c										 //
// 6	' '	20										 //
// 7	'世'	4e16								 //
// 10	'界'	754c								 //
///////////////////////////////////////////////////////