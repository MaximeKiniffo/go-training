# Erreurs métier

> Statut : **notion acquise** après l'exercice de réservation de places.

## Définition

Une erreur métier exprime qu’une règle du domaine n’est pas respectée. Ce n’est pas
un dysfonctionnement technique.

Exemples :

- réserver plus de places que la limite autorisée ;
- créer un compte avec un âge non éligible ;
- modifier une ressource qui n’existe pas.

## Erreur technique ou métier ?

| Situation | Catégorie |
| --- | --- |
| Conversion d’un texte invalide en nombre | technique / donnée invalide |
| Fichier introuvable | technique |
| Limite de places dépassée | métier |
| Solde insuffisant pour une opération autorisée | métier |

## Pourquoi une erreur explicite ?

Une erreur métier peut transporter les informations utiles à la décision : valeur
demandée, limite, identifiant concerné ou règle violée. L'appelant peut ensuite
présenter un message adapté ou adopter un comportement spécifique.

## Mise en œuvre retenue

- Une erreur sentinelle nommée représente chaque règle métier refusée.
- `errors.Is` permet de reconnaître la règle concernée sans comparer le texte du message.
- La fonction métier retourne le nouvel état avec l'erreur.
- En cas d'échec, elle retourne l'état initial afin de ne pas présenter une modification
  qui n'a pas eu lieu.
- Le code appelant transforme l'erreur interne en message destiné à l'utilisateur.

## Pièges

- Ne pas confondre une demande invalide avec une disponibilité insuffisante.
- Vérifier la quantité demandée pour valider la demande ; la disponibilité sert à
  vérifier si la réservation peut ensuite être effectuée.
- Ne pas utiliser `err.Error()` comme identifiant d'une situation métier.

## Exercice associé

[seat_reservation.go](../fundamentals/38-business-errors/seat_reservation.go) — réservation
acceptée, demande invalide et places insuffisantes.

## Questions de révision

1. Quelle différence fais-tu entre une demande invalide et une disponibilité insuffisante ?
2. Pourquoi retourner le nombre initial de places avec une erreur ?
