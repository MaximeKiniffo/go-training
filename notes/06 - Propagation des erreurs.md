# Propagation des erreurs

## Problème résolu

Une fonction intermédiaire ne sait pas toujours comment réagir à une erreur. Elle la
retourne donc à son appelant, qui dispose peut-être d’un meilleur contexte pour la
traiter.

```go
result, err := convert(input)
if err != nil {
	return err
}
return result, nil
```

## Cheminement à suivre

1. une fonction basse retourne une erreur ;
2. chaque appelant vérifie `err` immédiatement ;
3. il la traite, l’enveloppe avec `%w`, ou la retourne ;
4. le niveau qui connaît l’action à mener décide finalement du comportement.

## À retenir

Propager une erreur ne veut pas dire l’ignorer : la fonction signale explicitement
qu’elle n’a pas produit de résultat valide.

## Point de vigilance personnel

Dans un exercice précédent, distinguer correctement la chaîne d’entrée du nombre
converti demandait une attention particulière. Nommer clairement les variables aide à
suivre l’algorithme et le chemin de `err`.

## Exercices associés

[[Exercices terminés#37 — Propagation]] · [[Exercices terminés#Exercice complémentaire — Rapport]]
