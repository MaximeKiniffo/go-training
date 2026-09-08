# Package `testing` et premier test unitaire

## Idées clés

- Le package standard `testing` permet de vérifier automatiquement le comportement d'une fonction.
- Un test est placé dans un fichier dont le nom se termine par `_test.go`.
- Une fonction de test commence par `Test` et reçoit `t *testing.T`.
- Lorsque la valeur obtenue ne correspond pas à la valeur attendue, `t.Errorf` signale l'échec du test.
- La commande `go test` compile le package courant et exécute ses tests.

## Exemple étudié

L'exercice [[Exercices terminés#39 — Premier test unitaire|39 — Premier test unitaire]] vérifie que `SeatsAvailability(3, 5)` indique qu'une réservation est possible.

Fichiers associés :

- [seat_availability.go](../fundamentals/39-first-tests/seat_availability.go)
- [seat_availability_test.go](../fundamentals/39-first-tests/seat_availability_test.go)

## Pièges

- Un fichier de test mal nommé, par exemple sans le suffixe `_test.go`, n'est pas exécuté par `go test`.
- Une fonction qui ne commence pas par `Test` n'est pas reconnue comme un test.
- Le test doit comparer le résultat obtenu avec le résultat attendu : un test qui appelle une fonction sans vérifier son résultat ne protège pas le comportement.
