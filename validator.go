package aefire

import (
	"encoding/xml"
	"strconv"
	"strings"
)

func ValidateLocalCellPhoneNumber(phoneNumber string) bool {
	phoneNumber = strings.TrimSpace(strings.Replace(phoneNumber, "-", "", -1))

	return (strings.HasPrefix(phoneNumber, "010") ||
		strings.HasPrefix(phoneNumber, "011") ||
		strings.HasPrefix(phoneNumber, "016") ||
		strings.HasPrefix(phoneNumber, "017") ||
		strings.HasPrefix(phoneNumber, "018") ||
		strings.HasPrefix(phoneNumber, "019")) &&
		(len(phoneNumber) == 11 || len(phoneNumber) == 10)
}

func ValidateRRN(n string) bool {
	n = strings.Replace(n, "-", "", -1)
	n = strings.Replace(n, " ", "", -1)

	if len(n) != 13 {
		return false
	}

	multipliers := []int{2, 3, 4, 5, 6, 7, 8, 9, 2, 3, 4, 5}
	checksum := 0

	for i, v := range multipliers {
		digit, _ := strconv.Atoi(string(n[i]))
		checksum += v * digit
	}

	checksum = (11 - checksum%11) % 10

	lastDigit, _ := strconv.Atoi(string(n[12]))

	return lastDigit == checksum
}

func ValidateCorpNum(n string) bool {
	n = strings.Replace(n, "-", "", -1)
	n = strings.Replace(n, " ", "", -1)

	if len(n) != 10 {
		return false
	}

	multipliers := []int{1, 3, 7, 1, 3, 7, 1, 3, 5}
	checksum := 0

	for i, v := range multipliers {
		digit, _ := strconv.Atoi(string(n[i]))
		if i < 8 {
			checksum += v * digit % 10
		} else {
			checksum += (v * digit % 10) + (v * digit / 10)
		}
	}

	lastDigit, _ := strconv.Atoi(string(n[9]))
	checksum = checksum + lastDigit

	return checksum%10 == 0
}

func IsValidXML(data []byte) bool {
	return xml.Unmarshal(data, new(interface{})) != nil
}
