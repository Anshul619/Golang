package main

import (
	"fmt"
	"log"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t, _ := time.Parse("January 2, 2006 15:04:05", date)
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	formatted, _ := time.Parse("15:04", t.Format("15:04"))

	start, _ := time.Parse("15:04", "12:00")
	end, _ := time.Parse("15:04", "16:00")

	if formatted.After(start) && formatted.Before(end) {
		return true
	}

	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	return fmt.Sprintf("You have an appointment on %s.", Schedule(date).Format("Monday, January 2, 2006, at 15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", fmt.Sprintf("9/15/%d 00:00:00", time.Now().Year()))
	return t
}

func main() {
	log.Println(Schedule("7/25/2019 13:45:00"))
	log.Println(Schedule("7/13/2020 20:32:00"))
	log.Println(Schedule("11/28/1984 2:02:02"))
	log.Println(HasPassed("July 25, 2016 13:45:00"))
	log.Println(IsAfternoonAppointment("Thursday, July 25, 2019 11:45:00"))
	log.Println(Description("7/25/2019 13:45:00"))
	log.Println(AnniversaryDate())
}
