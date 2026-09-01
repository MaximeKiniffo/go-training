# Glossaire Go

| Terme | Définition courte |
| --- | --- |
| Package | Ensemble de fichiers Go d’un même dossier et d’une même responsabilité. |
| Module | Unité de projet définie par `go.mod`, utilisée pour résoudre les imports. |
| Exporté | Identifiant commençant par une majuscule, accessible depuis un autre package. |
| Struct | Type regroupant plusieurs champs nommés. |
| Slice | Séquence flexible de valeurs, construite sur un tableau sous-jacent. |
| Map | Collection associant une clé à une valeur. |
| Receiver | Valeur placée avant le nom d’une méthode et sur laquelle la méthode opère. |
| Pointeur | Valeur qui contient l’adresse d’une autre valeur. |
| `nil` | Absence de valeur pour certains types, notamment les pointeurs. |
| Interface | Ensemble de méthodes attendues par un comportement. |
| Erreur sentinelle | Valeur d’erreur partagée, utile pour reconnaître un cas précis avec `errors.Is`. |
| Enveloppement | Ajout de contexte à une erreur avec `%w` sans perdre l’erreur d’origine. |
| Erreur métier | Erreur représentant une règle du domaine non respectée. |
