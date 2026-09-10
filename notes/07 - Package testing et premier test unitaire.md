# Package `testing` et premier test unitaire

## Idées clés

- Le package standard `testing` permet de vérifier automatiquement le comportement d'une fonction.
- Un test est placé dans un fichier dont le nom se termine par `_test.go`.
- Une fonction de test commence par `Test` et reçoit `t *testing.T`.
- Lorsque la valeur obtenue ne correspond pas à la valeur attendue, `t.Errorf` signale l'échec du test.
- La commande `go test` compile le package courant et exécute ses tests.
- `go test -v` affiche le nom de chaque test exécuté et son résultat détaillé.
- `go test ./...` parcourt tous les packages du module ; `[no test files]` indique simplement qu'un package ne contient pas de test.

## Exemple étudié

L'exercice [[Exercices terminés#39 — Premier test unitaire|39 — Premier test unitaire]] vérifie que `SeatsAvailability(3, 5)` indique qu'une réservation est possible.

Fichiers associés :

- [seat_availability.go](../fundamentals/39-first-tests/seat_availability.go)
- [seat_availability_test.go](../fundamentals/39-first-tests/seat_availability_test.go)

## Tests des cas d’erreur

L’exercice [[Exercices terminés#40 — Tests des cas d’erreur|40 — Tests des cas d’erreur]]
vérifie trois résultats de `ValidateReservation` : un cas valide retourne `nil`, une
demande invalide retourne son erreur métier et une demande dépassant les places
disponibles retourne une autre erreur métier.

- Pour un cas valide, le test échoue si `err != nil`.
- Pour un cas d’erreur, `errors.Is(err, erreurAttendue)` vérifie l’identité de
  l’erreur sans dépendre de son texte.
- Chaque comportement est testé dans une fonction `TestXxx` séparée ; les sous-tests
  et les tests pilotés par tableaux seront étudiés plus tard.

Fichiers associés :

- [reservation_validation.go](../fundamentals/40-error-tests/reservation_validation.go)
- [reservation_validation_test.go](../fundamentals/40-error-tests/reservation_validation_test.go)

## Pièges

- Un fichier de test mal nommé, par exemple sans le suffixe `_test.go`, n'est pas exécuté par `go test`.
- Une fonction qui ne commence pas par `Test` n'est pas reconnue comme un test.
- Le test doit comparer le résultat obtenu avec le résultat attendu : un test qui appelle une fonction sans vérifier son résultat ne protège pas le comportement.
- Ne pas interpréter `[no test files]` comme un échec : cette ligne signale l'absence de fichier `*_test.go` dans le package concerné.
- Un cas valide ne doit pas être testé comme s’il devait retourner une erreur : il
  doit vérifier que `err` vaut `nil`.
- Ne pas comparer les textes d’erreur avec `err.Error()` lorsqu’une erreur nommée peut
  être vérifiée avec `errors.Is`.
