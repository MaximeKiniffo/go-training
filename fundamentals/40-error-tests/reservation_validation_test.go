package errortests

import (
	"errors"
	"testing"
)

func TestValidateReservation1(t *testing.T) {
	err := ValidateReservation(2, 5)
	if err != nil {
		t.Fatalf("Attendu: pas d'erreur, Obtenu: %v", err)
	}
}

func TestValidateReservation2(t *testing.T) {
	err := ValidateReservation(0, 5)
	if !errors.Is(err, ErrInvalidRequest) {
		t.Errorf("ValidateReservation(\"\") = %v, erreur attendue : %v", err, ErrInvalidRequest)
	}
}

func TestValidateReservation3(t *testing.T) {
	err := ValidateReservation(6, 5)
	if !errors.Is(err, ErrNotEnoughSeats) {
		t.Errorf("ValidateReservation(\"\") = %v, erreur attendue : %v", err, ErrNotEnoughSeats)
	}
}
