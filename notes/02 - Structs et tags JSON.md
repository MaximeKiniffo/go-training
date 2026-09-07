# Structs et tags JSON

## Struct

Une struct regroupe des champs qui forment une même entité.

```go
type Developer struct {
	Name              string
	MainLanguage      string
	YearsOfExperience int
}
```

Elle est plus explicite et plus sûre qu’une `map[string]any` pour des données dont la
forme est connue.

## Tags JSON

> La syntaxe a été observée avec les structs dans l'ancien ordre du programme. Son
> utilisation concrète sera consolidée en phase 8A, après l'introduction de
> `encoding/json`.

```go
type Developer struct {
	Name string `json:"name"`
}
```

Le tag indique le nom que prendra le champ lors d’un encodage JSON. Il n’agit que si
le package `encoding/json` est utilisé.

## Visibilité

Un champ commençant par une majuscule est exporté et peut notamment être encodé par
`encoding/json`. Un champ en minuscule est interne au package.

## Exercices associés

[[Exercices terminés#15 — Structs]] · [[Exercices terminés#16 — Tags JSON]]
