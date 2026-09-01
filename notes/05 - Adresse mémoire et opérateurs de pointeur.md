# Adresse mémoire et opérateurs de pointeur

## Les opérateurs

```go
score := 10
pointer := &score
*pointer = 20
```

- `&score` obtient l’adresse de `score`.
- `pointer` a le type `*int` : « pointeur vers un entier ».
- `*pointer` lit ou modifie la valeur à cette adresse.

Après la modification, `score` vaut `20`.

## Comparaison JavaScript

Les objets JavaScript sont souvent manipulés par référence, mais cette notion est
implicite. En Go, valeur et pointeur sont deux formes distinctes dont le type est
visible.

## Piège

Déférencer un pointeur `nil` provoque une erreur d’exécution. Toujours vérifier sa
présence lorsqu’il peut être absent.

## Exercices associés

[[Exercices terminés#26 — Adresse mémoire]] · [[Exercices terminés#27 — Opérateurs de pointeur]]
