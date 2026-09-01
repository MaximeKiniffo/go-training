# Composition

## Idée

Une struct peut contenir une autre struct pour représenter une relation « possède un ».

```go
type Contact struct {
	Email string
}

type Developer struct {
	Name    string
	Contact Contact
}
```

`Developer` possède un `Contact`. La composition évite une hiérarchie de classes qui
ne correspondrait pas au domaine.

## Comparaison TypeScript

C’est proche d’un objet imbriqué typé. En Go, la composition est souvent privilégiée
à l’héritage, qui n’existe pas comme dans les langages orientés objet classiques.

## Exercice associé

[[Exercices terminés#22 — Composition]]
