package main

import (
	"Golang-Hacktiv8/Sesi-03/export-unexport/helpers"
)

func main() {
	helpers.Greet()
	var person = helpers.Person{}
	person.Invokegreet()
}