package main

import "fmt"

func main() {
	age := 32                          // regular variable
	agePointer := &age                 // pointer variable
	ageValueFromPointer := *agePointer // pointer variable
	adultYears := getAdultYears(agePointer)

	fmt.Println("age:", age)
	fmt.Println("agePointer:", agePointer)
	fmt.Println("ageValueFromPointer:", ageValueFromPointer)
	fmt.Println("getAdultYears:", adultYears)
}

// * in the argument indicates that the function receives a pointer. No copy of the value is made
// *age is used to get the value from the pointer
func getAdultYears(age *int) int {
	return *age - 18
}
