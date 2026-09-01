# Erreurs métier

> Statut : **notion en cours**. Cette fiche résume le but de la notion sans fournir
> la solution de l’exercice actuel.

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
demandée, limite, identifiant concerné ou règle violée. L’appelant peut ensuite
présenter un message adapté ou adopter un comportement spécifique.

## Question de contrôle

Quelle règle concrète le programme de réservation doit-il protéger, et quelles
informations l’appelant doit-il connaître lorsque cette règle est refusée ?
