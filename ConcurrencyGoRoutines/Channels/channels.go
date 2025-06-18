package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://facebook.com",
		"http://linkedin.com",
		"http://google.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLinkStatus(link, c)
	}

	for l := range c {
		go func(url string) {
			time.Sleep(time.Second)
			checkLinkStatus(url, c)
		}(l)

		//go checkLinkStatus(<-c, c)
		//log.Println(<-c)
	}
}

func checkLinkStatus(url string, c chan string) {

	if _, err := http.Get(url); err != nil {
		log.Println(url, "is down")
		c <- url
		return
	}

	c <- url
	log.Println(url, "is up")
}