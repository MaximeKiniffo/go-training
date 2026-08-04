package main

func main() {
	jour := "mercredi"

	switch jour {
	case "lundi", "mercredi":
		println("Cours de Go")
	case "samedi":
		println("Révision")
	default:
		println("Repos ou autre activité")
	}
}
