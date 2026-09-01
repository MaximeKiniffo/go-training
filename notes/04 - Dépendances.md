# Dépendances

## Idée

Un service reçoit la collaboration dont il a besoin au lieu de la créer lui-même.

```go
type TaskService struct {
	Notifier Notifier
}
```

Le service dépend du contrat `Notifier`, non de `ConsoleNotifier` en particulier.

## Pourquoi c’est utile

- le comportement peut changer sans modifier le service ;
- le code a une responsabilité plus claire ;
- il sera plus simple à tester lorsque les tests seront étudiés.

## À ne pas faire

Ajouter une interface et une couche d’indirection si une seule implémentation stable
existe sans besoin de remplacement. L’abstraction doit résoudre un problème concret.

## Exercice associé

[[Exercices terminés#25 — Dépendances]]
