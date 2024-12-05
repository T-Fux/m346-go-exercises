package main

import "fmt"

type FullName struct {
	// TODO: add fields
	FirstName string
	LastName  string
}

// TODO: declare a structure for birth date

type BirthDate struct {
	dayOfBirth   int
	monthOfBirth int
	yearOfBirth  int
}

type Profile struct {
	// TODO: embed full name and birth date information
	Name             FullName
	Birth            BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		Name: FullName{
			FirstName: "Timo",
			LastName:  "Fuchs",
		},
		Birth: BirthDate{
			dayOfBirth:   14,
			monthOfBirth: 9,
			yearOfBirth:  2006,
		},
		NumberOfSiblings: 3,        // TODO: adjust
		ZodiacSign:       '\u264d', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
