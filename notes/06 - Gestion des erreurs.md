# Phase 6 — Gestion des erreurs

En Go, une erreur est une valeur qui accompagne le résultat normal d’une fonction.
L’appelant vérifie cette valeur et décide quoi faire.

- [[06 - Type error|Type error]]
- [[06 - errors.New|errors.New]]
- [[06 - fmt.Errorf|fmt.Errorf]]
- [[06 - Envelopper une erreur avec %w|Envelopper une erreur avec %w]]
- [[06 - errors.Is|errors.Is]]
- [[06 - errors.As|errors.As]]
- [[06 - Propagation des erreurs|Propagation des erreurs]]
- [[06 - Erreurs métier|Erreurs métier]]

## Réflexe général

```go
value, err := operation()
if err != nil {
	return err
}
```

Traiter ou propager une erreur juste après l’appel qui peut la retourner rend le code
plus sûr et plus lisible.
