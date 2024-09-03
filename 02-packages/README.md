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

## Using multiple packages

Each package should be in its own directory. The name of the directory should be the same as the name of the package. For example this file is in the `fileutils` directory and the package name is `fileutils`.

```go
package fileutils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(fileName string, floatValue float64) {
	dataText := fmt.Sprint(floatValue)
	os.WriteFile(fileName, []byte(dataText), 0644)
}

func GetFloatFromFile(fileName string, defaultValue ...float64) (float64, error) {
	// ...float64 is a variadic parameter: it allows the function to accept 0 or more float64 values
	var defVal float64
	data, err := os.ReadFile(fileName)

	if len(defaultValue) > 0 {
		defVal = defaultValue[0]
	} else {
		defVal = 0
	}

	if err != nil {
		return defVal, errors.New("error reading from file")
	}

	dataString := string(data)
	dataFloat, err := strconv.ParseFloat(dataString, 64)

	if err != nil {
		return defVal, errors.New("error parsing from file")
	}

	return dataFloat, nil
}
```

Now we can use the `fileutils` package in the `main.go` file by importing it including the module name.

```go

import (
	"fmt"

	"example.com/bank/fileutils"
)
```

**_Note:_** The module name is the name of the module that we defined in the `go.mod` file. And the exported functions in the package should start with an uppercase letter.

## Using third-party packages

To use third-party packages we need to import them in the file where we want to use them. We can use the `go get` command to download the package and add it to the `go.mod` file.

```bash
go get github.com/gorilla/mux
```

Now we can use the `gorilla/mux` package in our code.

```go
package main

import (
  "fmt"
  "net/http"

  "github.com/gorilla/mux"
)
```
