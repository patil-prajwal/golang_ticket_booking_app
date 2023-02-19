package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var remainingTickets uint = 50
var conferenceName = "Go Conference" //conferenceName := "Go Conference" //Since package level vars cant be declared using := (In main func)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// WaitGroup - Concept of threading - craeted..
var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicket { // all 3 needs to be TRUE

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1) // Adds the thread

		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("These people done bookings previously: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Tickets are sold out...!!!")
			//break // End Execution
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
	wg.Wait() // wait till thread to complete execution
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total %v tickets and %v tickets are available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your  tickets from here...")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
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

	// Create a map (Dictionary) to store info. of users.
	// map can hold same datatype of both key and pair - eg. int:int (Simplyy we cant mix datatypes)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank You %v %v for booking %v tickets. You will receive confirmation email at %v shortly !!!\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v Tickets are Remaining for %v.\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v Tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("***************")
	fmt.Printf("Sending Ticket :\n %v to email id %v\n", ticket, email)
	fmt.Println("***************")
	wg.Done() // removes the thread
}
