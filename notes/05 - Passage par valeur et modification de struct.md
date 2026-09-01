# Passage par valeur et modification de struct

## Passage par valeur

Les paramètres Go sont transmis par valeur. Une fonction qui reçoit un `int` ou une
struct reçoit une copie de cette valeur.

```go
func increment(score int) {
	score++
}
```

L’appel à `increment(score)` ne modifie pas la variable de l’appelant.

## Modifier une struct d’origine

Passer un pointeur permet de modifier la même struct :

```go
func promote(developer *Developer) {
	developer.YearsOfExperience++
}
```

## À retenir

Go passe toujours les arguments par valeur. Quand l’argument est un pointeur, la
valeur copiée est l’adresse ; elle permet encore de viser la donnée initiale.

## Exercices associés

[[Exercices terminés#28 — Passage par valeur]] · [[Exercices terminés#29 — Modification de struct]]
