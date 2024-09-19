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
