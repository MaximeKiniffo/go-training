# Receivers par valeur et par pointeur

## Receiver par valeur

```go
func (d Developer) GainExperience() {
	d.YearsOfExperience++
}
```

`d` est une copie : la struct d’origine ne change pas.

## Receiver par pointeur

```go
func (d *Developer) Promote() {
	d.YearsOfExperience++
}
```

`d` désigne la struct d’origine : la modification persiste.

## Choisir

Utiliser un receiver par pointeur lorsqu’une méthode doit modifier la struct ou qu’il
serait coûteux de la copier. Rester cohérent pour les méthodes d’un même type.

Go insère souvent automatiquement `&` lors de l’appel sur une valeur adressable, mais
il faut comprendre que la méthode travaille alors sur l’original.

## Exercices associés

[[Exercices terminés#20 — Receiver par valeur]] · [[Exercices terminés#21 — Receiver par pointeur]]
