package main

import (
	"fmt"
	"time"
)

const (
	openingHour     = 10 // 10 AM
	closingHour     = 22 // 10 PM
	maxReservations = 4
)

type Reservation struct {
	CustomerName string
	Date         string
	Time         string
}

var reservations = make(map[string][]Reservation)

func addReservation(name, date, timeStr string) {
	if !isTimeSlotAvailable(date, timeStr) {
		fmt.Println("Reservation slot not available for this time.")
		return
	}
	reservations[date] = append(reservations[date], Reservation{name, date, timeStr})
	fmt.Println("Reservation confirmed!")
}

func viewReservations(date string) {
	if reservations[date] == nil {
		fmt.Println("No reservations for this date.")
		return
	}
	fmt.Printf("Reservations for %s:\n", date)
	for _, res := range reservations[date] {
		fmt.Printf("Customer: %s, Time: %s\n", res.CustomerName, res.Time)
	}
}

func cancelReservation(name, date, timeStr string) {
	if reservations[date] == nil {
		fmt.Println("No reservations found for this date.")
		return
	}

	var updatedReservations []Reservation
	found := false
	for _, res := range reservations[date] {
		if res.CustomerName == name && res.Time == timeStr {
			found = true
		} else {
			updatedReservations = append(updatedReservations, res)
		}
	}

	if found {
		reservations[date] = updatedReservations
		fmt.Println("Reservation canceled!")
	} else {
		fmt.Println("Reservation not found.")
	}
}

func isTimeSlotAvailable(date, timeStr string) bool {
	if reservations[date] == nil {
		return true
	}

	count := 0
	for _, res := range reservations[date] {
		if res.Time == timeStr {
			count++
		}
	}
	return count < maxReservations
}

func isValidTime(timeStr string) bool {
	t, err := time.Parse("15:04", timeStr)
	if err != nil {
		return false
	}
	hour := t.Hour()
	return hour >= openingHour && hour < closingHour
}

func main() {
	for {
		fmt.Println("Choose an option:")
		fmt.Println("1 - Add Reservation")
		fmt.Println("2 - View Reservations")
		fmt.Println("3 - Cancel Reservation")
		fmt.Println("4 - Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name, date, timeStr string
			fmt.Println("Enter customer name:")
			fmt.Scan(&name)
			fmt.Println("Enter reservation date (YYYY-MM-DD):")
			fmt.Scan(&date)
			fmt.Println("Enter reservation time (HH:MM):")
			fmt.Scan(&timeStr)

			if !isValidTime(timeStr) {
				fmt.Println("Invalid reservation time. Please choose a time between 10:00 and 22:00.")
				continue
			}

			addReservation(name, date, timeStr)

		case 2:
			var date string
			fmt.Println("Enter date (YYYY-MM-DD):")
			fmt.Scan(&date)
			viewReservations(date)

		case 3:
			var name, date, timeStr string
			fmt.Println("Enter customer name:")
			fmt.Scan(&name)
			fmt.Println("Enter reservation date (YYYY-MM-DD):")
			fmt.Scan(&date)
			fmt.Println("Enter reservation time (HH:MM):")
			fmt.Scan(&timeStr)

			if !isValidTime(timeStr) {
				fmt.Println("Invalid reservation time. Please choose a time between 10:00 and 22:00.")
				continue
			}

			cancelReservation(name, date, timeStr)

		case 4:
			fmt.Println("Exiting the system.")
			return

		default:
			fmt.Println("Invalid choice. Please enter 1, 2, 3, or 4.")
		}
	}
}
