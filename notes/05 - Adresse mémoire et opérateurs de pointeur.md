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

## Comment choisir

- Une adresse mémoire est l'emplacement d'une valeur ; en Go, `&value` obtient un
  pointeur qui contient cette adresse.
- Un pointeur a un type comme `*int` ou `*Quota`. On l'utilise lorsqu'une fonction
  doit lire ou modifier la valeur originale, ou lorsqu'une valeur peut être absente.
- `*pointer` déréférence le pointeur et accède à la valeur située à cette adresse.
- Pour une structure, `pointer.Field` est une écriture abrégée de
  `(*pointer).Field`.
- Une adresse affichée avec `%p` sert surtout à observer ou diagnostiquer le
  programme ; elle ne remplace pas un identifiant métier.

## Comparaison JavaScript

Les objets JavaScript sont souvent manipulés par référence, mais cette notion est
implicite. En Go, valeur et pointeur sont deux formes distinctes dont le type est
visible.

## Piège

Déférencer un pointeur `nil` provoque une erreur d’exécution. Toujours vérifier sa
présence lorsqu’il peut être absent.

## Exercices associés

[[Exercices terminés#26 — Adresse mémoire]] · [[Exercices terminés#27 — Opérateurs de pointeur]] · [[Exercices terminés#Exercice complémentaire — Pointeurs et adresse mémoire]]
