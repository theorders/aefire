package aefire

import (
	"fmt"
	"strings"
)

func NormalizePhoneNumber(phoneNumber string, countryCode int) string {
	phoneNumber = strings.Replace(phoneNumber, "-", "", -1)

	countryCodePrefix := CountryCodePrefix(countryCode)

	if !strings.HasPrefix(phoneNumber, "+") {
		phoneNumber = fmt.Sprintf("%s%s", countryCodePrefix, phoneNumber[1:])
	}

	return phoneNumber
}

func LocalizePhoneNumber(phoneNumber string, countryCode int) (localNumber string) {
	countryCodePrefix := CountryCodePrefix(countryCode)

	if strings.HasPrefix(phoneNumber, countryCodePrefix) {
		return strings.Replace(phoneNumber, countryCodePrefix, "0", 1)
	} else {
		return phoneNumber
	}

	localNumber = strings.Replace(phoneNumber, countryCodePrefix, "0", 1)

	return
}

func LocalizeAndFormatPhoneNumber(phoneNumber string, countryCode int) (localNumber string) {
	localNumber = LocalizePhoneNumber(phoneNumber, countryCode)

	if len(localNumber) == 11 {
		localNumber = fmt.Sprintf("%s-%s-%s", localNumber[:3], localNumber[3:7], localNumber[7:])
	} else if len(localNumber) == 10 {
		localNumber = fmt.Sprintf("%s-%s-%s", localNumber[:3], localNumber[3:6], localNumber[6:])
	}

	return localNumber
}

func CountryCodePrefix(countryCode int) string {
	return fmt.Sprintf("+%d", countryCode)
}
