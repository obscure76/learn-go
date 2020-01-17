package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age, salary         int
}

type Address struct {
	city, state string
}
type Person struct {
	name string
	age int
	address Address
}

type PromotedPerson struct {
	name string
	age int
	Address
}

func define_ds() {
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}
	emp2 := Employee{"Thomas", "Paul", 29, 800}
	fmt.Println(emp1)
	fmt.Println(emp2)
}

func anonymous_ds() {
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}

	fmt.Println("Employee 3", emp3)
}

func access_members() {
	emp6 := Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("\n\nFirst Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d", emp6.salary)
}

func nested_ds() {
	var p Person
	p.name = "Naveen"
	p.age = 50
	p.address = Address {
		city: "Chicago",
		state: "Illinois",
	}
	fmt.Println("\n\nName:", p.name)
	fmt.Println("Age:",p.age)
	fmt.Println("City:",p.address.city)
	fmt.Println("State:",p.address.state)
}

func promoted_ds() {
	var p PromotedPerson
	p.name = "Naveen"
	p.age = 50
	p.Address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.city) //city is promoted field
	fmt.Println("State:", p.state)
}

type name struct {
	firstName string
	lastName string
}

func equal_structs() {
	name1 := name{"Steve", "Jobs"}
	name2 := name{"Steve", "Jobs"}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name{firstName:"Steve", lastName:"Jobs"}
	name4 := name{}
	name4.firstName = "Steve"
	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}
}

type image struct {
	data map[int]int
}

func img_diff()  {
	image1 := image{data: map[int]int{
		0: 155,
	}}
	image2 := image{data: map[int]int{
		0: 155,
	}}
	if len(image1.data) == len(image2.data) {
		fmt.Println("image1 and image2 are equal")
	}
}

func main() {
	//anonymous_ds()
	//access_members()
	//
	//nested_ds()
	//promoted_ds()

	equal_structs()
	img_diff()
}
