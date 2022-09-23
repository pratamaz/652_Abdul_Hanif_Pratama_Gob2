package main

import (
	"fmt"
)

// Embedded Struct
type Person struct {
	name string
	age int
}

type Employee struct {
	division string
	person   Person
}

func main() {
	fmt.Println("====================================")
	fmt.Println("       Giving Value to Struct       ")
	fmt.Println("====================================")
	fmt.Println()
	
	// var employee1 Employee
	
	// employee1.name = "Abdul"
	// employee1.age = 22
	// employee1.division = "Back End Developer"
	
	// fmt.Println(employee1.name)
	// fmt.Println(employee1.age)
	// fmt.Println(employee1.division)
	
	fmt.Println()

	
	fmt.Println("====================================")
	fmt.Println("        Initializing Struct         ")
	fmt.Println("====================================")
	fmt.Println()
	// var employee2 = Employee{name: "Hanif", age: 23, division: "Front End Developer"}
	
	// fmt.Printf("Employee1: %+v\n", employee1)
	// fmt.Printf("Employee2: %+v\n", employee2)
	
	// fmt.Println()
	
	// // Pointer to a Struct
	fmt.Println("====================================")
	fmt.Println("        Pointer to a Struct         ")
	fmt.Println("====================================")
	fmt.Println()
	// var employee3 = Employee{name: "Pratama", age: 22, division: "Flutter Developer"}
	
	// var employee4 *Employee = &employee3

	// fmt.Println("Employee3 name", employee3.name)
	// fmt.Println("Employee4 name", employee4.name)
	
	// fmt.Println(strings.Repeat("#", 20))
	
	// employee4.name = "Doel"
	
	// fmt.Println("Employee3 name", employee3.name)
	// fmt.Println("Employee4 name", employee4.name)
	
	fmt.Println()

	fmt.Println("====================================")
	fmt.Println("          Embedded Struct           ")
	fmt.Println("====================================")

	var employee5 = Employee{}

	employee5.person.name = "Amril"
	employee5.person.age = 24
	employee5.division = "Devops Engineer"

	fmt.Printf("%+v", employee5)
	fmt.Println()
	fmt.Println()


	fmt.Println("====================================")
	fmt.Println("           Anonymous Struct         ")
	fmt.Println("====================================")
	
	// Anonymous struct tanpa pengisian field
	var employee1 = struct {
		person		Person
		division 	string 
	}{}
	employee1.person = Person{name: "Abdul", age: 23}
	employee1.division = "UX Resercher"

	// Anonymous struct dengan pengisian Field
	var employee2  = struct {
		person 		Person
		division 	string
	}{
		person: Person{name: "Hanif", age: 22},
		division: "UI Designer",
	}

	fmt.Printf("Employee1: %+v\n", employee1)
	fmt.Printf("Employee2: %+v\n", employee2)

	fmt.Println()


	fmt.Println("====================================")
	fmt.Println("           Slice of Struct          ")
	fmt.Println("====================================")
	
	var people = []Person{
		{name: "Abdul", age: 22},
		{name: "Hanif", age: 21},
		{name: "Pratama", age: 23},
	}

	for _,v := range people {
		fmt.Printf("%+v\n", v)
	}

	fmt.Println()


	fmt.Println("====================================")
	fmt.Println("     Slice of Anonymous Struct      ")
	fmt.Println("====================================")

	var employee = []struct{
		person		Person
		division	string
	}{
		{person: Person{name: "Abdul", age: 22}, division: "Back End Developer"},
		{person: Person{name: "Hanif", age: 21}, division: "Front End Developer"},
		{person: Person{name: "Prata", age: 23}, division: "Android Developer"},
	}

	for _, v := range employee {
		fmt.Printf("%+v\n", v)
	}
	fmt.Println()
}