package main

import (
	"fmt"
)

type SalaryCalculatorLocallocal interface {
	DisplaySalary()
}

type LeaveCalculatorLocal interface {
	CalculateLeavesLeft() int
}

type EmployeeOperations interface {
	SalaryCalculatorLocallocal
	LeaveCalculatorLocal
}

type EmployeeLocal struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e EmployeeLocal) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e EmployeeLocal) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {
	e := EmployeeLocal{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var empOp EmployeeOperations = e
	empOp.DisplaySalary()
	fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())
}

