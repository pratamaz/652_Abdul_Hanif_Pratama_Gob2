package main

import (
	"Sesi-03/export-unexport/helpers"
)

func main() {
	helpers.Greet()
	var person = helpers.Person{}
	person.Invokegreet()
}