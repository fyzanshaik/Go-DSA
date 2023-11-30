package main

//fmt : format package
import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Println("Welcome to Booking App for", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "left!")
	fmt.Println("Get your tickets here to attend")

}
