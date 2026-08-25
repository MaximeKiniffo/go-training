package main

import (
	"fmt"
	"strconv"
)

func parseScore(text string) (int, error) {
	return strconv.Atoi(text)
}

func main() {
	score, err := parseScore("100")
	score2, err2 := parseScore("abc")
	if err != nil {
		fmt.Println("Impossible de convertir le score : ", err)
	} else {
		fmt.Println("Score valide :", score)
	}
	if err2 != nil {
		fmt.Println("Impossible de convertir le score : ", err2)
	} else {
		fmt.Println("Score valide :", score2)
	}
}
