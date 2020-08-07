package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel)
	}
	for l := range channel {
		go func(l string) {
			time.Sleep(time.Second * 5) // remember there's differences when sleep was put in the main routine or child go routine
			checkLink(l, channel)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
