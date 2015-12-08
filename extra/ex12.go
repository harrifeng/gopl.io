package main

import "fmt"

func main() {

	s := "h世界"
	fmt.Println(len(s))
	for k, v := range s {
		fmt.Println(k, v)
	}

	fmt.Println(s)
}
