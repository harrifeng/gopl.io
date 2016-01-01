package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	var dilbert Employee
	fmt.Println(dilbert.Salary)
	dilbert.Salary += 5000
	fmt.Println(dilbert.Salary)

	dilbert.Position = "Engineer"
	fmt.Println(dilbert.Position)
	position := &dilbert.Position
	*position = "Senior " + *position
	fmt.Println(dilbert.Position)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 0											  //
// 5000											  //
// Engineer										  //
// Senior Engineer								  //
////////////////////////////////////////////////////
