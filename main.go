package main

import (
	"fmt"
	"sync"
	"time"

	// "strconv"
	"strings"
	// "strings"
)

// Package level varaible, accessible to all but within a package
var conferenceName = "Go Conference"
const conferenceTickets = 50  
var remainingTickets uint = 50
// var bookings []string

// empty list of map
// var bookings = make([]map[string]string, 0)

// empty list of struct 'UserData'
var bookings = make([]UserData, 0)


//  struct type:
// Collection of different data type of data
// then we can create struct type, a custom data type
type UserData struct{
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

// WaitGroup - Waits for the launched goroutine to finish
// Pakackage "sync" provides basic synchronization functionality
// Add: Sets the number of goroutine to wait for(increase the counter by the provided number)
// Wait: Blocks until the WaitGroup conuter is 0
// Done: it removed thread from the waiting list.
// It decrease the WaitGroup counter by 1
// This is called by goroutine to indicate that it is finished.
var wg = sync.WaitGroup{}

func main(){

	greetUsers()
    
	// for remainingTickets > 0 {
	
		firstName, lastName, email, userTickets := getUserInput()
        isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)
       
		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			// Add number of threads/go routine to wait
			wg.Add(1)

			// Goroutine:
			// Go is using, what's called a 'Green thread'
			// Abstraction of an actual thread
			// Menaged by the go runtime, we are only interacting with these high level goroutines
			// Cheaper and lightweight
			// You can run hundread of thousands or millions goroutines
			// without affecting the performance
			// Goroutine has 'Channel' to commuinate between goroutines
			go sendTicket(userTickets, firstName, lastName, email)
	
            firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			
			if remainingTickets == 0 {
				// end program
				fmt.Printf("Our %v is booked out. Come back next year.", conferenceName)
				// break
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
		wg.Wait()
	// }
}

func greetUsers(){
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string{

	firstNames := []string{}
	for _, booking := range bookings{
		firstNames = append(firstNames, booking.firstName)
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

	// create a map for a user - map is a key-value pair data structure
	// var mySlice []string   
	// var myMap map[string]string
	// var userData = make(map[string]string)

	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)


	// bookings = append(bookings, firstName +" "+ lastName)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("########################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("########################")

	// Done: it removed thread from the waiting list.
	// It decrease the WaitGroup counter by 1
	// This is called by goroutine to indicate that it is finished.
	wg.Done()
}