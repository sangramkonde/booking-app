package main

import (
	"fmt"
	"strings"
)

// Package level varaible, accessible to all but within a package
var conferenceName = "Go Conference"
const conferenceTickets = 50  
var remainingTickets uint = 50
var bookings []string

func main(){

	greetUsers()
    
	for remainingTickets > 0 {
	
		firstName, lastName, email, userTickets := getUserInput()
        isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)
       
		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
	
            firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			
			if remainingTickets == 0 {
				// end program
				fmt.Printf("Our %v is booked out. Come back next year.", conferenceName)
				break
			}
		}else{
			if !isValidName{
				fmt.Println("First name or last name you entered is too short. Both should be more than 2 characters.")
			}
			if !isValidEmail{
				fmt.Println("Email you entered doesn't contain '@' sign. Please provide valid email address.")
			}
			if !isValidTicketNumber{
			    fmt.Println("Number of tickes you entered is invalid. Please enter valid input.")
			}
		}
	}
}

func greetUsers(){
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string{

	firstNames := []string{}
	for _, booking := range bookings{
	  var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool){
	isValidName:= len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput()(string, string, string, uint){

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// We use fmt.Scan to get the input from user and store into variable
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string){

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName +" "+ lastName)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}