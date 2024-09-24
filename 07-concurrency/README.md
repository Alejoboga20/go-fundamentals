# Concurrency

Concurrency is the ability to run multiple tasks at the same time. This is different from parallelism, which is the ability to run multiple tasks simultaneously. Concurrency is achieved by interleaving tasks in time, while parallelism is achieved by running tasks at the same time.

Go has built-in support for concurrency. It provides goroutines, which are lightweight threads managed by the Go runtime. Goroutines are functions that run concurrently with other functions. They are created using the `go` keyword followed by a function call.

Also Go has channels, which are used to communicate between goroutines. Channels are used to send and receive values between goroutines. They are created using the `make` function with the `chan` keyword.

```go
package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)

	doneChan <- true
}

func main() {
	done := make(chan bool)
	go slowGreet("How ... are ... you ...?", done)
	greet("How are you?")
	<-done
}
```
