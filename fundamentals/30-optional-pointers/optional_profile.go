package main

import "fmt"

type Profile struct {
	Name     string
	Nickname *string
}

func main() {
	profile := Profile{
		Name:     "Alice",
		Nickname: nil,
	}
	fmt.Println("Nom : ", profile.Name)
	if profile.Nickname != nil {
		fmt.Println("Surnom : ", *profile.Nickname)
	} else {
		fmt.Println("Surnom : Aucun surnom")
	}
	optionalNickname := "Jean Kevin"

	secondProfile := Profile{
		Name:     "Bob",
		Nickname: &optionalNickname,
	}
	fmt.Println("Nom : ", secondProfile.Name)
	if secondProfile.Nickname != nil {
		fmt.Println("Surnom : ", *secondProfile.Nickname)
	} else {
		fmt.Println("Surnom : Aucun surnom")
	}
}
