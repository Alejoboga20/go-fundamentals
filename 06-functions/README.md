# Functions Deep Dive

Functions are first-class citizens in Go. They can be assigned to variables, passed as arguments, and returned from other functions. Functions can also be anonymous, which means they don't have a name.

Here's an example of passing a function as an argument and returning a function from another function:

```go
func transformNumbers(numbers *[]int, transform func(num int) int) []int {
	doubledNumbers := []int{}

	for _, value := range *numbers {
		doubledNumbers = append(doubledNumbers, transform(value))
	}

	return doubledNumbers
}

func getTransformerFunction(numbers *[]int) transformFunc {
	if (*numbers)[0] == 1 {
		return double
	}

	return triple
}
```

## Anonymous Functions

Anonymous functions are functions without a name. They are useful when you need to define a function inline. Here's an example of an anonymous function:

```go
func main() {
  func() {
    fmt.Println("Hello, World!")
  }()
}
```

## Factory Functions

Factory functions are functions that return another function. They are useful when you need to create a function with a specific behavior. Here's an example of a factory function:

```go

// factory function
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
```

Also note the use of closures in the above example. Closures are functions that reference variables from outside their body. In the above example, the inner function `func(number int) int` references the `factor` variable from the outer function `createTransformer`.

## Recursion

Recursion is a technique in which a function calls itself. It is useful when you need to solve a problem that can be broken down into smaller subproblems. Here's an example of a recursive function:

```go
func factorial(n int) int {
  if n == 0 {
    return 1
  }

  return n * factorial(n-1)
}
```

## Variadic Functions

Variadic functions are functions that can take a variable number of arguments. The `...` syntax is used to indicate that a function is variadic. Here's an example of a variadic function:

```go
package main

import "fmt"

func main() {
	sum := sumup(1, 2, 3, 4, 5)
	fmt.Println("Sum of numbers is", sum)
}

func sumup(srartingValue int, numbers ...int) int {
	sum := 0

	fmt.Println("Starting value is", srartingValue)

	for _, number := range numbers {
		sum += number
	}

	return sum
}
```
