# Erreurs et pièges

Cette note sert à conserver les erreurs réellement rencontrées ou les confusions qui
risquent de revenir. Une erreur comprise devient un point de révision, pas un échec.

## Propagation d’erreurs

- Lire l’énoncé en identifiant séparément l’entrée, la transformation et la sortie.
- Ne pas inverser une chaîne à convertir et le nombre qui résulte de la conversion.
- Suivre `err` juste après chaque appel pouvant échouer.
- Ne jamais poursuivre le calcul normal avec une donnée lorsque `err != nil`.

## Pointeurs

- `&value` prend l'adresse ; `*pointer` lit ou modifie la valeur visée.
- Une fonction reçoit toujours une valeur. Si cette valeur est un pointeur, elle peut
  viser la donnée de l'appelant.
- Ne pas confondre le pointeur avec la valeur pointée : un pointeur conserve une
  adresse, tandis que `*pointer` permet d'accéder à la valeur située à cette adresse.
- Un entier passé par valeur est copié ; il faut retourner la nouvelle valeur ou
  transmettre un pointeur si la fonction doit modifier l'entier de l'appelant.
- Vérifier un pointeur optionnel avant de le déréférencer.

## Slices

- `append` retourne la slice à conserver.
- Une affectation de slice ne garantit pas une copie indépendante.

## À enrichir

Ajouter ici uniquement les erreurs qui se répètent ou qui ont demandé une explication.
Format conseillé : contexte, mauvaise hypothèse, bonne règle, exemple personnel.
