package main

import "fmt"

func main(){
	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",conferenceTickets,remainingTickets,conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)	
	fmt.Printf("we have total of %v tickets and %v are stil available\n", conferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here")

	



	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// asking user for their name
	fmt.Print("enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("enter your email: ")
	
	
	fmt.Scan(&email)

	fmt.Print("enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)


	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email %v\n", firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	fmt.Printf("These are all the bookings: %v\n", bookings)	

}
