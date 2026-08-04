package main

import "fmt"

func contact() (string, int) {
	return "Lina", 22
}

func main() {
	name, age := contact()
	fmt.Printf(`%s a %d ans`, name, age)
}
