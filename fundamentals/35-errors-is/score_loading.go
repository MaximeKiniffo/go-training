package main

import (
	"errors"
	"fmt"
)

var (
	ErrScoreNotFound      = errors.New("score introuvable")
	ErrStorageUnavailable = errors.New("stockage indisponible")
)

func loadScore(scoreExists bool, storageAvailable bool) error {
	if !storageAvailable {
		return fmt.Errorf("chargement du score impossible : %w", ErrStorageUnavailable)
	}

	if !scoreExists {
		return fmt.Errorf("chargement du score impossible : %w", ErrScoreNotFound)
	}

	return nil
}

func handleScoreResult(err error) {
	switch {
	case err == nil:
		fmt.Println("Score chargé avec succès")
	case errors.Is(err, ErrScoreNotFound):
		fmt.Println("Aucun score sauvegardé : un nouveau score peut être créé")
	case errors.Is(err, ErrStorageUnavailable):
		fmt.Println("Stockage indisponible : réessaie plus tard")
	default:
		fmt.Println("Erreur inattendue :", err)
	}
}

func main() {
	handleScoreResult(loadScore(true, true))
	handleScoreResult(loadScore(false, true))
	handleScoreResult(loadScore(false, false))
}
