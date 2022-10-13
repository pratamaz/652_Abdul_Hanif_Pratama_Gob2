package main

// import "fmt"

func main() {
	var randomValues interface{}
	_ = randomValues
	randomValues = "Jalan Sudirman"
	randomValues = 20
	randomValues = true
	randomValues = []string{"Abdul", "Hanif"}

	// Empty Interface - Type Assetion
	var v interface{}

	v = 20

	if value, ok := v.(int); ok == true {
		v = value * 9
	}


	// Map & Slice with Empty Interface
	rs := []interface{}{1, "Abdul", true, 2, "Hanif", true}

	rm := map[string]interface{}{
		"Nama" : "Abdul",
		"Status" : true,
		"Age":	23,
	}

	_, _ = rs, rm
}