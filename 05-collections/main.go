package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "John"
	userNames[1] = "Doe"

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)

	courseRatings["Go"] = 4.7
	courseRatings["Python"] = 4.5
	courseRatings["Java"] = 4.8

	courseRatings.output()
}
