---
aliases:
  - Exercices terminés
---

# Exercices terminés

Les fichiers Go sont les sources de vérité. Ces liens servent à retrouver rapidement
le code correspondant à une fiche de notion.

## Phase 0

- Module et premier programme : [hello.go](../fundamentals/01-first-program/hello.go)

## Phase 1 — Fondamentaux

### 01 — Premier programme

[hello.go](../fundamentals/01-first-program/hello.go) — `package main`, `main`, `fmt.Println`.

### 02 — Variables et constantes

[variables.go](../fundamentals/02-variables/variables.go) — profil affiché avec des valeurs typées.

### 03 — Fonctions

[profile.go](../fundamentals/03-functions/profile.go) — paramètres et chaîne retournée.

### 04 — Conditions

[access.go](../fundamentals/04-conditions/access.go) — accès selon l’âge.

### 05 — Boucle `for`

[countdown.go](../fundamentals/05-loops/countdown.go) — compte à rebours de 5 à 1.

### 06 — Retours multiples

[contact.go](../fundamentals/06-multiple-returns/contact.go) — prénom et âge retournés.

### 07 — `switch`

[activity.go](../fundamentals/07-switch/activity.go) — activité selon le jour.

## Phase 2 — Collections et données

### 08 — Tableaux

[languages.go](../fundamentals/08-arrays/languages.go) — tableau fixe de langages.

### 09 — Slices

[skills.go](../fundamentals/09-slices/skills.go) — liste de compétences modifiable.

### 10 — `append`

[skills_append.go](../fundamentals/10-append/skills_append.go) — ajout de `Go` à une slice.

### 11 — `len` et `cap`

[slice_capacity.go](../fundamentals/11-len-cap/slice_capacity.go) — évolution longueur/capacité.

### 12 — Copie de slices

[independent_skills.go](../fundamentals/12-slice-copy/independent_skills.go) — copie indépendante.

### 13 — Maps

[profile_map.go](../fundamentals/13-maps/profile_map.go) — profil clé/valeur.

### 14 — `range`

[languages_range.go](../fundamentals/14-range/languages_range.go) — index et valeur d’une slice.

### 15 — Structs

[developer_profile.go](../fundamentals/15-structs/developer_profile.go) — données regroupées dans `Developer`.

### 16 — Tags JSON

[developer_json_tags.go](../fundamentals/16-json-tags/developer_json_tags.go) — noms JSON des champs.

## Phase 3 — Organisation du code

### 17 — Plusieurs fichiers

[presentation.go](../fundamentals/17-multiple-files/presentation.go) et [application.go](../fundamentals/17-multiple-files/application.go) — même package `main`.

### 18 — Packages locaux et module

[application.go](../fundamentals/18-local-packages/application.go) et [profile.go](../fundamentals/18-local-packages/profile/profile.go) — import d’un package local et fonctions exportées.

## Phase 4 — Méthodes et abstraction

### 19 — Méthodes

[developer_method.go](../fundamentals/19-methods/developer_method.go) — méthode de présentation.

### 20 — Receiver par valeur

[developer_experience.go](../fundamentals/20-value-receivers/developer_experience.go) — modification d’une copie.

### 21 — Receiver par pointeur

[developer_promotion.go](../fundamentals/21-pointer-receivers/developer_promotion.go) — modification de la struct originale.

### 22 — Composition

[developer_contact.go](../fundamentals/22-composition/developer_contact.go) — `Developer` contient `Contact`.

### 23 — Interface

[presentable.go](../fundamentals/23-interfaces/presentable.go) — comportement commun à plusieurs types.

### 24 — Interface implicite

[notification.go](../fundamentals/24-implicit-interfaces/notification.go) — `Email` et `SMS` satisfont `Notifier`.

### 25 — Dépendances

[task_notification.go](../fundamentals/25-dependencies/task_notification.go) — service dépendant d’un contrat.

## Phase 5 — Pointeurs et mémoire

### 26 — Adresse mémoire

[address.go](../fundamentals/26-memory-address/address.go) — valeurs et adresses avec `&`.

### 27 — Opérateurs de pointeur

[pointer_operators.go](../fundamentals/27-pointer-operators/pointer_operators.go) — lecture et modification via `*`.

### 28 — Passage par valeur

[pass_by_value.go](../fundamentals/28-pass-by-value/pass_by_value.go) — modification d’une copie.

### 29 — Modification de struct

[profile.go](../fundamentals/29-struct-modification/profile.go) — struct modifiée via pointeur.

### 30 — Pointeurs optionnels

[optional_profile.go](../fundamentals/30-optional-pointers/optional_profile.go) — surnom éventuellement absent.

### Exercice complémentaire — Pointeurs et adresse mémoire

[quota.go](../exercises/02-pointers-memory-address/quota.go) — modification d'un quota via un pointeur, décrémentation d'un compteur transmis par pointeur et observation de la même adresse mémoire avant et après les appels.

## Phase 6 — Gestion des erreurs

### 31 — Type `error`

[score_parser.go](../fundamentals/31-errors/score_parser.go) — conversion et vérification de `err`.

### 32 — `errors.New`

[age_validation.go](../fundamentals/32-errors-new/age_validation.go) — validation d’âge.

### 33 — `fmt.Errorf`

[score_validation.go](../fundamentals/33-fmt-error/score_validation.go) — message dynamique.

### 34 — Enveloppement

[score_loading.go](../fundamentals/34-error-wrapping/score_loading.go) — contexte avec `%w`.

### 35 — `errors.Is`

[score_loading.go](../fundamentals/35-errors-is/score_loading.go) — erreur sentinelle malgré l’enveloppement.

### 36 — `errors.As`

[score_validation_details.go](../fundamentals/36-errors-as/score_validation_details.go) — récupération d’une erreur typée.

### 37 — Propagation

[score_propagation.go](../fundamentals/37-error-propagation/score_propagation.go) — retour de l’erreur à travers une fonction intermédiaire.

### Exercice complémentaire — Rapport

[response_report.go](../exercises/01-error-propagation-report/response_report.go) — erreur lorsqu’une liste de mesures est vide.

### 38 — Erreurs métier

[seat_reservation.go](../fundamentals/38-business-errors/seat_reservation.go) — deux erreurs
métier nommées, distinction avec `errors.Is`, retour du nouvel état et conservation de
l'état initial en cas de refus.

## Phase 7 — Tests

### 39 — Premier test unitaire

[seat_availability.go](../fundamentals/39-first-tests/seat_availability.go) et [seat_availability_test.go](../fundamentals/39-first-tests/seat_availability_test.go) — vérification automatisée d'une règle de disponibilité avec le package `testing`.
