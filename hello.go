package main

import (
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

// var errRequestFailed = errors.New("request failed")

func main() {
	// map should be initialized before putting values in
	// var results = make(map[string]string)
	results := map[string]string{}

	c := make(chan result)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- result) {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "OK"
	if (err != nil) || (resp.StatusCode >= 400) {
		fmt.Println(err, resp.StatusCode)
		status = "FAILED"
	}
	c <- result{url: url, status: status}
}
