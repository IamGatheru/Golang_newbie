package main

import (
    "strings"
    "fmt"
)

func main() {

    conferenceName := "Go Conference"
    const conferenceTickets uint= 50
    var remainingTickets uint = 50
    var bookings [] string

    fmt.Printf("ConferenceTickets is %T, remainingTickets is %T, conferenceName is %T.", conferenceTickets, remainingTickets, conferenceName)

    fmt.Printf("Welcome to %v booking application\n", conferenceName)
    fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
    fmt.Println("Get your tickets here to attend")

    for {
	    var firstName string
	    var lastName string
	    var userTickets uint
	    var email string

	    // Ask user for their first name
	    fmt.Println("Please enter your first name:")
	    fmt.Scan(&firstName)

	    fmt.Println("Enter your last name:")
	    fmt.Scan(&lastName)

	    fmt.Println("Enter your email address:")
	    fmt.Scan(&email)

	    fmt.Println("Enter number of tickets:")
	    fmt.Scan(&userTickets)

	    if userTickets <=  remainingTickets {
	    
	    remainingTickets = remainingTickets - userTickets
	    bookings = append(bookings, firstName + " " + lastName)

	    fmt.Printf("Thank you %v %v for booking %v ticket(s) with us.\n", firstName, lastName, userTickets)
	    fmt.Printf("You will receive a confirmation email at %v.\n", email)

	    firstNames  := []string{}
	    for _, booking := range bookings {
		    var names = strings.Fields(booking)
		    var firstName = names[0]
		    firstNames = append(firstNames, firstName)
	    }
	    fmt.Printf("The first names of bookings are : %v\n", firstNames)

	    noTicketsRemaining := remainingTickets == 0
	    if noTicketsRemaining {
		    //end program
		    fmt.Println("Our conference is booked out. Come back next year.")
		    break
	    }
    } else {
	    fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
	    continue
	}
}
}
