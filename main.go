package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	//  ask user for their name
	println("Enter your first name:")
	fmt.Scan(&firstName)

	println("Enter your last name:")
	fmt.Scan(&lastName)

	println("Enter your email Address:")
	fmt.Scan(&email)

	println("How many tickets do you want to buy?:")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for buying %v Tickets, the receipt will be sent to %v.", firstName, lastName, userTickets, email)
}
