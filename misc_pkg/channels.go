package main

import (
	"fmt"
	"net/http"
	"time"
)

var links = []string{
	"http://google.com",
	"http://facebook.com",
	"http://stackoverflow.com",
	"http://golang.org",
	"http://amazon.com",
}

func iterateUrls(c chan string) {
	for _, link := range links {
		time.Sleep(time.Second)
		checkLink(link, c)
	}

	fmt.Println("Waiting for channel message ....")
	fmt.Println(<-c) // <- Blocking call, Main routine will say "Okay, I'm going to wait until there is a message in the channel"
	// Receiving messages from a channel is a blocking thing.

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
