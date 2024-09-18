package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "John"
	userNames[1] = "Doe"

	fmt.Println(userNames)
}
