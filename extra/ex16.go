package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	fmt.Println(strconv.FormatInt(int64(x), 2))

	fmt.Printf("x=%b\n", x)

	x2, _ := strconv.Atoi("123")
	y2, _ := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits

	fmt.Println(x2)
	fmt.Println(y2)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 123 123										  //
// 1111011										  //
// x=1111011									  //
// 123											  //
// 123											  //
////////////////////////////////////////////////////
