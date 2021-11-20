package main

import (
	"fmt"

	"Github.com/SehonKwon/learnGo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "Fisrt Word"}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, err := dictionary.Search(word)
	fmt.Println(hello)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}
