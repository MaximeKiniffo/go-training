package main

import "fmt"

type Presentable interface {
	Description() string
}

type Developer struct {
	Name string
}

type Project struct {
	Title string
}

func (d Developer) Description() string {
	return fmt.Sprintf("Développeur: %s", d.Name)
}

func (p Project) Description() string {
	return fmt.Sprintf("Projet: %s", p.Title)
}

func afficher(item Presentable) {
	fmt.Println(item.Description())
}

func main() {
	dev := Developer{Name: "Maxime"}
	proj := Project{Title: "Interface en Go"}
	afficher(dev)
	afficher(proj)
}
