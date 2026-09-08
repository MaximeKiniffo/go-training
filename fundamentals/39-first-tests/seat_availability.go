package firsttests

func SeatsAvailability(request int,
	availability int) bool {
	if request > availability {
		return false
	}
	return true
}
