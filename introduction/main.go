package main

import (
	"fmt"
	"strings"
)

func main() {
	// var conferenceName string = "Go Conference"
	conferenceName := "Go Conference" // Same as previous line // Syntactical Sugar

	const conferenceTickets int = 50
	var remainingTickets uint = 50 //
	var bookings []string          // Slice

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T.\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Println("Welcome to", conferenceName, "our ticket booking application.")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name

	for {
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email:")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining.\n\n", remainingTickets)
			continue
		}

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets, You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("These are all our bookings: %v\n\n", firstNames)

		noTicketsRemaining := remainingTickets == 0

		if noTicketsRemaining {
			fmt.Printf("Our conference is booked out. Comeback next year.")
			break
		}
	}
}
