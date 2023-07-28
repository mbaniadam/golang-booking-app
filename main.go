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

	var firstName string
	var lastName string
	var email string
	var userTicket int
	// ask user for their name
	fmt.Println("Enter your first name: ")
	// & is a pointer for memory address of variable
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTicket)

	fmt.Printf("user %v %v booked %v tickets.\n", firstName, lastName, userTicket)

	// to print type of variables
	fmt.Printf("conferenceName type is %T and conferenceTicketes is %T and userTicket is %T\n", conferenceName, conferenceTickets, userTicket)

	// we can also assign type to variables when we define it.
	// in some cases we need for example only positiv number for integer
	var userAge uint = 20
	fmt.Printf("%v %v is %v years old!\n", firstName, lastName, userAge)
	// got error with this >>> userAge = -1
	// in the calculation due to the negative result, a random number will be printed. >>> userAge = userAge - 23
	// fmt.Printf("userage is %v\n", userAge)

	// another way to declare a variable
	userBirthCity := "Tehran"
	fmt.Println(userBirthCity)

}
