package main

import "fmt"

func main() {
	profile := map[string]string{
		"prenom":   "Maxime",
		"langage":  "TypeScript",
		"objectif": "Devenir développeur Go",
	}
	fmt.Println(profile["prenom"])
	fmt.Println(profile["langage"])
	fmt.Println(profile["objectif"])

	profile["langage"] = "Go"
	fmt.Println("----------------------")
	fmt.Println(profile["prenom"])
	fmt.Println(profile["langage"])
	fmt.Println(profile["objectif"])
}
