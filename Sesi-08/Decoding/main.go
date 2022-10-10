package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName	string		`json:"full_name"`
	Email		string		`json:"email"`
	Age			int			`json:"age"`
}

func main() {
	var jsonString = `[
			{
				"full_name": "Abdul Hanif",
				"email": "ahanifpratama@gmail.com",
				"age": 23
			},
			{
				"full_name": "Salman Alfarizi",
				"email": "salman@gmail.com",
				"age": 24
			}
		]	
	`
	
	// Decoding JSON to Struct
	// var result Employee

	// Decoding JSON to Map
	// var result map[string]interface{}

	// Decoding JSON To Empty Interface
	// var temp interface{}

	// Decoding JSON To Slice Of Struct
	var result []Employee
	
	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// var result = temp.(map[string]interface{})
	for i, v := range result {
		fmt.Printf("Index %d: %+v\n", i+1, v)
	}

	// fmt.Println("Full_name :", result["full_name"])
	// fmt.Println("Email     :", result["email"])
	// fmt.Println("Age       :", result["age"])
}