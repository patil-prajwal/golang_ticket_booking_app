package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference" //conferenceName := "Go Conference" //Since package level vars cant be declared using := (In main func)
var remainingTickets uint = 50
var bookings []string // bookings := []string{} //Since package level vars cant be declared using := (In main func)

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicket := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicket { // all 3 needs to be TRUE
			bookTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("These people done bookings previously: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Tickets are sold out...!!!")
				break // End Execution
			}

		} else {
			if !isValidName {
				fmt.Println("Either First Name or Last Name is Invalid. Please Try Again...")
			}
			if !isValidEmail {
				fmt.Println("Invalid Email Entered. Please Try Again..")
			}
			if !isValidTicket {
				fmt.Println("Invalid Ticket number. Please Try Again..")
			}
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total %v tickets and %v tickets are available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your  tickets from here...")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicket
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter Your First Name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter Your Last Name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter Your Email ID: ")
	fmt.Scan(&email)
	fmt.Println("Enter Number of tickets you want to buy: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Thank You %v %v for booking %v tickets. You will receive confirmation email at %v shortly !!!\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v Tickets are Remaining for %v.", remainingTickets, conferenceName)
}

