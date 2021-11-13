package main

import (
	"fmt"
	"strings"
)

//기본적인 함수 생성법
func multiply(a int, b int) int {
	return a * b
}

//인자가 같은 타입일땐 뒤에 한번으로 가능
func multiply2(a, b int) int {
	return a * b
}

//여러개 반환도 가능
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

//인자를 무제한으로 받는법 ... 쓰기
func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println("Hello wolrd") //대문자로 시작하면 export되어 다른 패키지에서 사용가능

	//아래는 동일한 문법이다  : 키워드가 auto 인듯
	var name string = "Sehoon"
	fmt.Println(name)

	name2 := "Sehoon2"
	fmt.Println(name2)

	fmt.Println((multiply(2, 2)))

	totalLenth, upperName := lenAndUpper("sehoon")
	fmt.Println(totalLenth, upperName)

	totalLenth2, _ := lenAndUpper("jihyun") // _는 무시되는 값(인자 숫자맞출때 사용)
	fmt.Println(totalLenth2)                // _을 안했떠라면 오류

	repeatMe("sehoon", "jihyun", "jonghoon")
}
