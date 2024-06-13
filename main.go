package main

import "fmt"

func main(){
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)	
	fmt.Printf("we have total of %v tickets and %v are stil available\n", conferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here")

	var userName string
	var userTickets int
	// asking user for their name

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName,userTickets)

	

}
