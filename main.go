package main

import "fmt"

func main() {
	// variables
	var conferenceName = "Go Conference"
	// consts are like variable but cant changed!
	const conferenceTickets = 50
	var remainingTickets = 30
	fmt.Println("Welcome to", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are still available!\n", conferenceTickets, remainingTickets)

	var userName string
	var userTicket int
	// ask user for their name
	userName = "Tom"
	userTicket = 2
	fmt.Printf("user %v booked %v tickets.\n", userName, userTicket)

	// to print type of variables
	fmt.Printf("conferenceName type is %T and conferenceTicketes is %T and userTicket is %T\n", conferenceName, conferenceTickets, userTicket)

}