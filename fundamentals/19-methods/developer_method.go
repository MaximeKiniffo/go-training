package main

import "fmt"

type Developer struct {
	Name         string
	MainLanguage string
}

func (developer Developer) Introduction() string {
	return "Je suis " + developer.Name + " et mon langage principal est " + developer.MainLanguage + "."
}

func main() {
	developer := Developer{Name: "Léa", MainLanguage: "Go"}
	fmt.Println(developer.Introduction())
}
