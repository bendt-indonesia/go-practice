package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string {
		"https://www.google.com",
		"https://www.bendt.io",
		"https://www.amazon.com",
		"https://www.facebook.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go fetchUrl(url, c)
	}

	//forever loop to always fetch the channel
	for l := range c {
		time.Sleep(time.Second * 5)
		go fetchUrl(l, c)
	}
}

func fetchUrl (url string, c chan string) {
	fmt.Println("Fetching: ", url)
	_, err := http.Get(url)

	if err != nil {
		c <- url
		return
	}

	c <- url
	return
}
