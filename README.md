# Go Training

Dépôt personnel consacré à mon apprentissage de Go avant mon alternance.

## Objectif

Apprendre progressivement :

- les fondamentaux de Go ;
- le développement backend ;
- les tests ;
- les API HTTP ;
- PostgreSQL ;
- Docker ;
- les bases d’une architecture microservices.

## Organisation du dépôt

```text
go-training/
├── fundamentals/  # Notions Go étudiées, conservées par sujet
├── exercises/     # Exercices complémentaires et évaluations
├── mini-projects/ # Projets progressifs
├── services/      # Services backend
└── notes/         # Notes de révision
```

Les exercices fondamentaux sont exécutables séparément. Leur fichier porte un nom descriptif, par exemple `variables.go`, même s’il contient une fonction `main`.

## Exécuter un exercice

Depuis la racine du dépôt :

```powershell
go run ./fundamentals/01-first-program
go run ./fundamentals/02-variables
```

## Vérifier le projet

```powershell
go fmt ./...
go vet ./...
go test ./...
