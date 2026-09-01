# `errors.Is`

## À quoi cela sert

Déterminer si une erreur est, ou enveloppe, une erreur sentinelle précise.

```go
if errors.Is(err, errScoreNotFound) {
	// Proposer une création du score.
}
```

## Comportement

`errors.Is` fonctionne même si l’erreur a été enrichie avec `%w` par plusieurs
fonctions intermédiaires.

## Cas d’usage

Un appelant peut prendre une décision différente pour :

- une ressource absente ;
- une erreur de validation ;
- une indisponibilité technique.

## À retenir

La comparaison `err == errSentinelle` ne suffit pas lorsque l’erreur a été
enveloppée. Employer `errors.Is`.

## Exercice associé

[[Exercices terminés#35 — errors.Is]]
