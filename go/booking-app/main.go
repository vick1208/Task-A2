package main

import (
	"fmt"
	"strings"
)

const conTic int = 50

var nameCon = "Go Con"
var remTic uint = 50
var books = []string{}

func main() {

	introUser()

	// Ask user for ticket ;p

	for {

		firstName, lastName, email, userTic := getUserInput()

		isVName, isVEmail, isVTNum := ValidInput(firstName, lastName, email, userTic, remTic)

		if isVName && isVEmail && isVTNum {

			firstNames := getFirstName()

			bookTicket(userTic, firstName, lastName, email)
			fmt.Printf("These are our bookings: %v\n", firstNames)

			if remTic == 0 {
				fmt.Println("Our con is booked out. Come back next month")
				break
			}
		} else {
			if !isVName {
				fmt.Println("the name is too short")
			} else if !isVEmail {
				fmt.Println("there is no @")
			} else if !isVTNum {
				fmt.Println("invalid number")
			}
		}
	}
}

func introUser() {
	fmt.Printf("Welcome to %v booking app\n", nameCon)
	fmt.Printf("We have %v tickets and %v are still available\n", conTic, remTic)
	fmt.Println("Get ur ticket now")
}

func getFirstName() []string {
	firstNames := []string{""}
	for _, book := range books {
		var names = strings.Fields(book)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTic uint

	fmt.Print("First Name: ")
	fmt.Scan(&firstName)
	fmt.Print("Last Name: ")
	fmt.Scan(&lastName)
	fmt.Print("Email: ")
	fmt.Scan(&email)
	fmt.Print("Numbers of Ticket(s): ")
	fmt.Scan(&userTic)

	return firstName, lastName, email, userTic

}

func bookTicket(userTic uint, firstName string, lastName string, email string) {
	remTic -= userTic
	books = append(books, firstName+" "+lastName)

	fmt.Printf("Thanks %v %v for booking %v ticket(s). You'll receive a verify email at %v\n", firstName, lastName, userTic, email)
	fmt.Printf("%v %v tickets remains\n", remTic, nameCon)
}
