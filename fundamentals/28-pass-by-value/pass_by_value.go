package main

import "fmt"

// Pass by value means that when you pass a variable to a function, the function receives a copy of the variable's value. Any changes made to the parameter inside the function do not affect the original variable outside the function.
func incrementByTen(number int) {
	number += 10
	fmt.Println("Dans la fonction :", number)
}

// In this example, we define a function called incrementByTen that takes an integer parameter number. Inside the function, we increment the value of number by 10 and print it. However, since we are passing the variable by value, the original variable in main remains unchanged.
func main() {
	number := 50
	incrementByTen(number)
	fmt.Println("Dans main:", number)
}
