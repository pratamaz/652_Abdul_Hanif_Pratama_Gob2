package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host	= "localhost"
	port	= 8082
	user	= "postgres"
	password= "pratama21"
	dbname  = "db-go-sql"
)

var (
	db 	*sql.DB
	err error
)

type Employee struct {
	ID 			int
	Full_name	string
	Email		string
	Age			int
	Division 	string
}

func main() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s " + 
		"password=%s dbname=%s sslmode=disable", 
		host, port, user, password, dbname,
	)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("successfully connected to database")

	// CreateEmployee()
	// GetEmployees()
	// UpdateEmployee()
	// DeleteEmployee()
}

func CreateEmployee() {
	var employee = Employee{}

	sqlStatement := `
	INSERT INTO employees (full_name, email, age, division)
	VALUES	($1, $2, $3, $4)
	Returning *
	`

	err = db.QueryRow(
		sqlStatement, "Salma Alfarizi", "salman123@gmail.com", 22, "QA Engineer",
	).Scan(
		&employee.ID, 
		&employee.Full_name, 
		&employee.Email, 
		&employee.Age, 
		&employee.Division,
	)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data : %+v\n", employee)
}

func GetEmployees() {
	var results = []Employee{}

	sqlStatement := `SELECT * from employees`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var employee = Employee{}

		err = rows.Scan(
			&employee.ID,
			&employee.Full_name,
			&employee.Email,
			&employee.Age,
			&employee.Division,	
		)
		
		if err != nil {
			panic(err)
		}
		
		results = append(results, employee)
	}
	
	fmt.Println("Employee Data:", results)
}

func UpdateEmployee() {
	sqlStatement := `
	UPDATE employees
	SET full_name = $2, email = $3, division = $4, age = $5
	WHERE id = $1;	
	`

	res, err := db.Exec(
		sqlStatement, 1, "Abdul Hanif Pratama", "ahanifpratama@gmail.com", "DevOps", 22,
	)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Println("Updated Data Amount:", count)
}

func DeleteEmployee() {
	sqlStatement := `
	DELETE from employees
	WHERE id = $1;
	`

	res, err := db.Exec(sqlStatement, 1)
	if err != nil {
		panic(err)
	}
	
	count, err := res.RowsAffected() 
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted Data Amount:", count)
}