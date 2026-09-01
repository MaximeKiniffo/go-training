# Longueur, capacité et copie de slices

## `len` et `cap`

- `len(slice)` : nombre d’éléments accessibles.
- `cap(slice)` : espace disponible dans le tableau sous-jacent avant une éventuelle
  réallocation.

La capacité peut évoluer lors d’un `append`, selon la stratégie du runtime Go. Il ne
faut pas dépendre d’une valeur précise après un agrandissement.

## Copier une slice

```go
original := []string{"Go", "SQL"}
copyOfOriginal := make([]string, len(original))
copy(copyOfOriginal, original)
```

La nouvelle slice peut alors être modifiée sans changer `original`.

## À retenir

Une affectation simple, `other := original`, ne fait pas une copie indépendante des
éléments : les deux slices peuvent référencer le même stockage.

## Exercices associés

[[Exercices terminés#11 — Len et cap]] · [[Exercices terminés#12 — Copie de slices]]
