## Functions Deep Dive

Functions are first-class citizens in Go. They can be assigned to variables, passed as arguments, and returned from other functions. Functions can also be anonymous, which means they don't have a name.

Here's an example of passing a function as an argument:

```go
func transformNumbers(numbers *[]int, transform func(num int) int) []int {
	doubledNumbers := []int{}

	for _, value := range *numbers {
		doubledNumbers = append(doubledNumbers, transform(value))
	}

	return doubledNumbers
}
```
