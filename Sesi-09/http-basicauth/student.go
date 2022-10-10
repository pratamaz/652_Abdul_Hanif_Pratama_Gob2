package main

var students = []*Student{}

type Student struct {
	Id		string
	Name	string
	Grade	string
}

func GetStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}

	return nil
}

func init() {
	students = append(students, &Student{Id: "s001", Name: "Abdul", Grade: 2})
	students = append(students, &Student{Id: "s002", Name: "Hanif", Grade: 2})
	students = append(students, &Student{Id: "s003", Name: "Pratama", Grade: 3})
}