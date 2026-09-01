# Envelopper une erreur avec `%w`

## Problème résolu

Une erreur technique peut être compréhensible mais manquer de contexte : « quelle
étape du programme a échoué ? »

```go
return fmt.Errorf("chargement du score : %w", err)
```

Le message gagne du contexte tout en conservant l’erreur d’origine.

## Chaîne d’erreurs

L’erreur enveloppée devient une nouvelle erreur qui référence l’ancienne. Des appels
ultérieurs à `errors.Is` et `errors.As` peuvent parcourir cette chaîne.

## Piège

`%v` formate seulement le message. `%w` crée une relation d’enveloppement.

## Exercice associé

[[Exercices terminés#34 — Error wrapping]]
