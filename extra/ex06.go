package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[1] = "One"
	m[2] = "Two"

	var v string
	var ok bool
	v, ok = m[2]
	fmt.Println(v, ok)
}
