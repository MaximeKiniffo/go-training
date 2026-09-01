# Maps

## Idée

Une map associe une clé à une valeur.

```go
profile := map[string]string{
	"name":     "Maxime",
	"language": "Go",
}
profile["language"] = "TypeScript"
```

## Comparaison JavaScript / TypeScript

Une `map[string]string` est proche d’un `Record<string, string>` ou d’un `Map`, selon
l’usage. Les clés d’une map Go sont d’un type déterminé.

## Piège important

L’accès à une clé absente renvoie la valeur zéro du type. Pour distinguer une clé
absente d’une valeur réellement vide, utiliser la forme `value, ok := m[key]`.

## Exercice associé

[[Exercices terminés#13 — Maps]]
