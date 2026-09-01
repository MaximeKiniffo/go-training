# Méthodes

## Idée

Une méthode est une fonction associée à un type grâce à un receiver.

```go
type Developer struct {
	Name string
}

func (d Developer) Introduction() string {
	return "Je suis " + d.Name
}
```

Elle s’appelle avec `developer.Introduction()`.

## Comparaison TypeScript

Cela ressemble à une méthode de classe, mais Go n’a pas de classes. La méthode est
déclarée séparément du type et son receiver est une valeur explicite.

## À retenir

Une méthode est pertinente lorsque son comportement a du sens pour le type concerné.

## Exercice associé

[[Exercices terminés#19 — Méthodes]]
