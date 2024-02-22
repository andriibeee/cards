package internal

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

func HasCardExpired(month string, year string, date time.Time) bool {
	month = strings.TrimSpace(month)
	year = strings.TrimSpace(year)

	// filter out non-numeric characters
	month = regexp.MustCompile(`[^0-9]`).ReplaceAllString(month, "")
	year = regexp.MustCompile(`[^0-9]`).ReplaceAllString(year, "")

	if month == "" || year == "" {
		return true
	}

	m, err := strconv.Atoi(month)
	if err != nil {
		return true
	}

	y, err := strconv.Atoi(year)
	if err != nil {
		return true
	}

	if 2000 > y {
		y = 2000 + y
	}

	thisYear := y == date.Year()
	if thisYear {
		return m < int(date.Month())
	}
	return y < date.Year()
}

func CheckCreditCardMonth(month string) bool {
	month = strings.TrimSpace(month)

	// filter out non-numeric characters
	month = regexp.MustCompile(`[^0-9]`).ReplaceAllString(month, "")
	if month == "" {
		return false
	}
	n, err := strconv.Atoi(month)
	return err != nil || n >= 1 && n <= 12
}

func CheckCreditCardNumber(number string) bool {
	number = strings.TrimSpace(number)

	// filter out non-numeric characters
	number = regexp.MustCompile(`[^0-9]`).ReplaceAllString(number, "")

	// discard invalid input fast
	if number == "" {
		return false
	}

	// the luhn algo

	var total int64 = 0

	for i := 0; i < len(number); i++ {
		digit, err := strconv.ParseInt(string(number[i]), 10, 0)
		if err != nil {
			return false
		}
		if (i&1)^(len(number)&1) == 0 { // double every second digit
			digit = digit * 2
		}
		if digit > 9 {
			digit = digit - 9
		}
		total += digit
	}

	return (total % 10) == 0
}
