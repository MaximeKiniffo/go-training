# Range

## Idée

`range` parcourt une collection.

```go
for index, language := range languages {
	fmt.Println(index, language)
}
```

Pour ignorer une valeur, employer `_` :

```go
for _, language := range languages { }
```

## Attention avec les maps

L’ordre de parcours d’une map n’est pas garanti. Il ne faut pas baser une logique ou
un test sur cet ordre.

## Exercice associé

[[Exercices terminés#14 — Range]]
