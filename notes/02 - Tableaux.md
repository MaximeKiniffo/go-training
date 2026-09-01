# Tableaux

## Idée

Un tableau possède une taille fixe qui fait partie de son type.

```go
languages := [3]string{"Go", "TypeScript", "JavaScript"}
```

Ici, `[3]string` ne désigne pas le même type que `[4]string`.

## À retenir

Un tableau convient lorsqu’on connaît exactement le nombre d’éléments. En pratique,
les slices sont beaucoup plus fréquentes.

## Comparaison TypeScript

Il ne correspond pas directement à `string[]`, car ce dernier peut changer de taille.

## Exercice associé

[[Exercices terminés#08 — Tableaux]]
