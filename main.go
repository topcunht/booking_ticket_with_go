package main

import (
	"fmt"
	"sync"
	"time"
)

const totalTickets int = 50

var conferanceName = "Go Conference"
var remainingTickets uint = 50
var participantList = make([]UserData, 0)

type UserData struct {
	firstName        string
	lastName         string
	email            string
	userBookedTicket uint
}

func main() {
	// greeting to user
	greetUsers()

	firstName, lastName, email, userBookedTicket := getUserInput()
	isValidName, isValidEmail, isValidTicketCount := ValidateUserInput(firstName, lastName, email, userBookedTicket, remainingTickets)

	if isValidEmail && isValidTicketCount && isValidName {
		// booking
		bookingTicket(firstName, lastName, email, userBookedTicket)
		wg.Add(1)
		go sendTickets(userBookedTicket, firstName, lastName, email)
		// participants printing
		participants := printParticipants()
		fmt.Printf("Participant List : %v \n", participants)

		var checkingTicket bool = remainingTickets == 0
		if checkingTicket {
			// no ticket
			fmt.Println("Our tickets are sold out :(")
		}
	} else {
		// invalid data input
		if !isValidName {
			fmt.Printf("Your Name or Surname is too short. You entered %v %v . Please correct it . \n ", firstName, lastName)
		}
		if !isValidEmail {
			fmt.Printf("Your email does not contain @. You entered %v . Please correct it. \n", email)
		}
		if !isValidTicketCount {
			fmt.Printf("Number of tickets you booked is invalid. You entered %v . Please correct. \n", userBookedTicket)
		}

	}
	wg.Wait()
}

var wg = sync.WaitGroup{}

func greetUsers() {
	fmt.Printf("Welcome to %v about booking app \n", conferanceName)
	fmt.Printf("We have %v and %v left \n ", totalTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")
}

func printParticipants() []string {
	firstNames := []string{}
	for _, participants := range participantList {
		firstNames = append(firstNames, participants.firstName)
	}
	return firstNames

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userBookedTicket uint
	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email adress:")
	fmt.Scan(&email)

	fmt.Println("Please enter how many tickets will you book:")
	fmt.Scan(&userBookedTicket)

	return firstName, lastName, email, userBookedTicket
}

func bookingTicket(firstName string, lastName string, email string, userBookedTicket uint) {
	remainingTickets = remainingTickets - userBookedTicket

	// mapping
	var userData = UserData{
		firstName:        firstName,
		lastName:         lastName,
		email:            email,
		userBookedTicket: userBookedTicket,
	}

	participantList = append(participantList, userData)
	fmt.Printf("List of participats' information: %v \n", participantList)

	fmt.Printf("%v %v booked %v tickets for the %v . You will recieve an email via %v \n ", firstName, lastName, userBookedTicket, conferanceName, email)
	fmt.Printf("%v tickets left for the %v \n", remainingTickets, conferanceName)
}

func sendTickets(userBookedTicket uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticketMessage = fmt.Sprintf("%v tickets booked by %v %v for the %v", userBookedTicket, firstName, lastName, conferanceName)
	fmt.Println("################")
	fmt.Printf("Sending ticket: \n %v \n to %v \n", ticketMessage, email)
	fmt.Println("################")
	wg.Done()
}
