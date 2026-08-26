package main

import (
	"errors"
	"fmt"
)

type ScoreRangeError struct {
	Score         int
	ScoreMinValue int
	ScoreMaxValue int
}

func (e *ScoreRangeError) Error() string {
	return fmt.Sprintf("Le score %d n'est pas valide, il doit être compris entre %d et %d", e.Score, e.ScoreMinValue, e.ScoreMaxValue)
}

func checkScore(score int) error {
	err := &ScoreRangeError{
		Score:         score,
		ScoreMinValue: 0,
		ScoreMaxValue: 100,
	}
	if err.Score > err.ScoreMaxValue {
		return fmt.Errorf("Le score %d est trop haut : %w", err.Score, err)
	} else if err.Score < err.ScoreMinValue {
		return fmt.Errorf("Le score %d est trop bas : %w", err.Score, err)
	}
	return nil
}

func handleScore(score int) {
	err := checkScore(score)

	var scoreErr *ScoreRangeError

	switch {
	case err == nil:
		fmt.Printf("Score %d accepté\n", score)
	case errors.As(err, &scoreErr):
		if scoreErr.ScoreMinValue > scoreErr.Score {
			point := scoreErr.ScoreMinValue - scoreErr.Score
			fmt.Printf("Le score %d est trop bas, il manque %d\n", score, point)
		} else {
			point := scoreErr.Score - scoreErr.ScoreMaxValue
			fmt.Printf("Le score %d est trop haut, il faut retirer %d\n", score, point)
		}
	default:
		fmt.Println("Erreur inattendue :", err)
	}
}

func main() {
	handleScore(85)
	handleScore(-4)
	handleScore(106)
}
