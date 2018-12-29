package codeWars

import (
	"fmt"
	"strings"
)

// Given a string with friends to visit in different states:

// John Daggett, 341 King Road, Plymouth MA
// Alice Ford, 22 East Broadway, Richmond VA
// Sal Carpenter, 73 6th Street, Boston MA

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
	sortdAddrs := make(map[string][]string)
	splitAddrs := strings.Split(addressList, "\n")
	// states := []string{}

	// For address in list of friend's addresses
	for _, address := range splitAddrs {
		splitAddress := strings.Split(address, ",")
		name := strings.Split(splitAddress[0], " ")
		street := splitAddress[1]
		cityState := strings.Split(splitAddress[2], " ")
		city := cityState[1]
		stateAbbr := cityState[2]
		fullState := statesMap[stateAbbr]

		fmt.Println(name[len(name)-2:])
		fmt.Println(len(name))

		reformattedAddress := fmt.Sprintf("%v%v %v %v", name, street, city, fullState)
		sortdAddrs[fullState] = append(sortdAddrs[fullState], reformattedAddress)
	}

	finalString := ""
	for k, v := range sortdAddrs {
		finalString = finalString + fmt.Sprintf("%v\n", k)
		for _, addr := range v {
			finalString = finalString + fmt.Sprintf("..... %v\n", addr)
		}
	}
	return ""
}

// Example result
// Massachusetts
// ..... John Daggett 341 King Road Plymouth Massachusetts
// ..... Sal Carpenter 73 6th Street Boston Massachusetts
// Virginia
// ..... Alice Ford 22 East Broadway Richmond Virginia
