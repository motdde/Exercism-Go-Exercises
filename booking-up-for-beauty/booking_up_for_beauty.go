package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	timeStruct, _ := time.Parse("1/02/2006 15:04:05", date)
	return timeStruct
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	timeStruct, _ := time.Parse("January 1, 2006 15:04:05", date)
	return timeStruct.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	timeStruct, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if timeStruct.Hour() >= 12 && timeStruct.Hour() <= 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	timeStruct, _ := time.Parse("1/2/2006 15:04:05", date)
	year, month, day := timeStruct.Date()
	description := fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", timeStruct.Weekday(), month.String(), day, year, timeStruct.Hour(), timeStruct.Minute())
	return description
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	currentYear := time.Now().Year()
	anniversaryDate := time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
	return anniversaryDate
}
