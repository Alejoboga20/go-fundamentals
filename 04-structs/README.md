# Structs & Custom Types

A Struct is a user-defined data type that groups related data together. It is a collection of fields or members that can have different types. Structs are used to create custom data types that can be used in a program.
To declare a struct, you use the `struct` keyword followed by the name of the struct and a block of fields enclosed in curly braces. Each field has a name and a type.

```go
type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

appUser := User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
}
```

If we omit the field names, the fields are initialized in the order they are declared in the struct.

```go
type User struct {
  firstName string
  lastName  string
  birthdate string
  createdAt time.Time
}
```

Also we can omit assigning values to fields, in this case the fields will be initialized with their zero values.

We can also attach functions into a struct, these functions are called methods. Methods are functions that are associated with a particular type. They are similar to functions but are called using a receiver. To declare a method, you use the `func` keyword followed by the receiver type and the method name.

```go
type UserWithMethods struct {
	firstName string
	lastName  string
}

func (user UserWithMethods) outputUserDetailsOfUser() {
	fmt.Println("First Name:", user.firstName)
	fmt.Println("Last Name:", user.lastName)
}
```

To call a method, you use the dot operator followed by the method name.

```go
userWithMethod.outputUserDetailsOfUser()
```
