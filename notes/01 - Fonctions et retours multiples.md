# Fonctions et retours multiples

## À quoi cela sert

Donner un nom à un traitement réutilisable et obtenir un résultat.

```go
func presentation(name string, age int) string {
	return name + " a " + fmt.Sprint(age) + " ans"
}
```

Une fonction précise les types de ses paramètres et de sa valeur de retour.

## Retours multiples

Go peut retourner plusieurs valeurs, souvent une donnée et une erreur.

```go
func contact() (string, int) {
	return "Maxime", 20
}

name, age := contact()
```

## Comparaison JavaScript / TypeScript

Au lieu de retourner un objet ou un tableau à déconstruire, Go possède une syntaxe
naturelle pour plusieurs résultats. Leur ordre doit correspondre à la signature.

## À retenir

La signature d’une fonction est un contrat : nombre, ordre et type des paramètres et
des retours sont importants.

## Exercices associés

[[Exercices terminés#03 — Fonctions]] · [[Exercices terminés#06 — Retours multiples]]
