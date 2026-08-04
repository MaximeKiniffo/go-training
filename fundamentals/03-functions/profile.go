package main

import "fmt"

func presentation(nom string, age int) string {
	return fmt.Sprintf("Je m'appelle %s et j'ai %d ans.", nom, age)
}

func addition(a int, b int) int {
	return a + b
}

func multiplication(a int, b int) int {
	return a * b
}

func main() {
	fmt.Println(presentation("Maxime", 33))
	fmt.Println(addition(8, 4))
	fmt.Println(multiplication(8, 4))
}
