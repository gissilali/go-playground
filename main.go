package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var eventName = "Go Conf"
	const conferenceTickets uint8 = 2
	var remainingTickets = conferenceTickets
	var bookings []string

	for {
		greetUsers(eventName, conferenceTickets, remainingTickets)
		scanner := bufio.NewScanner(os.Stdin)
		var userTickets uint8
		fmt.Println("What's yo name?")
		scanner.Scan()
		userName := scanner.Text()
		fmt.Println("How many tickets you want?")
		fmt.Scan(&userTickets)
		if userTickets > remainingTickets {
			fmt.Printf("Ok bruh you can't buy %v tickets when only %v tickets remain\n", userTickets, remainingTickets)
			continue
		}
		fmt.Printf("Ok %v you bought %v tikcets\n", userName, userTickets)
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, userName)
		for _, booking := range bookings {
			fmt.Println(booking, "Has booked a fucking ticket")
		}
		fmt.Print(remainingTickets, " tickets remaining\n")
		if remainingTickets == 0 {
			break
		}
	}
}

func greetUsers(eventName string, conferenceTickets uint8, remainingTickets uint8) {
	fmt.Printf("Welcome to %v Conference\n", eventName)
	fmt.Printf("We have a total of %v and %v remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get yo tickets here")
}
