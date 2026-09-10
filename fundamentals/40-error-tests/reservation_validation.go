package errortests

import "errors"

var ErrInvalidRequest = errors.New("Le nombre de place demandé doit etre superieur à zero")
var ErrNotEnoughSeats = errors.New("Il ne reste pas assez de place")

func ValidateReservation(request int, available int) error {
	if request <= 0 {
		return ErrInvalidRequest
	}

	if request > available {
		return ErrNotEnoughSeats
	}

	return nil
}
