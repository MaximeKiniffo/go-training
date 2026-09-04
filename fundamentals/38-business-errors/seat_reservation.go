package main

import (
	"errors"
	"fmt"
)

var ErrInvalidReservation = errors.New("le nombre de places doit être positif.")
var ErrInsufficientSeats = errors.New("il ne reste pas assez de places.")

func reserveSeats(availableSeats int, requestedSeats int) (int, error) {
	if requestedSeats <= 0 {
		return availableSeats, ErrInvalidReservation
	} else if requestedSeats > availableSeats {
		return availableSeats, ErrInsufficientSeats
	}
	return availableSeats - requestedSeats, nil
}

func handleReservation(availableSeats int, requestedSeats int) {
	remainingSeats, err := reserveSeats(availableSeats, requestedSeats)
	switch {
	case err == nil:
		fmt.Println("Réservation acceptée. Places restantes : ", remainingSeats)
	case errors.Is(err, ErrInsufficientSeats):
		fmt.Println("Réservation refusée : il ne reste pas assez de places.")
	case errors.Is(err, ErrInvalidReservation):
		fmt.Println("Réservation refusée : le nombre de places doit être positif.")
	default:
		fmt.Println("Erreur inattendue :", err)
	}
}

func main() {
	handleReservation(5, 2)
	handleReservation(5, 0)
	handleReservation(2, 4)
}
