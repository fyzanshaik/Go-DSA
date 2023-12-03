package main

import "strings"

func ValidateUserinput(firstName string, lastName string, email string, userTickets uint64) (bool, bool, bool) {
	isValidname := len(firstName) >= 2 && len(lastName) >= 2
	isValidemail := strings.Contains(email, "@")
	isValidticket := userTickets > 0 && userTickets <= remainingTickets

	return isValidemail, isValidname, isValidticket
}
