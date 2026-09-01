# Pointeurs optionnels et nil

## Problème résolu

Une chaîne vide et une valeur absente ne veulent pas toujours dire la même chose. Un
pointeur peut distinguer ces deux situations.

```go
type Profile struct {
	Nickname *string
}
```

- `nil` : aucun surnom n’a été fourni ;
- pointeur vers `""` : un surnom vide a été fourni ;
- pointeur vers `"Max"` : un surnom existe.

## Vérification

```go
if profile.Nickname != nil {
	fmt.Println(*profile.Nickname)
}
```

## À retenir

Un pointeur n’est pas nécessaire pour toute donnée. Il devient pertinent lorsqu’il
faut modifier une valeur existante ou distinguer clairement l’absence d’une valeur.

## Exercice associé

[[Exercices terminés#30 — Pointeurs optionnels et nil]]
