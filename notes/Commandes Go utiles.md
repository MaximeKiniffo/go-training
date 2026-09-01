# Commandes Go utiles

Exécuter ces commandes dans PowerShell, depuis le dossier adapté au programme ou à la
racine du module selon la commande.

## Vérifier l’installation

```powershell
go version
```

Affiche la version de Go disponible dans le terminal.

## Lancer un programme

```powershell
go run .
```

À lancer dans le dossier contenant le package `main` à exécuter. Go compile
temporairement et lance le programme.

## Compiler un programme

```powershell
go build .
```

Compile le package courant. Sous Windows, cela peut produire un exécutable `.exe`
dans le dossier courant.

## Formater le code

```powershell
go fmt ./...
```

À lancer à la racine du module. Formate tous les packages du module selon les règles
Go. Une absence de sortie est normale.

## Créer un module

```powershell
go mod init go-training
```

Crée `go.mod` dans le dossier courant. À faire une seule fois lors de l’initialisation
d’un nouveau module.

## Erreurs courantes

- `go` n’est pas reconnu : ouvrir un nouveau PowerShell après l’installation ou
  vérifier le `PATH`.
- `no Go files` : être dans le bon dossier ou cibler le bon package.
- import inutilisé : retirer l’import ou utiliser réellement le package.
