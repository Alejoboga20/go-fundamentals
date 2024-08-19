# Go Essentials

## Organizing Code with Packages

When creating code in go we split it into packages. Every GO file should start with package name

We need at least one package in our code, the `main` package. This package is the entry point of our code. But we can have multiple packages in our code.

This way we can import/export code between packages.

Go also uses the concept of **modules**. Modules are a collection of related Go packages that are versioned together as a single unit.

To create a module we need to run:

```bash
go mod init [module-name]
```

This will create a `go.mod` file in the root of our project. This file contains the module name and the dependencies of our project.

After this we can use `go build` to build our project and generate and executable file.

We should always have a `main` package in our project. This package is the entry point of our code. We should also have a `main` function in this package so GO knows where to start executing our code.

```go
package main

import "fmt"

func main() {
	fmt.Print("Hello, World!")
}
```

We can only have one `main` function in our `main` package. If we are building a library we should not have a `main` package because they are not intended to be executed as a program

## Understanding Functions

Functions are code blocks that can be called from other parts of our code.Functions in Go are defined using the `func` keyword. They can have zero or more parameters and zero or more return values.
We can define the type of the parameters and the return values. They can receive no parameters and return no values as well.

```go
func add(x int, y int) int {
  result = x + y

  return result
}
```

## Resources

- [Go Standard Packages](https://pkg.go.dev/std)
