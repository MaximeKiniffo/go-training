# `fmt.Errorf`

## À quoi cela sert

Créer un message d’erreur qui contient une valeur dynamique.

```go
return fmt.Errorf("score invalide : %d", score)
```

Le message précise ici quel score est invalide.

## Choisir entre `errors.New` et `fmt.Errorf`

| Situation | Outil |
| --- | --- |
| Message constant | `errors.New` |
| Message avec une valeur | `fmt.Errorf` |
| Ajouter un contexte à une erreur existante | `fmt.Errorf` avec `%w` |

## Attention

`fmt.Errorf` avec `%v` affiche une erreur existante mais ne préserve pas sa relation
dans la chaîne ; utiliser `%w` lorsqu’elle doit rester détectable.

## Exercice associé

[[Exercices terminés#33 — fmt.Errorf]]
