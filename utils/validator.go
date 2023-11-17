package utils

import (
	"regexp"
)

func ValidateEmail(email string) bool {
	var re = regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`)

	return re.MatchString(email)
}

func ValidatePhone(phone string) bool {
	var re = regexp.MustCompile(`^(?m)[0][0-9]{9}$`)

	return re.MatchString(phone)
}

func ValidatePassword(password string) bool {
	var re = regexp.MustCompile(`(?m)^.{8,}$`)

	return re.MatchString(password)
}
