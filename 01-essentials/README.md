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

## Resources

- [Go Standard Packages](https://pkg.go.dev/std)
