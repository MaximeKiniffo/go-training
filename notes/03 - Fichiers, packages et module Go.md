# Fichiers, packages et module Go

## Plusieurs fichiers, même package

Deux fichiers d’un même dossier peuvent partager le même package :

```go
package main
```

Ils sont compilés ensemble. Les fonctions non exportées sont donc accessibles entre
ces fichiers, puisqu’ils appartiennent au même package.

## Package local

Un sous-dossier peut former un package à importer :

```go
import "go-training/fundamentals/18-local-packages/profile"
```

Le début du chemin d’import vient du nom déclaré dans `go.mod`.

## `go.mod`

`go.mod` identifie le module et les dépendances du projet. Il ne contient pas le code
du programme, mais permet à Go de résoudre les imports de façon cohérente.

## Comparaison JavaScript / TypeScript

Un package Go ressemble à un module, mais tous les fichiers d’un dossier composent
une unité de compilation. L’import vise le dossier/package, pas un export nommé d’un
fichier précis.

## Exercices associés

[[Exercices terminés#17 — Plusieurs fichiers]] · [[Exercices terminés#18 — Packages locaux et module]]
