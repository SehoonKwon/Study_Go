package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {

	urls := []string{
		"https://github.com/SehoonKwon",
		"http://www.samsungcareers.com/main.html",
		"https://www.naver.com",
	}

	for _, url := range urls {
		hitURL(url)
	}
}

var errRequestFailed = errors.New("Request failed")

func hitURL(url string) error {
	fmt.Println("Checking :", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
