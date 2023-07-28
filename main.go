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

	// we can also assign type to variables when we define it.
	// in some cases we need for example only positiv number for integer
	var userAge uint = 20
	fmt.Printf("%v is %v years old!\n", userName, userAge)
	// got error with this >>> userAge = -1
	// in the calculation due to the negative result, a random number will be printed. >>> userAge = userAge - 23
	// fmt.Printf("userage is %v\n", userAge)

}
