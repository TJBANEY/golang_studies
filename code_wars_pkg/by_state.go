package codeWars

import (
	"fmt"
	"strings"
)

// Given a string with friends to visit in different states:

// ad3="John Daggett, 341 King Road, Plymouth MA
// Alice Ford, 22 East Broadway, Richmond VA
// Sal Carpenter, 73 6th Street, Boston MA"

// we want to produce a result that sorts the names by state and lists the name of the state
// followed by the name of each person residing in that state (people's names sorted). When the result is printed we get:

var statesMap = map[string]string{
	"AZ": "Arizona",
	"CA": "California",
	"ID": "Idaho",
	"IN": "Indiana",
	"MA": "Massachusetts",
	"OK": "Oklahoma",
	"PA": "Pennsylvania",
	"VA": "Virginia",
}

// ByState - ..
func ByState(addressList string) (stateString string) {
	splitAddresses := strings.Split(addressList, "\n")
	states := []string{}

	for i, address := range splitAddresses {
		splitAddress := strings.Split(address, ",")
		fmt.Println(splitAddress[2])

		if (i+1)%3 == 0 {
			// fmt.Println(addressSection)
			states = append(states, statesMap[address])
		}
	}

	// fmt.Println(states)

	return "some string"
}

// Example result
// Massachusetts
// ..... John Daggett 341 King Road Plymouth Massachusetts
// ..... Sal Carpenter 73 6th Street Boston Massachusetts
// Virginia
// ..... Alice Ford 22 East Broadway Richmond Virginia
