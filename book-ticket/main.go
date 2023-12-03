package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceName string = "Go Conference"
const conferenceTickets uint64 = 50

var remainingTickets uint64 = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOftickets uint64
}

var wg = sync.WaitGroup{}

func main() {
	greetUser()

	firstName, lastName, email, userTickets := getUserinput()
	isValidEmail, isValidName, isValidTicket := ValidateUserinput(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicket {
		remainingTickets, bookings = bookTicket(userTickets, bookings, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		firstNames := getFirstNames(bookings)
		fmt.Printf("The names of Booking are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out!")
		}
	} else {
		if !isValidName {
			fmt.Printf("Your input name is invalid\n")
		}
		if !isValidEmail {
			fmt.Printf("Your input email is invalid and doesn't contain @ sign\n")
		}
		if !isValidTicket {
			fmt.Printf("Number of tickets you choose is invalid\n")
		}
	}

	fmt.Println()
	wg.Wait()
}

func greetUser() {
	fmt.Println()
	fmt.Printf("Welcome to Booking App for %v\n", conferenceName)
	fmt.Println()
	fmt.Printf("We have a total of %v tickets, and only %v are left!\n", conferenceTickets, remainingTickets)
	fmt.Println()
	fmt.Println("Get your tickets here to attend")
	fmt.Println()
}

func getFirstNames(bookings []userData) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func validationPrint(firstName string, lastName string, userTickets uint64, email string, remainingTickets uint64) {
	fmt.Println()
	fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email @%s\n", firstName, lastName, userTickets, email)
	fmt.Printf("Only %d tickets left\n", remainingTickets)
}

func getUserinput() (string, string, string, uint64) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint64
	fmt.Print("What is your first and last name: ")
	fmt.Scan(&firstName, &lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter the number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint64, bookings []userData, firstName string, lastName string, email string) (uint64, []userData) {
	remainingTickets -= userTickets

	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOftickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("\nList of bookings are %v\n", bookings)
	validationPrint(firstName, lastName, userTickets, email, remainingTickets)

	return remainingTickets, bookings
}

func sendTicket(userTickets uint64, firstName string, lastName string, email string) {

	time.Sleep(15 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println()
	fmt.Printf("Sending ticket: \n %v \nto email address %v\n", ticket, email)
	fmt.Println()
	wg.Done()
}
