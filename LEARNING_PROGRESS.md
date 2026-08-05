# Progression Go

## Niveau actuel

Débutant — phase 1 terminée.

## Notions étudiées

### Phase 0 — Installation et environnement

- [x] Installation de Go — Go 1.26.5 installé et vérifié avec `go version`.
- [x] Création d'un module Go et lecture de `go.mod`.
- [x] Exécution et formatage d'un programme avec `go run` et `go fmt`.

### Phase 1 — Fondamentaux

- [x] Programme exécutable — `package main` et fonction `main`.
- [x] Import d'un package standard — `fmt`.
- [x] Affichage avec `fmt.Println`.
- [x] Variables, constantes, types primitifs et inférence — `var`, `:=`, `const`.
- [x] Fonctions — paramètres, valeur de retour et appel d'une fonction.
- [x] Retours multiples.
- [x] Conditions — `if` et `else`.
- [x] `switch`.
- [x] Boucles `for`.

### Phase 2 — Collections et données

- [x] Tableaux.
- [x] Slices.
- [ ] `append`.
- [ ] `len` et `cap`.
- [ ] Copie de slices.
- [ ] Maps.
- [ ] Boucles avec `range`.
- [ ] Structs.
- [ ] Tags JSON.

### Phase 3 — Organisation du code

- [ ] Plusieurs fichiers.
- [ ] Packages locaux.
- [ ] Modules Go.
- [ ] `go.mod`.
- [ ] Fonctions exportées.
- [ ] Documentation.
- [ ] Organisation simple d'un projet.

### Phase 4 — Méthodes et abstraction

- [ ] Méthodes.
- [ ] Receivers par valeur.
- [ ] Receivers par pointeur.
- [ ] Composition.
- [ ] Interfaces.
- [ ] Interfaces implicites.
- [ ] Dépendances.

### Phase 5 — Pointeurs et mémoire

- [ ] Adresse mémoire.
- [ ] Opérateurs `&` et `*`.
- [ ] Passage par valeur.
- [ ] Modification d'une struct.
- [ ] Pointeurs optionnels.
- [ ] Valeurs `nil`.

### Phase 6 — Gestion des erreurs

- [ ] Type `error`.
- [ ] `errors.New`.
- [ ] `fmt.Errorf`.
- [ ] `%w`.
- [ ] `errors.Is`.
- [ ] `errors.As`.
- [ ] Propagation d'erreurs.
- [ ] Erreurs métier.
- [ ] Erreurs HTTP.

### Phase 7 — Tests

- [ ] Package `testing`.
- [ ] Premiers tests unitaires.
- [ ] Tests pilotés par tableaux.
- [ ] Sous-tests avec `t.Run`.
- [ ] Couverture.
- [ ] Tests des cas d'erreur.
- [ ] Commandes `go test` et `go test ./...`.

### Phase 8 — JSON, fichiers et HTTP

- [ ] Lecture et écriture de fichiers.
- [ ] Encodage et décodage JSON.
- [ ] `net/http`.
- [ ] Création d'un serveur.
- [ ] Handlers.
- [ ] Routes.
- [ ] Paramètres.
- [ ] Query parameters.
- [ ] Codes HTTP.
- [ ] Headers.
- [ ] Middlewares simples.
- [ ] Clients HTTP.
- [ ] Timeouts.
- [ ] Contexte.

### Phase 9 — Base de données

- [ ] SQL avec Go.
- [ ] Package `database/sql`.
- [ ] Connexion.
- [ ] Requêtes.
- [ ] Paramètres.
- [ ] `Scan`.
- [ ] Transactions.
- [ ] Repository.
- [ ] Migrations.
- [ ] Gestion des erreurs SQL.

### Phase 10 — Concurrence

- [ ] Différence entre concurrence et parallélisme.
- [ ] Goroutines.
- [ ] Channels.
- [ ] Channels bufferisés.
- [ ] `select`.
- [ ] `sync.WaitGroup`.
- [ ] Mutex.
- [ ] Race conditions.
- [ ] Race detector.
- [ ] Annulation.
- [ ] Timeouts.
- [ ] `context.Context`.

### Phase 11 — Service backend complet

- [ ] API HTTP.
- [ ] Handlers.
- [ ] Couche métier.
- [ ] Repository.
- [ ] Base PostgreSQL.
- [ ] Validations.
- [ ] Erreurs métier.
- [ ] Tests.
- [ ] Configuration par variables d'environnement.
- [ ] Logs structurés.
- [ ] Endpoint de santé.
- [ ] Arrêt propre du serveur.
- [ ] Dockerfile.

### Phase 12 — Introduction aux microservices

- [ ] Limites et responsabilités d'un service.
- [ ] Communication HTTP entre deux services.
- [ ] Clients HTTP.
- [ ] Timeouts.
- [ ] Propagation du contexte.
- [ ] Erreurs réseau.
- [ ] Retries.
- [ ] Idempotence.
- [ ] Base de données par service.
- [ ] Événements.
- [ ] Communication asynchrone.
- [ ] Docker Compose.
- [ ] Observabilité.
- [ ] Introduction à gRPC et Protocol Buffers.

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
- Boucles `for` : `fundamentals/05-loops/countdown.go`, avec un compte à rebours de 5 à 1.
- Retours multiples : `fundamentals/06-multiple-returns/contact.go`, avec une fonction retournant un prénom et un âge.
- `switch` : `fundamentals/07-switch/activity.go`, avec le choix d'une activité selon le jour.
- Tableaux : `fundamentals/08-arrays/languages.go`, avec un tableau fixe de trois langages et l'accès à son premier et son dernier élément.
- Slices : `fundamentals/09-slices/skills.go`, avec une liste de compétences modifiable et son affichage.

## Planning des mini-projets

- [ ] Après la phase 7 : démarrer le projet 1, programme en ligne de commande en mémoire, dans `mini-projects/`.
- [ ] Après la phase 8 : terminer le projet 1 avec la persistance dans un fichier JSON.
- [ ] Phase 11 : réaliser le projet 2, API REST monolithique, dans `services/`.
- [ ] Phase 12, après les appels HTTP inter-services : démarrer le projet 3, deux services.
- [ ] Phase 12, après les événements et la communication asynchrone : démarrer le projet 4, service de notification.

Règle de décision : lorsqu’un jalon est atteint, la section `Prochaine étape` doit indiquer explicitement le premier exercice du mini-projet concerné. Cette ligne reste la source de vérité pour démarrer une nouvelle séance.

## Organisation du dépôt

- Architecture d’apprentissage appliquée : chaque exercice est conservé dans un dossier dédié et les fichiers Go ont un nom descriptif.

## Prochaine étape

Phase 2 — découvrir `append` pour ajouter un élément à une slice.
