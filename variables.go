package main

import "fmt"

func abtest()  {
	a, b := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n", a, b)
	sum := a + b
	diff := a - b
	fmt.Println("sum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)
}

func constantTest() {
	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar",intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var",complex64Var)

}

func ageTest() {
	var age int // variable declaration
	fmt.Println("my age is ", age)
	age = 29 //assignment
	fmt.Println("my age is", age)
	age = 54 //assignment
	fmt.Println("my new age is", age)
}

func nameTest() {
	var (
		name   = "naveen"
		ager    = 29
		height int
	)
	fmt.Println("\n\nmy name is", name, ", age is", ager, "and height is", height)
}

func abctest() {
	a, b := 20, 30 // declare variables a and b
	fmt.Println("a is", a, "b is", b)
	b, c := 40, 50 // b is already declared but c is new
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90 // assign new values to already declared variables b and c
	fmt.Println("changed b is", b, "c is", c)
}

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func findV2(num int, nums []int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func mapTest() {
	personSalary := make(map[string]int)
	personSalary["steve"] = 12000
	personSalary["jamie"] = 15000
	personSalary["mike"] = 9000
	fmt.Println("personSalary map contents:", personSalary)
}

func maptest()  {
	personSalary := map[string]string{
		"steve": "12000",
		"jamie": "15000",
	}
	personSalary["mike"] = "9000"
	newEmp := "steve"
	value, ok := personSalary[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
	} else {
		fmt.Println(newEmp,"not found")
	}

	for key, value := range personSalary {
		fmt.Printf("personSalary[%s] = %s\n", key, value)
	}
}

func printBytes(s string) {
	for i:= 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	for i:= 0; i < len(s); i++ {
		fmt.Printf("%c ",s[i])
	}
}

func printRunes(s string) {
	runes := []rune(s)
	for i:= 0; i < len(runes); i++ {
		fmt.Printf("%c ",runes[i])
	}
}

func pointerIntro() {
	b := 255
	var a *int = &b
	fmt.Println("Value of b is", b)
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

}

func zeroPointer() {
	a := 25
	var b *int
	if b == nil {
		fmt.Println("\n\nb is", b)
		b = &a
		fmt.Println("b after initialization is", b)
	}
}

func sizeDemo() {
	size := new(int)
	fmt.Printf("\n\nSize value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size)
	changeValue(size)
	fmt.Println("Changed size value is", *size)
}

func changeValue(val *int) *int{
	*val = 155
	return val
}

func modify(arr []int) {
	arr[0] = 000
}

func arr(){
	a := [3]int{89, 90, 91}
	fmt.Println(a)
	modify(a[:])
	fmt.Println(a)
}

func main() {
	//ageTest()
	//nameTest()
	//abctest()
	//abtest()
	//constantTest()

	//find(89, 89, 90, 95)
	//
	//findV2(45, []int{56, 67, 45, 90, 109})
	//findV2(78, []int{38, 56, 98})
	//
	//mapTest()
	//maptest()
	//
	//printBytes("lol")
	//
	//printChars("Señor")
	//printRunes("Señor")

	//pointerIntro()
	//zeroPointer()
	//sizeDemo()

	arr()
}

func init() {
}