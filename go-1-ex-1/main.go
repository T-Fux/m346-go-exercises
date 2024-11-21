package main

import (
	"fmt"
)

func main() {
	var firstName string = "Timo"
	lastName := "Fuchs"

	var dayOfBirth int = 14
	monthOfBirth := 9
	yearOfBirth := 2006

	var numberOfSiblings int = 3
	var heightInMeters float32 = 1.73
	var zodiacSign rune = 'v'

	// TODO: Declare and initialize the variables being used in the output!
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
