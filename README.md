# automapper

AutoMapper is a simple little library built to solve a deceptively complex problem - getting rid of code that mapped one object to another. 


# Example
```go
package main

import (
	"fmt"
	auto "github.com/erhankrygt/automapper"
)

type Personal struct {
	FullName string
	Age      int
	Salary   float64
	HasCar   bool
}

type Manager struct {
	FullName string
	Age      int
	Salary   float64
	Manager  bool
	CarHas   bool
}

// Main function
func main() {
	p := Personal{FullName: "John Doe", Age: 33, Salary: 100, HasCar: false}
	m := Manager{}
	auto.Mapper(p, &m, map[string]string{"HasCar": "CarHas"})

	fmt.Println(m)
	// output
	{ "John Doe", 33, 100, false, true }
	
	auto.Mapper(p, &m, nil)
	fmt.Println(m)
	{ "John Doe", 33, 100, false, false }
	
}
```
# Installation
```
go get -u github.com/erhankrygt/automapper
```
