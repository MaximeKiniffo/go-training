# Interfaces implicites

## Idée

Une interface décrit les méthodes attendues. Un type la satisfait automatiquement
lorsqu’il possède ces méthodes avec les signatures exactes.

```go
type Notifier interface {
	Notify() string
}
```

`Email` et `SMS` n’ont pas besoin d’écrire qu’ils « implémentent » `Notifier` : leur
méthode `Notify() string` suffit.

## Intérêt

Une fonction dépend d’un comportement, pas d’un type concret : elle peut recevoir
plusieurs implémentations compatibles.

## Comparaison TypeScript

TypeScript est aussi structurel dans beaucoup de cas. En Go, cette souplesse est
centrale : le type concret ne dépend pas nécessairement du package qui définit
l’interface.

## Piège

Ne pas créer une interface sans besoin réel de plusieurs comportements ou d’une
dépendance à isoler.

## Exercices associés

[[Exercices terminés#23 — Interface]] · [[Exercices terminés#24 — Interface implicite]]
