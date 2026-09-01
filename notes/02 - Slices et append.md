# Slices et append

## Idée

Une slice est une vue flexible sur une séquence de valeurs.

```go
skills := []string{"Git", "SQL"}
skills = append(skills, "Go")
```

`append` retourne la slice résultante : il faut conserver son retour.

## Comparaison TypeScript

`[]string` est proche de `string[]`. La différence importante est que la capacité et
le stockage sous-jacent font partie du comportement des slices Go.

## Pièges

- oublier `skills =` devant `append` ;
- croire que deux slices partageant le même tableau sont toujours indépendantes.

## Exercices associés

[[Exercices terminés#09 — Slices]] · [[Exercices terminés#10 — Append]]
