# Progression Go

## Niveau actuel

Débutant — phase 6 en cours.

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
- [x] `append`.
- [x] `len` et `cap`.
- [x] Copie de slices.
- [x] Maps.
- [x] Boucles avec `range`.
- [x] Structs.
- [x] Tags JSON.

### Phase 3 — Organisation du code

- [x] Plusieurs fichiers d'un même package.
- [x] Packages locaux.
- [x] Modules Go.
- [x] `go.mod`.
- [x] Fonctions exportées.
- [x] Documentation.
- [x] Organisation simple d'un projet — `package main` coordonne le programme et les packages locaux regroupent une responsabilité précise.

### Phase 4 — Méthodes et abstraction

- [x] Méthodes.
- [x] Receivers par valeur — une méthode reçoit une copie de la struct et ne modifie donc pas la valeur d'origine.
- [x] Receivers par pointeur — une méthode recevant un pointeur peut modifier la struct d'origine.
- [x] Composition.
- [x] Interfaces.
- [x] Interfaces implicites — un type satisfait automatiquement une interface lorsqu'il possède toutes les méthodes attendues avec les signatures exactes.
- [x] Dépendances — une struct reçoit la collaboration nécessaire via un champ plutôt que de créer elle-même une implémentation concrète.

### Phase 5 — Pointeurs et mémoire

- [x] Adresse mémoire — une variable occupe un emplacement en mémoire, dont l'adresse est obtenue avec `&`.
- [x] Opérateurs `&` et `*` — `&` obtient l'adresse d'une variable et `*` permet de lire ou modifier la valeur visée par un pointeur.
- [x] Passage par valeur — une fonction reçoit une copie d'une valeur et la modification de cette copie ne modifie pas la variable originale.
- [x] Modification d'une struct — une struct passée par valeur est copiée ; un pointeur permet de modifier la struct originale.
- [x] Pointeurs optionnels — un pointeur permet de représenter une valeur qui peut être absente.
- [x] Valeurs `nil` — `nil` indique qu'un pointeur ne référence aucune valeur et doit être vérifié avant déréférencement.

### Phase 6 — Gestion des erreurs

- [x] Type `error` — une fonction peut retourner une erreur comme valeur, et l'appelant la vérifie avec `err != nil`.
- [x] `errors.New` — création d'une erreur simple avec un message et vérification de la valeur retournée avec `err != nil`.
- [x] `fmt.Errorf` — création d'une erreur dont le message contient une valeur dynamique.
- [x] `%w` — contextualiser une erreur avec `fmt.Errorf` tout en conservant l'erreur d'origine.
- [x] `errors.Is` — vérifier si une erreur correspond à une erreur d'origine, même lorsqu'elle est enveloppée avec `%w`.
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
- `append` : `fundamentals/10-append/skills_append.go`, avec l’ajout de `"Go"` à une slice de langages.
- `len` et `cap` : `fundamentals/11-len-cap/slice_capacity.go`, avec l’observation de la longueur et de la capacité lors d’ajouts successifs.
- Copie de slices : `fundamentals/12-slice-copy/independent_skills.go`, avec une copie indépendante modifiée sans altérer la slice originale.
- Maps : `fundamentals/13-maps/profile_map.go`, avec un profil associant des clés à des valeurs et la mise à jour du langage.
- Boucles avec `range` : `fundamentals/14-range/languages_range.go`, avec le parcours d'une slice de langages et l'affichage de chaque index et valeur.
- Structs : `fundamentals/15-structs/developer_profile.go`, avec un profil de développeur regroupant un nom, un langage principal et des années d'expérience.
- Tags JSON : `fundamentals/16-json-tags/developer_json_tags.go`, avec les noms JSON `name`, `main_language` et `years_of_experience` associés aux champs d'un profil de développeur.
- Plusieurs fichiers : `fundamentals/17-multiple-files/presentation.go` et `application.go`, avec une fonction appelée depuis un autre fichier du même package `main`.
- Packages locaux : `fundamentals/18-local-packages/application.go` et `profile/profile.go`, avec l'import du package `profile` via le chemin du module `go-training`.
- Modules Go et `go.mod` : compréhension du nom de module comme préfixe des imports locaux.
- Fonctions exportées : `profile.Introduction` et `profile.Welcome`, appelées depuis `package main` ; une majuscule initiale rend une fonction accessible hors de son package. Cette notion a été abordée avec les packages locaux dans le même exercice.
- Documentation GoDoc : commentaires placés au-dessus des fonctions exportées `Introduction` et `Welcome`, commençant par leur nom.
- Organisation simple d'un projet : distinction entre le package `main`, qui appelle les packages, et le package `profile`, responsable des fonctionnalités liées au profil.
- Méthodes avec receiver par valeur : `fundamentals/19-methods/developer_method.go`, avec `Developer.Introduction()` appelée sur une valeur `Developer`.
- Receivers par valeur : `fundamentals/20-value-receivers/developer_experience.go`, avec `Developer.GainExperience()` qui modifie seulement sa copie locale.
- Receivers par pointeur : `fundamentals/21-pointer-receivers/developer_promotion.go`, avec `Developer.Promote()` qui incrémente l'expérience de la struct d'origine.
- Composition : `fundamentals/22-composition/developer_contact.go`, avec un `Developer` contenant un `Contact` pour regrouper des coordonnées liées.
- Interfaces : `fundamentals/23-interfaces/presentable.go`, avec une interface `Presentable` utilisée par la fonction `afficher` pour accepter un `Developer` ou un `Project`.
- Interfaces implicites : `fundamentals/24-implicit-interfaces/notification.go`, avec `Email` et `SMS` transmis à `sendNotification` grâce à leur méthode `Notify() string`.
- Dépendances : `fundamentals/25-dependencies/task_notification.go`, avec `TaskService` recevant un `Notifier` et restant indépendant de `ConsoleNotifier`.
- Adresse mémoire : `fundamentals/26-memory-address/address.go`, avec l'affichage des valeurs `langage` et `yearsOfExperience` ainsi que de leurs adresses via `&`.

- `fundamentals/27-pointer-operators/pointer_operators.go`, avec la lecture et la modification d'une variable via un pointeur.
- `fundamentals/28-pass-by-value/pass_by_value.go`, avec une fonction qui modifie une copie d'un score sans modifier la variable originale.
- `fundamentals/29-struct-modification/profile.go`, avec la modification de l'expérience d'une struct `Developer` grâce à un pointeur.
- `fundamentals/30-optional-pointers/optional_profile.go`, avec une struct `Profile` dont le surnom peut être absent grâce à un pointeur `nil`.
- `fundamentals/31-errors/score_parser.go`, avec une fonction `parseScore` retournant une erreur de conversion et sa vérification dans `main`.
- `fundamentals/32-errors-new/age_validation.go`, avec une validation d'âge retournant une erreur créée avec `errors.New` lorsque l'âge est inférieur à 18 ans.
- `fundamentals/33-fmt-error/score_validation.go`, avec une validation de score retournant une erreur créée avec `fmt.Errorf` et contenant la valeur invalide.
- `fundamentals/34-error-wrapping/score_loading.go`, avec une erreur d'origine contextualisée grâce à `fmt.Errorf` et `%w`.
- `fundamentals/35-errors-is/score_loading.go`, avec une erreur sentinelle détectée grâce à `errors.Is` malgré son enveloppage avec `%w`.

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

Phase 6 — `errors.As` : récupérer une erreur d'un type précis dans une chaîne d'erreurs.
