# Programme Go, packages et affichage

## À quoi cela sert

Créer et exécuter un premier programme.

## Éléments indispensables

- `package main` indique un programme exécutable.
- `func main()` est son point de départ.
- `import "fmt"` donne accès aux fonctions de formatage et d’affichage.
- `fmt.Println` affiche une ou plusieurs valeurs puis un retour à la ligne.

## Exemple minimal

```go
package main

import "fmt"

func main() {
	fmt.Println("Bonjour Go")
}
```

Exécution : `go run .` depuis le dossier contenant le programme ou le module.

## Comparaison JavaScript

`func main()` est le point d’entrée explicite. Contrairement à un script Node.js,
le fichier est compilé avant d’être lancé.

## À retenir

Chaque fichier Go commence par une déclaration de package. Les imports doivent être
utilisés, sinon le compilateur signale une erreur.

## Exercice associé

[[Exercices terminés#01 — Premier programme]]
