package domain

import (
	"fmt"
	"regexp"
)

func (d *domain) NormalizePhone(phone string) (string, error) {

	var normalized string

	for letter := range phone {
		if phone[letter] >= '0' && phone[letter] <= '9' {
			normalized += string(phone[letter])
		}
	}

	if normalized[0] == '8' {
		normalized = "+7" + normalized[1:]
	}

	if normalized[0] == '7' {
		normalized = "+" + normalized
	}

	normalized = normalized[:2] + "-" + normalized[2:5] + "-" + normalized[5:8] + "-" + normalized[8:]

	return normalized, nil

}

func (d *domain) ValidatePhone(phone string) (bool, error) {
	patterns := []string{
		`^(\+7|7|8)([0-9]{10})$`,
		`^(\+7|7|8)(\(|)([0-9]){3}(\)|)([0-9]{3})(\-|)([0-9]{2})(\-|)([0-9]{2})$`,
		`^(\+7|7|8)(\(|)([0-9]){3}(\)|)([0-9]{3})([0-9]{4})$`,
		`^(\+7|7|8)([0-9]){3}([0-9]{3})([0-9]{4})$`,
		`^(\+7|7|8) ([0-9]{3}) ([0-9]{3}) ([0-9]{4})$`,
		`^(\+7|7|8) ([0-9]{3}) ([0-9]{3}) ([0-9]{2}) ([0-9]{2})$`,
	}

	for _, pattern := range patterns {
		pattern, err := regexp.Compile(pattern)
		if err != nil {
			return false, fmt.Errorf("regexp.Compile: %w", err)
		}
		matches := pattern.MatchString(phone)
		if matches {
			return matches, nil
		}
	}

	return false, nil
}
