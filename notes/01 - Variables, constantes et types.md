# Variables, constantes et types

## À quoi cela sert

Conserver des données : texte, nombre entier, nombre décimal ou booléen.

## Syntaxes utiles

```go
var language string = "Go"
years := 1
const maxUsers = 100
```

- `var` permet de déclarer explicitement une variable et, si besoin, son type.
- `:=` déclare une variable locale en laissant Go inférer son type.
- `const` représente une valeur qui ne changera pas.

## Types rencontrés

| Go | Rôle |
| --- | --- |
| `string` | texte |
| `int` | entier |
| `float64` | nombre décimal courant |
| `bool` | vrai ou faux |

## Comparaison JavaScript / TypeScript

`:=` ressemble à `const` ou `let`, mais Go infère un type statique. Une valeur
déclarée `int` ne peut pas devenir une chaîne plus tard.

## Pièges

- `:=` ne s’utilise pas pour affecter une variable déjà déclarée seule : employer `=`.
- une constante ne peut pas être réaffectée.

## Exercice associé

[[Exercices terminés#02 — Variables et constantes]]
