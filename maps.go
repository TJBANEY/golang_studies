package main

import "fmt"

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for", value, "is", key)
	}
}
