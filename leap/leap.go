// Package leap check checks for leap yest
package leap

// IsLeapYear returns true if it leap year
func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}
