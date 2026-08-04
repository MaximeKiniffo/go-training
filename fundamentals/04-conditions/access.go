package main

import "fmt"

func main() {
	age := 99

	if age >= 18 {
		fmt.Println("Accès autorisé.")
	} else {
		fmt.Println("Accès refusé : réservé aux majeurs.")
	}
}
