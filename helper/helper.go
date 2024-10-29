package helper

import "strings"

var RemainingTicket int = 10

func IsValid(firstname string, lastname string, email string, numberOftickets int) (bool, bool, bool) {
	isValidName := len(firstname) >= 0 && len(lastname) >= 0
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTickets := numberOftickets > 0 && numberOftickets <= RemainingTicket

	return isValidName, isValidEmail, isValidTickets
}
