package main

//fmt : format package
import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint64 = 50
	var remainingTickets uint64 = 50
	fmt.Println()
	fmt.Printf("Welcome to Booking App for %v\n", conferenceName)
	fmt.Println();
	fmt.Printf("We have total of %v tickets and only %v are left! \n", conferenceTickets, remainingTickets)
	fmt.Println();
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint64
	//Asking user for information

	fmt.Print("What is your first and last name: ")
	fmt.Scan(&firstName, &lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("User %s %s booked %d tickets", firstName, lastName, userTickets)

	fmt.Println()
	fmt.Println()
}
