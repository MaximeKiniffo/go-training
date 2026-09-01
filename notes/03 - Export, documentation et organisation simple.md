# Export, documentation et organisation simple

## Visibilité par la majuscule

- `introduction` est privée au package.
- `Introduction` est exportée et utilisable depuis un autre package.

Cette règle s’applique aux fonctions, types, variables, constantes et champs de
struct.

## Documentation GoDoc

Une fonction exportée reçoit un commentaire immédiatement au-dessus d’elle, qui
commence par son nom :

```go
// Introduction construit une phrase de présentation.
func Introduction(name string) string { }
```

## Organisation simple

- `package main` coordonne l’exécution.
- un package local regroupe une responsabilité précise, par exemple le profil.
- éviter de créer des packages vides ou abstraits « au cas où ».

## À retenir

On découpe le code parce qu’une responsabilité est distincte et devient difficile à
lire, pas simplement parce qu’un fichier dépasse un nombre arbitraire de lignes.

## Exercice associé

[[Exercices terminés#18 — Packages locaux et module]]
