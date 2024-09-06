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

When we create methods to modify the strcuct fields, we need to use a pointer receiver. This is because Go is a pass-by-value language, so when we pass a struct to a function, a copy of the struct is created. If we want to modify the original struct, we need to pass a pointer to the struct.

```go
type UserWithPointerReceiver struct {
  firstName string
  lastName  string
}

func (user *UserWithMethods) clearUserName() {
	user.firstName = ""
	user.lastName = ""
}
```

Constructor functions are functions that return a new instance of a struct. They are used to create and initialize a struct. We return a pointer to the struct because we want to modify the struct fields.

```go
func newUser(firstName, lastName, birthdate string) *User {
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}
```

If we create a package for a struct, the same rules for exporting and unexporting apply. If we want to export a struct, we need to capitalize the first letter of the struct name. If we don't want to export a struct, we need to use a lowercase letter. The same apply for some fields in the struct. If the fields should not be exported, we need to use a lowercase letter (for example `FirstName` is exported while `createdAt` is not)

```go
package user

type User struct {
  FirstName string
  LastName  string
  Birthdate string
  createdAt time.Time
}
```

We can also embed a struct into another struct. This is called composition. When we embed a struct into another struct, the fields and methods of the embedded struct are promoted to the outer struct. This means that we can access the fields and methods of the embedded struct as if they were part of the outer struct.

```go
type Admin struct {
	emial    string
	password string
	user User
}


func NewAdmin(email, password string) Admin {
	return Admin{
		emial:    email,
		password: password,
		user: User{
			FirstName: "Admin",
			LastName:  "Admin",
			Birthdate: "01/01/2000",
			CreatedAt: time.Now(),
		},
	}
}
```
