package main

import "fmt"

func Contact() (string, int) {
	return "Lina", 22
}

func main() {
	name, age := Contact()
	fmt.Printf(`%s a %d ans`, name, age)
}
