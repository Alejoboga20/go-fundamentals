// import "fmt"

// // func main() {
// // 	productNames := [4]string{"laptop", "mouse", "keyboard"}
// // 	prices := []float64{1.50, 2.35, 3.123, 4, 5}

// // 	fmt.Println(productNames)
// // 	productNames[0] = "monitor"
// // 	fmt.Println(productNames)

// // 	for i := range prices {
// // 		fmt.Println(prices[i])
// // 	}

// // 	// Slices
// // 	featuredPricesOne := prices[1:3]
// // 	featuredPricesTwo := prices[:3]
// // 	featuredPricesThree := prices[1:]
// // 	concatenatedPrices := featuredPricesOne[:1]

// // 	fmt.Println(featuredPricesOne, featuredPricesTwo, featuredPricesThree, concatenatedPrices)

// // 	// Slices metadata
// // 	fmt.Println(len(prices), cap(prices))
// // 	fmt.Println(len(featuredPricesOne), cap(featuredPricesOne))

// // 	// we can re-slice and select more elements to the right
// // 	concatenatedPrices = concatenatedPrices[:3]
// // 	fmt.Println(featuredPricesOne, concatenatedPrices)

// // }

// func main() {
// 	// if we don't specify the size of the array, it will be a slice, not an array
// 	// with an slice we can add or remove elements, no matter the size
// 	prices := []float64{10.99, 5.99, 3.99}
// 	fmt.Println(prices[0])

//		// we can not add elements to indexes that are not in the slice e.g. prices[3] = 1.99
//		// to add elements we use the append function
//		updatedPrices := append(prices, 1.99)
//		pricesWithRemovedElements := prices[1 : len(prices)-1]
//		fmt.Println(updatedPrices, prices, pricesWithRemovedElements)
//	}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

package lists

import "fmt"

type Product struct {
	id     int
	title  string
	prices float64
}

func main() {
	hobbies := [3]string{"reading", "coding", "gaming"}

	fmt.Println(hobbies)
	fmt.Println(hobbies[0])

	slicedHobbies := hobbies[1:3]
	fmt.Println(slicedHobbies)

	goals := []string{"learn go", "learn react"}
	fmt.Println(goals)

	goals[1] = "learn python"
	updatedGoals := append(goals, "learn rust")
	fmt.Println(updatedGoals)

	products :=
		[]Product{
			{1, "laptop", 1000.00},
			{2, "mouse", 20.00},
		}

	updatedProducts := append(products, Product{3, "keyboard", 50.00})
	fmt.Println(updatedProducts)

	prices := []float64{10.99, 5.99, 3.99}
	fmt.Println(prices[0])

	updatedPrices := append(prices, 1.99)
	pricesWithRemovedElements := prices[1 : len(prices)-1]
	fmt.Println(updatedPrices, prices, pricesWithRemovedElements)

	discountPrices := []float64{1.99, 2.99, 3.99}
	// to unpack the slice we use the ... operator
	updatedPrices = append(updatedPrices, discountPrices...)
	fmt.Println(updatedPrices)
}
