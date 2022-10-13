package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) SetName(name string) {
	s.Name = name
}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	var reflectType = reflectValue.Type()
	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama :", reflectType.Field(i).Name)
		fmt.Println("tipe data :", reflectType.Field(i).Type)
		fmt.Println("nilai :", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func main() {
	// Identifying Data Type & Value
	// var number = 23
	// var reflectValue = reflect.ValueOf(number)
	// //var nilai = reflectValue.Interface().(int)

	// fmt.Println("Tipe variabel :", reflectValue.Type())

	// if reflectValue.Kind() == reflect.Int {
	// 	fmt.Println("Nilai Variabel :", reflectValue.Int())
	// }

	// // Accessing Value Using Interface Method
	// fmt.Println()

	// fmt.Println("Tipe Variabel  :", reflectValue.Type())
	// fmt.Println("Nilai Variabel :", reflectValue.Interface())
	fmt.Println()

	// Pengaksesan Informasi Property Variabel Objek
	var s = &student{Name: "wick", Grade: 2}
	s.getPropertyInfo()

	// Identifying Method Information
	var s1 = &student{Name: "Abdul Hanif", Grade: 3}
	fmt.Println("Nama :", s1.Name)

	var reflectValue = reflect.ValueOf(s1)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("Hanif"),
	})

	fmt.Println("Nama :", s1.Name)
}
