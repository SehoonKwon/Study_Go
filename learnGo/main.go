package main

import (
	"fmt"
	"strings"
)

//여러개 반환도 가능
func lenAndUpper(name string) (int, string) {
	defer fmt.Println("I'm done") //함수가 return 된 후에 실행된다 defer!!
	return len(name), strings.ToUpper(name)
}

func main() {

	totalLenth, upperName := lenAndUpper("sehoon")
	fmt.Println(totalLenth, upperName)
}
