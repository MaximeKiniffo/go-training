package main

import "fmt"

func main() {
	languages := []string{"JavaScript", "TypeScript", "Go"}
	for i, language := range languages {
		fmt.Printf("Index: %d, Language: %s\n", i, language)
	}
}
