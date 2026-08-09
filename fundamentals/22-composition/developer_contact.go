package main

import "fmt"

type Contact struct {
	Email string
	City  string
}

type Developer struct {
	Name         string
	MainLanguage string
	Contact      Contact
}

func main() {
	developer := Developer{
		Name:         "Maxime",
		MainLanguage: "Go",
		Contact: Contact{
			Email: "developer@example.com",
			City:  "Paris",
		},
	}
	fmt.Printf("%s développe en %s. Contact: %s, %s", developer.Name, developer.MainLanguage, developer.Contact.Email, developer.Contact.City)
}
