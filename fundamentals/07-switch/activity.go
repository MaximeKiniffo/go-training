package main

import "fmt"

func main() {
	jour := "mercredi"

	switch jour {
	case "lundi", "mercredi":
		fmt.Println("Cours de Go")
	case "samedi":
		fmt.Println("Révision")
	default:
		fmt.Println("Repos ou autre activité")
	}
}
