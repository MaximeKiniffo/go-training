package main

import "fmt"

func main() {
	// Declare an array of strings with a length of 3
	languages := [3]string{"JavaScript", "TypeScript", "Go"}
	fmt.Printf("Premier langage : %s\n", languages[0])
	fmt.Printf("Dernier langage : %s\n", languages[2])
}
