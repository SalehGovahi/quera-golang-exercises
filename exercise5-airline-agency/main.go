package main

import (
	"fmt"
)

func main () {
	var numberOfCountries int
	fmt.Scan(&numberOfCountries)

	var countriesMap = make(map[string]string)

	var country string
	var phonePrefix string
	for i := 0; i < numberOfCountries; i++ {
		fmt.Scanf("%s %s",&country, &phonePrefix)
		countriesMap[country] = phonePrefix
	}

	var numberOfPhoneNumbers int
	fmt.Scan(&numberOfPhoneNumbers)

	var phoneNumbers = make([]string,numberOfPhoneNumbers)
	for i := 0; i < numberOfPhoneNumbers; i++ {
		fmt.Scanf("%s",&phoneNumbers[i])
	}

	var resultOfPhoneNumbers = make([]string, numberOfPhoneNumbers)
	for i := 0; i < numberOfPhoneNumbers; i++ {
		var foundMatch = false
		for country, phonePrefix := range countriesMap {
			if phonePrefix == phoneNumbers[i][:3] {
				foundMatch = true
				resultOfPhoneNumbers[i] = country
				break
			} 
		}
		if (!foundMatch) {
			resultOfPhoneNumbers[i] = "Invalid Number"
		}
	}

	for i := 0; i < numberOfPhoneNumbers; i++ {
		fmt.Printf("%s\n", resultOfPhoneNumbers[i])	
	}
}
