# Progression Go

## Niveau actuel

Débutant — phase 1 en cours.

## Notions étudiées

### Phase 0 — Installation et environnement

- [x] Installation de Go — Go 1.26.5 installé et vérifié avec `go version`.
- [x] Création d'un module Go et lecture de `go.mod`.
- [x] Exécution et formatage d'un programme avec `go run` et `go fmt`.

### Phase 1 — Fondamentaux

- [x] Programme exécutable — `package main` et fonction `main`.
- [x] Import d'un package standard — `fmt`.
- [x] Affichage avec `fmt.Println`.
- [x] Variables et constantes — `var`, `:=`, `const`.
- [x] Fonctions — paramètres, valeur de retour et appel d'une fonction.
- [x] Conditions — `if` et `else`.
- [ ] Boucles `for`.

### Phase 2 — Collections et données

- [ ] Tableaux et slices.
- [ ] Maps.
- [ ] Structs.

### Phase 3 — Organisation du code

- [ ] Plusieurs fichiers dans un même package.
- [ ] Packages locaux.
- [ ] Visibilité des identifiants — majuscule/minuscule.
- [ ] Organisation simple d'un projet.

### Phase 4 — Méthodes et abstraction

- [ ] Méthodes et receivers par valeur.
- [ ] Receivers par pointeur.
- [ ] Composition.
- [ ] Interfaces implicites et dépendances.

### Phase 5 — Pointeurs et mémoire

- [ ] Adresses mémoire et opérateurs `&` et `*`.
- [ ] Passage par valeur et modification d'une struct.
- [ ] Pointeurs optionnels et valeurs `nil`.

### Phase 6 — Gestion des erreurs

- [ ] Type `error`, `errors.New` et `fmt.Errorf`.
- [ ] Wrapping avec `%w`, `errors.Is` et `errors.As`.
- [ ] Propagation et erreurs métier.

### Phase 7 — Tests

- [ ] Package `testing` et premiers tests unitaires.
- [ ] Tests pilotés par tableaux et sous-tests.
- [ ] Couverture et tests des cas d'erreur.

### Phase 8 — JSON, fichiers et HTTP

- [ ] Lecture et écriture de fichiers.
- [ ] Encodage et décodage JSON.
- [ ] Serveur HTTP, handlers et routes avec `net/http`.
- [ ] Paramètres, codes HTTP, headers et middlewares simples.
- [ ] Clients HTTP, timeouts et contexte.

### Phase 9 — Base de données

- [ ] SQL avec `database/sql`.
- [ ] Requêtes, paramètres, `Scan` et transactions.
- [ ] Repository, migrations et gestion des erreurs SQL.

### Phase 10 — Concurrence

- [ ] Goroutines et channels.
- [ ] Channels bufferisés et `select`.
- [ ] `sync.WaitGroup`, mutex et race conditions.
- [ ] Annulation, timeouts et `context.Context`.

### Phase 11 — Service backend complet

- [ ] API HTTP, handlers et couche métier.
- [ ] Repository PostgreSQL, validations et erreurs métier.
- [ ] Tests et configuration par variables d'environnement.
- [ ] Logs structurés, endpoint de santé et arrêt propre.
- [ ] Dockerfile.

### Phase 12 — Introduction aux microservices

- [ ] Limites et responsabilités d'un service.
- [ ] Communication HTTP entre services, contexte et timeouts.
- [ ] Erreurs réseau, retries et idempotence.
- [ ] Base de données par service et communication asynchrone.
- [ ] Docker Compose, observabilité, gRPC et Protocol Buffers.

## Difficultés rencontrées

- La commande `go` n'était pas disponible dans une fenêtre PowerShell ouverte avant la mise à jour du `PATH`. Résolu en ouvrant une nouvelle fenêtre.

## Erreurs récurrentes

Aucune pour le moment.

## Exercices terminés

- Création du module `go-training` et lecture de `go.mod`.
- Premier programme : `fundamentals/01-first-program/hello.go`, formaté avec `go fmt` et exécuté avec `go run`.
- Variables et constantes : `fundamentals/02-variables/variables.go`, avec affichage d'un profil via `fmt.Println`.
- Fonctions : `fundamentals/03-functions/profile.go`, avec une fonction `presentation` qui reçoit des paramètres et retourne une chaîne.
- Conditions : `fundamentals/04-conditions/access.go`, avec un contrôle d'accès selon l'âge.

## Organisation du dépôt

- Architecture d’apprentissage appliquée : chaque exercice est conservé dans un dossier dédié et les fichiers Go ont un nom descriptif.

## Prochaine étape

Phase 1 — découvrir les boucles `for`.
