package main

import (
	"fmt"
	"strings"
)

const conferenceTicket int = 50

var conferenceName = "GopherCon"
var remainingTicket int = 10
var bookings = []string{}

func main() {
	// conferenceYear := 2021
	// conferenceLocation := "Online"
	// conferenceDate := "May 27-28"

	// fmt.Printf("Conference Year: %d\n", conferenceYear)
	// fmt.Println("Conference Location: ", conferenceLocation)
	// fmt.Println("Conference Date: ", conferenceDate)

	// const conferenceDay uint = 1
	// fmt.Println("Conference Day: ", conferenceDay)

	greetUsers()

	for {

		firstname, lastname, email, numberOftickets := getUserInput()
		isValidName, isValidEmail, isValidTickets := isValid(firstname, lastname, email, numberOftickets)

		if isValidName && isValidEmail && isValidTickets {
			bookTicket(firstname, lastname, numberOftickets)

			firstnames := getFirstName(bookings)
			fmt.Printf("List of attendees: %v\n", firstnames)

			if remainingTicket == 0 {
				fmt.Println("Sorry, all tickets are sold out")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Invalid name name must more than 0 characters")
				continue
			} else if !isValidEmail {
				fmt.Println("Invalid email format must contain @ and .")
				continue
			} else if !isValidTickets {
				fmt.Printf("Sorry, you can only book maximum %v tickets\n", remainingTicket)
				continue
			}
			fmt.Println("Invalid input")

		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Total ticket of %v ticket and %v are still available\n", conferenceTicket, remainingTicket)
	fmt.Print("Get your ticket now!\n")
}

func getFirstName(bookings []string) []string {
	firstnames := []string{}
	for _, booking := range bookings {
		var name = strings.Fields(booking)
		var firstname = name[0]
		firstnames = append(firstnames, firstname)
	}
	return firstnames
}

func getUserInput() (string, string, string, int) {
	var firstname string
	var lastname string
	var numberOftickets int
	var email string

	fmt.Print("Enter your first name: ")
	fmt.Scanln(&firstname)
	fmt.Print("Enter your last name: ")
	fmt.Scanln(&lastname)
	fmt.Print("Enter your email: ")
	fmt.Scanln(&email)
	fmt.Print("Enter number of tickets: ")
	fmt.Scanln(&numberOftickets)

	return firstname, lastname, email, numberOftickets
}

func isValid(firstname string, lastname string, email string, numberOftickets int) (bool, bool, bool) {
	isValidName := len(firstname) >= 0 && len(lastname) >= 0
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTickets := numberOftickets > 0 && numberOftickets <= remainingTicket

	return isValidName, isValidEmail, isValidTickets
}

func bookTicket(firstname string, lastname string, numberOftickets int) {
	user := firstname + " " + lastname
	bookings = append(bookings, user)
	remainingTicket -= numberOftickets
	fmt.Printf("Congratulations %v, you have successfully booked %v tickets\n", user, numberOftickets)
	fmt.Printf("Total remaining tickets: %v\n", remainingTicket)
}
