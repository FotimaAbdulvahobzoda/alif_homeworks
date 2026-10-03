package validator

import "strings"

func IsValidEmail(email string) bool {
	if !strings.Contains(email, "@") {
		return false
	}

	parts := strings.Split(email, "@")

	if len(parts) != 2 {
		return false
	}

	domain := parts[1]

	dotIndex := strings.LastIndex(domain, ".")

	if dotIndex == -1 {
		return false
	}

	if dotIndex == len(domain)-1 {
		return false
	}

	return true
}

func IsAdult(age int) bool {
	if age >= 18 {
		return true
	} else {
		return false
	}
}
