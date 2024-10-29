package main

import (
	"booking-app/helper"
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

const conferenceTicket int = 50

var conferenceName = "GopherCon"
var bookings = make([]UserData, 0)

type UserData struct {
	firstname string
	lastname  string
	email     string
	tickets   uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstname, lastname, email, numberOftickets := getUserInput()
	isValidName, isValidEmail, isValidTickets := helper.IsValid(firstname, lastname, email, numberOftickets)

	if isValidName && isValidEmail && isValidTickets {
		bookTicket(firstname, lastname, numberOftickets, email)

		wg.Add(1)
		go sendTicket((uint)(numberOftickets), firstname, lastname, email)

		firstnames := getFirstName()
		fmt.Printf("List of attendees: %v\n", firstnames)

		if helper.RemainingTicket == 0 {
			fmt.Println("Sorry, all tickets are sold out")
		}
	} else {
		if !isValidName {
			fmt.Println("Invalid name name must more than 0 characters")
		} else if !isValidEmail {
			fmt.Println("Invalid email format must contain @ and .")
		} else if !isValidTickets {
			fmt.Printf("Sorry, you can only book maximum %v tickets\n", helper.RemainingTicket)
		}
		fmt.Println("Invalid input")

	}

	wg.Wait()
	fmt.Println("All tickets have been sent")
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Total ticket of %v ticket and %v are still available\n", conferenceTicket, helper.RemainingTicket)
	fmt.Print("Get your ticket now!\n")
}

func getFirstName() []string {
	firstnames := []string{}
	for _, booking := range bookings {
		firstnames = append(firstnames, booking.firstname)
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
	lastname = readMultipleWords()
	fmt.Print("Enter your email: ")
	fmt.Scanln(&email)
	fmt.Print("Enter number of tickets: ")
	fmt.Scanln(&numberOftickets)

	return firstname, lastname, email, numberOftickets
}

func bookTicket(firstname string, lastname string, numberOftickets int, email string) {
	user := firstname + " " + lastname
	helper.RemainingTicket -= numberOftickets

	var userData = UserData{
		firstname: firstname,
		lastname:  lastname,
		email:     email,
		tickets:   (uint)(numberOftickets),
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thank you %v, you have successfully booked %v tickets\n", user, numberOftickets)
	fmt.Printf("You will notifed by your email %v\n", email)
	fmt.Printf("Total remaining tickets: %v\n\n", helper.RemainingTicket)
}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(time.Second * 10)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTicket, firstName, lastName)
	fmt.Printf("\n\n##################\n")
	fmt.Printf("Sending ticket:\n%v to email address %v\n", ticket, email)
	fmt.Println("##################")
	fmt.Println()
	wg.Done()
}

func readMultipleWords() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
