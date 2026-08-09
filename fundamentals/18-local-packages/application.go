package main

import (
	"fmt"
	"go-training/fundamentals/18-local-packages/profile"
)

// Ce programme illustre l'import d'un package local et l'appel de fonctions exportées.
func main() {
	fmt.Println(profile.Introduction("Maxime", "Go"))
	fmt.Println(profile.Welcome("Maxime"))
}
