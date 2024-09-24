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
