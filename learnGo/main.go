package main

import (
	"fmt"
)

//여러개 반환도 가능
func superAdd(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	return 1
}

func Add(numbers ...int) int {
	total := 0
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		total += i
	}
	return total
}

func main() {
	result := Add(1, 2, 3, 4)
	fmt.Println(result)
}
