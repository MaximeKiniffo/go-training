package main

import "fmt"

func main() {
	skills := []string{"JavaScript", "TypeScript", "Go"}
	fmt.Println("Longueur de skills:", len(skills))
	fmt.Println("Capacité de skills:", cap(skills))
	fmt.Println("Original skills:", skills)
	copySkills := make([]string, len(skills))

	fmt.Println("Longueur de copySkills:", len(copySkills))
	fmt.Println("Capacité de copySkills:", cap(copySkills))
	fmt.Println("Before copySkills:", copySkills)

	copy(copySkills, skills)
	fmt.Println("After copySkills:", copySkills)
	copySkills[2] = "SQL"

	fmt.Println("----------------------")
	fmt.Println("Original skills:", skills)
	fmt.Println("Copy skills:", copySkills)
}
