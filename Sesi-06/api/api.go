package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"html/template"
)

type Employee struct {
	ID 			int
	Name		string
	Age			int
	Division	string
}

var employees = []Employee {
	{ID: 1, Name: "Abdul", Age: 20, Division: "IT"},
	{ID: 2, Name: "Hanif", Age: 21, Division: "Finance"},
	{ID: 3, Name: "Pratama", Age: 22, Division: "Operators"},
}

var PORT = ":8081"

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		tpl, err := template.ParseFiles("template.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tpl.Execute(w, employees)
		return
	}

	http.Error(w, "Invalid Method", http.StatusBadRequest)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "Invalid Age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			ID:		len(employees) + 1,
			Name:	name,
			Age:	convertAge,
			Division: division,
		}
		
		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}

	http.Error(w, "Invalid Method", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/employees", getEmployees)

	http.HandleFunc("/employee", createEmployee)

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}