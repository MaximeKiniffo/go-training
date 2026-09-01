# Conditions et switch

## À quoi cela sert

Faire varier le comportement du programme selon une situation.

```go
if age >= 18 {
	fmt.Println("Accès autorisé")
} else {
	fmt.Println("Accès refusé")
}
```

`switch` est utile lorsqu’une même valeur peut correspondre à plusieurs cas :

```go
switch day {
case "samedi", "dimanche":
	fmt.Println("Repos")
default:
	fmt.Println("Travail")
}
```

## Différence importante avec JavaScript

Les parenthèses autour de la condition ne sont pas nécessaires. Un `case` Go ne
continue pas automatiquement dans le suivant : pas de `break` à ajouter.

## À retenir

Une branche doit représenter une décision différente, et non uniquement deux textes
presque identiques.

## Exercices associés

[[Exercices terminés#04 — Conditions]] · [[Exercices terminés#07 — Switch]]
