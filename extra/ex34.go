package main

import "fmt"

var m = make(map[string]int)

func k(list []string) string { return fmt.Sprintf("%q", list) }

func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }

func main() {

	Add([]string{"1", "2", "3"})
	Add([]string{"1", "2", "3"})
	Add([]string{"1", "2", "3", "4"})
	Add([]string{"1", "2", "3"})

	fmt.Println(Count([]string{"1", "2", "3"}))

	for k, v := range m {
		fmt.Println(k, v)
	}

}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 3											  //
// ["1" "2" "3" "4"] 1							  //
// ["1" "2" "3"] 3								  //
////////////////////////////////////////////////////
