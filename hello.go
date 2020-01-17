package main


import "fmt"

/*
 * init function added
 */
func init() {
	fmt.Println("rectangle package initialized")
}

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}

func lol() {
	num := 10
	if num % 2 == 0 { //checks if number is even
		fmt.Println("the number is even")
	}  else {
		fmt.Println("the number is odd")
	}
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a // a copy of a is assigned to b
	b[0] = "Singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)
	for i := 0; i < len(a); i++ { //looping from 0 to the length of the array
		fmt.Printf("%d th element of a is %s\n", i, a[i])
	}

	for i, v := range a {//range returns both the index and value
		fmt.Printf("%d the element of a is %s\n", i, v)
	}

	test := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}
	printarray(test)
	var bs [3][2]string
	bs[0][0] = "apple"
	bs[0][1] = "samsung"
	bs[1][0] = "microsoft"
	bs[1][1] = "google"
	bs[2][0] = "AT&T"
	bs[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(bs)


	asp := [5]int{76, 77, 78, 79, 80}
	fmt.Println(asp)
	var bsp []int = asp[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(bsp)

	darr := []int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("\n\narray before",darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after",darr)


	numa := [3]int{78, 79 ,80}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("\n\narray before change 1",numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)

	var names []string //zero value of a slice is nil
	if names == nil {
		fmt.Println("\n\nslice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:",names)
	}


	veggies := []string{"potatoes","tomatoes","brinjal"}
	fruits := []string{"oranges","apples"}
	food := append(veggies, fruits...)
	fmt.Println("\n\nfood:",food)

	nos := []int{8, 7, 6}
	fmt.Println("\n\nslice before function call", nos)
	subtactOne(nos)                               //function modifies the slice
	fmt.Println("slice after function call", nos)

	pls := [][]string {
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}

	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}

func switchcase() {
	finger := 4
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")

	}
}
func main() {
	lol()
	switchcase()
}