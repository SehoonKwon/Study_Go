package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {
	var results = map[string]string{}
	c := make(chan requestResult)
	urls := []string{
		"https://github.com/SehoonKwon",
		"http://www.samsungcareers.com/main.html",
		"https://www.naver.com",
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

//chan<- == 이 함수는 send only라고 명시
func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		c <- requestResult{url: url, status: "Failed"}
	} else {
		c <- requestResult{url: url, status: "OK"}
	}
}
