package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {

	var results = map[string]string{} //map 초기화방법, 2)는 make 사용하기
	urls := []string{
		"https://github.com/SehoonKwon",
		"http://www.samsungcareers.com/main.html",
		"https://www.naver.com",
	}

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "Failed"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
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
