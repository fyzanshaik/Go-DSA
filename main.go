package main

//fmt : format package
import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets uint64 = 50
	var remainingTickets uint64 = 50

	bookings := []string{}

	greetUser(conferenceName, conferenceTickets, remainingTickets)

	for {
		firstName,lastName,email,userTickets := getUserinput();
		isValidemail, isValidname, isValidticket := validateUserinput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidname && isValidemail && isValidticket {
			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			validationPrint(firstName, lastName, userTickets, email, remainingTickets)
			fmt.Printf("The first names of booking are: %v", printFirstNames(bookings))

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out!")
				break
			}
		} else {
			if !isValidname {
				fmt.Printf("Your input name is invalid\n")
			}
			if !isValidemail {
				fmt.Printf("Your input email is invalid doesn't contain @sign\n")
			}
			if !isValidticket {
				fmt.Printf("Number of tickets you choose are invalid\n")
			}
		}

		fmt.Println()
		fmt.Println()
	}
}

func greetUser(name string, tickets uint64, remainingTickets uint64) {
	fmt.Println()
	fmt.Printf("Welcome to Booking App for %v\n", name)
	fmt.Println()
	fmt.Printf("We have total of %v tickets and only %v are left! \n", tickets, remainingTickets)
	fmt.Println()
	fmt.Println("Get your tickets here to attend")
	fmt.Println()
}

func printFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, element := range bookings {
		var names = strings.Fields(element)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validationPrint(firstName string, lastName string, userTickets uint64, email string, remainingTickets uint64) {
	fmt.Println()
	fmt.Printf("Thank you %s %s for booking %d tickets. You will recieve a confirmation email @%s\n", firstName, lastName, userTickets, email)
	fmt.Printf("Only %d tickets left\n", remainingTickets)
}

func validateUserinput(firstName string, lastName string, email string, userTickets uint64, remainingTickets uint64) (bool, bool, bool) {
	isValidname := len(firstName) >= 2 && len(lastName) >= 2
	isValidemail := strings.Contains(email, "@")
	isValidticket := userTickets > 0 && userTickets <= remainingTickets

	return isValidemail, isValidname, isValidticket
}

func getUserinput() (string,string,string,uint64){
	var firstName string
	var lastName string
	var email string
	var userTickets uint64
	fmt.Print("What is your first and last name: ")
	fmt.Scan(&firstName, &lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName,lastName,email,userTickets;
}

func bookTicket() {

	
}
