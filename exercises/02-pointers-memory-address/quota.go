package main

import "fmt"

type Quota struct {
	User  string
	Limit int
	Used  int
}

func utiliserQuota(quota *Quota, remaining *int) {
	*remaining--
	if quota.Used >= quota.Limit {
		fmt.Println("Quota utilisé :", quota.Used)
		fmt.Println("Quota dépassé pour l'utilisateur", quota.User)
		fmt.Printf("Nouvelles requêtes à traiter : %v\n", *remaining)
		return
	}
	(*quota).Used++
	fmt.Println("Quota utilisé :", quota.Used)
	fmt.Printf("Requete acceptée\n")
	fmt.Printf("Nouvelles requêtes à traiter : %v\n", *remaining)
}

func main() {
	quota := Quota{
		User:  "atelier-go",
		Limit: 2,
		Used:  1,
	}

	remaining := 2

	fmt.Printf("Adresse mémoire avant : %p\n", &quota)
	fmt.Printf("Client : %v\n", quota.User)
	fmt.Printf("Limite : %v\n", quota.Limit)

	fmt.Printf("Requêtes déjà utilisées : %v\n", quota.Used)

	utiliserQuota(&quota, &remaining)
	utiliserQuota(&quota, &remaining)

	fmt.Printf("Adresse mémoire après : %p\n", &quota)
}
