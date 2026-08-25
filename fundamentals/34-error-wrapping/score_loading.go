package main

import (
	"errors"
	"fmt"
)

func locateScore() error {
	return errors.New("score introuvable")
}

func loadScore() error {
	err := locateScore()
	if err != nil {
		return fmt.Errorf("chargement du score : %w", err)
	}
	return nil
}

func main() {
	err := loadScore()
	if err != nil {
		fmt.Println(err)
	}
}
