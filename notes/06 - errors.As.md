# `errors.As`

## À quoi cela sert

Récupérer une erreur d’un type particulier présente dans une chaîne d’erreurs, afin
d’utiliser ses données.

Une erreur de plage de score peut, par exemple, contenir la valeur reçue et la limite
attendue. L’appelant récupère l’erreur typée puis choisit un comportement selon ces
informations.

## Différence avec `errors.Is`

| Besoin | Fonction |
| --- | --- |
| Savoir si une erreur correspond à une valeur sentinelle | `errors.Is` |
| Obtenir une erreur d’un type précis et ses champs | `errors.As` |

## Point technique essentiel

`errors.As` reçoit l’adresse d’une variable cible afin de pouvoir y placer l’erreur
trouvée. Le type de cette cible doit correspondre à l’erreur recherchée.

## Exercice associé

[[Exercices terminés#36 — errors.As]]
