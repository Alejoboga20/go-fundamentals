# Pointers

## Introduction

Pointers are variables that store the memory address instead of values. To retrieve the memory address of a variable, we use the `&` operator. To retrieve the value of a memory address, we use the `*` operator.

In Go pointers are used to pass the reference of a variable to a function, instead of passing the value. This is useful when we want to modify the value of a variable inside a function. Another use case is when we want to avoid copying large data structures.

Here's an example of the use of pointers to avoid unnecessary copying of data

```go
package main

import "fmt"

func main() {
	age := 32
	agePointer := &age
	ageValueFromPointer := *agePointer
	adultYears := getAdultYears(agePointer)

	fmt.Println("age:", age)
	fmt.Println("agePointer:", agePointer)
	fmt.Println("ageValueFromPointer:", ageValueFromPointer)
	fmt.Println("getAdultYears:", adultYears)
}

func getAdultYears(age *int) int {
	return *age - 18
}
```

The other use case is for data mutation by passing the reference of a variable to a function. Here's an example:

```go
func mutateWithPointer(age *int) {
	*age = *age + 1
}
```

function `mutateWithPointer` will increment the value of the variable passed as a pointer by 1.
