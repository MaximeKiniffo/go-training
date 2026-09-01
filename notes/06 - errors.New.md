# `errors.New`

## À quoi cela sert

Créer une erreur simple avec un message constant.

```go
var errAgeTooLow = errors.New("l’âge doit être au moins 18 ans")
```

Cette forme convient lorsque le message ne dépend pas d’une donnée dynamique.

## Quand l’utiliser

- règle simple et stable ;
- erreur sentinelle qui pourra être comparée plus tard avec `errors.Is`.

## Piège

Créer plusieurs erreurs avec le même texte ne les rend pas identiques. Pour une
erreur sentinelle, déclarer une valeur partagée une seule fois.

## Exercice associé

[[Exercices terminés#32 — errors.New]]
