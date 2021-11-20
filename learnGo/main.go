package main

import (
	"fmt"

	"Github.com/SehonKwon/learnGo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "Fisrt Word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
