# Tableau de bord Go

## Niveau actuel

Débutant avancé : phase 7 en cours.

Dernière notion acquise : [[07 - Package testing et premier test unitaire|Package `testing` et premier test unitaire]].

Dernier renforcement : [[05 - Adresse mémoire et opérateurs de pointeur|pointeurs et adresse mémoire]] avec l'exercice `quota.go`.

Notion en cours : exécuter les tests du module avec `go test ./...` et `go test -v`,
puis lire leurs sorties.

## Parcours

- [[01 - Fondamentaux]]
- [[02 - Collections et données]]
- [[03 - Organisation du code]]
- [[04 - Méthodes et abstraction]]
- [[05 - Pointeurs et mémoire]]
- [[06 - Gestion des erreurs]]
- [[07 - Package testing et premier test unitaire]]

Phase en cours : tests. Les erreurs HTTP seront abordées pendant la phase HTTP,
après les handlers et les codes de statut.

## Pour réviser efficacement

1. Lire une fiche de notion.
2. Ouvrir le fichier Go lié dans [[Exercices terminés]].
3. Expliquer à voix haute le cheminement du programme.
4. Noter les difficultés répétées dans [[Erreurs et pièges]].

## À retenir maintenant

Un test Go est une fonction `TestXxx` placée dans un fichier `*_test.go`. `go test`
compile le package et exécute ses tests.

## Liens utiles

- [[Commandes Go utiles]]
- [[Glossaire Go]]
- [[Erreurs et pièges]]
- [[À revoir]]
