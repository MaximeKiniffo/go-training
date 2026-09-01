# Type `error`

## Idée

`error` est une interface standard qui représente un échec. Une fonction peut
retourner une valeur utile et une erreur :

```go
func parseScore(input string) (int, error)
```

## Utilisation

```go
score, err := parseScore("12")
if err != nil {
	fmt.Println("Erreur :", err)
	return
}
fmt.Println(score)
```

`nil` signifie l’absence d’erreur.

## Comparaison JavaScript

Au lieu de compter sur `throw` et `catch` pour les erreurs attendues, Go les retourne
explicitement. Elles font partie du contrat de la fonction.

## À retenir

Ne pas continuer avec une valeur qui pourrait être invalide si `err != nil`.

## Exercice associé

[[Exercices terminés#31 — Type error]]
