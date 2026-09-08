package firsttests

import "testing"

func TestSeatsAvailability(t *testing.T) {
	got := SeatsAvailability(3, 5)
	want := true

	if got != want {
		t.Errorf("SeatsAvailability(3, 5) = %t, attendu %t", got, want)
	}
}
