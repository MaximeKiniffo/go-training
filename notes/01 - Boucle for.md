# Boucle for

## À quoi cela sert

Répéter une action selon une condition ou un compteur.

```go
for count := 5; count >= 1; count-- {
	fmt.Println(count)
}
```

Go n’a qu’un seul mot-clé de boucle : `for`.

## Formes courantes

```go
for i := 0; i < 3; i++ { }
for condition { }
for { }
```

La dernière forme est une boucle infinie : elle exige généralement un `break` ou un
`return` pour pouvoir s’arrêter.

## Comparaison JavaScript

Il n’existe pas de `while` en Go ; `for condition` remplit ce rôle.

## Exercice associé

[[Exercices terminés#05 — Boucle for]]
