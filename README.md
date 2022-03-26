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
	HasCar   bool
}

// Main function
func main() {
	// *** Entity to Entity ***
	p := Personal{FullName: "John Doe", Age: 33, Salary: 100, HasCar: true}
	m := Manager{}

	auto.Mapper(p, &m)
	fmt.Println(m)
	// output
	{ "John Doe", 33, 100, true }
	
	// ******
	
	// *** List to List ***
	l := []Personal{
		{
			FullName: "Name 1",
			Age:      30,
			HasCar:   false,
			Salary:   100,
		},
		{
			FullName: "Name 2",
			Age:      40,
			HasCar:   true,
			Salary:   200,
		},
		{
			FullName: "Name 3",
			Age:      50,
			HasCar:   false,
			Salary:   300,
		},
	}

	var tObj = automapper.MapperForList(l, []Manager{})
	mList, _ := tObj.([]Manager)
	fmt.Println(mList)
	// output
	[
	   {Name 1 30 100 false} 
	   {Name 2 40 200 true} 
	   {Name 3 50 300 false}
	]
	
	// ******
}
```
# Installation
```
go get -u github.com/erhankrygt/automapper
```
