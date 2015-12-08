package main

import "fmt"

func main() {
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i

	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// (-5+10i)										  //
// -5											  //
// 10											  //
////////////////////////////////////////////////////
