package main

import (
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, userBookedTicket uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketCount := userBookedTicket > 0 && userBookedTicket <= remainingTickets
	return isValidName, isValidEmail, isValidTicketCount
}
