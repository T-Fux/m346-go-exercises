package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		FirstName string
		LastName  string
	}
	// TODO: declare a type for Class (consisting of multiple students)
	type Class struct {
		Name     string
		Students []Student
	}
	// TODO: declare a map of modules being attended by multiple classes
	type Modulverwaltung map[int][]Class

	student1 := Student{FirstName: "Timo", LastName: "Fuchs"}
	student2 := Student{FirstName: "Leandro", LastName: "Schwegler"}
	student3 := Student{FirstName: "Nando", LastName: "Schürmann"}

	student4 := Student{FirstName: "Simon", LastName: "Mettler"}
	student5 := Student{FirstName: "Artur", LastName: "Ferreira Cruz"}
	student6 := Student{FirstName: "Dylan", LastName: "Belamusto"}

	class1 := Class{Name: "Ina23cL", Students: []Student{student1, student2, student3}}
	class2 := Class{Name: "Ina23aL", Students: []Student{student4, student5, student6}}

	modules := Modulverwaltung{
		346: {class1, class2},
		114: {class1},
		117: {class2},
	}

	// TODO: output everything using fmt.Println()
	fmt.Println(modules)
}
