package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("foo.txt")
	fmt.Println(f)
	fmt.Println(err)
	//! ERROR no new variables on left side of :=
	// f, err := os.Open("foo.txt")
	f, err = os.Open("foo.txt")
}
