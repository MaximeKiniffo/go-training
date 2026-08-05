package main

import "fmt"

func main() {
	languages := []string{"JavaScript", "TypeScript", "Go"}
	languages[2] = "Go Backend"

	fmt.Printf("%v\n", languages[0])
	fmt.Println(languages)

}
