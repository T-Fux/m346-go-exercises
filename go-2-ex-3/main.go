package main

import "fmt"

func main() {
	// TODO: create a map called "modules"

	modules := map[int]string{
		104: "Datenmodell implementieren",
		117: "Informatik- und Netzinfrastruktur für ein kleines Unternehmen realisieren",
		346: "Cloud Lösungen konzipieren und realisieren",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	delete(modules, 104)
	// TODO: add one
	modules[320] = "Objektorientiert Programmieren"
	// TODO: replace one
	modules[346] = "Cloud-Architekturen gestalten"
	fmt.Println(modules)
}
