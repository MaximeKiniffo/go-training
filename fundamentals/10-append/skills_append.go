package main

import "fmt"

func main() {
	skills := []string{"JavaScript", "TypeScript"}
	skills = append(skills, "Go")

	fmt.Println(skills)
}
