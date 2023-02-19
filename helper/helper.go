package helper

import "strings"

// We need to export functions in packages explicitely. To export, just capitalize the first letter of function.
// Same applies to export variables too.
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicket
}
