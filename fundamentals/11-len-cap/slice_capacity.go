package main

import "fmt"

func main() {
	skills := make([]string, 0, 3)
	fmt.Println("Length:", len(skills))
	fmt.Println("Capacity:", cap(skills))

	skills = append(skills, "Go")
	fmt.Println("----------------------")
	fmt.Println("Skills:", skills)
	fmt.Println("Length:", len(skills))
	fmt.Println("Capacity:", cap(skills))

	skills = append(skills, "SQL")
	fmt.Println("----------------------")
	fmt.Println("Skills:", skills)
	fmt.Println("Length:", len(skills))
	fmt.Println("Capacity:", cap(skills))

	skills = append(skills, "Docker")
	fmt.Println("----------------------")
	fmt.Println("Skills:", skills)
	fmt.Println("Length:", len(skills))
	fmt.Println("Capacity:", cap(skills))

	skills = append(skills, "Git")
	fmt.Println("----------------------")
	fmt.Println("Skills:", skills)
	fmt.Println("Length:", len(skills))
	fmt.Println("Capacity:", cap(skills))
}
