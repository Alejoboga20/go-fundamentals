# Go Packages

## Splitting code across files

We can just create a new file and add the package declaration at the top of the file. The package declaration is the first line of the file and it should be the same for all files in the same package.

```go
package main

import (
	"fmt"
)

func presentOptions() {
	fmt.Println("What woud you like to do?")

	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	fmt.Print("Your choice: ")
}
```

Now we can use the `presentOptions` function in the `main.go` file without any issues as long as both files belong to the same package.
