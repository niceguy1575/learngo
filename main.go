package main

import (
	"fmt"
	"learngo/part1"
)

/* go는 compile할때 특정 function을 찾는다. == go의 시작점 */
func main() {

	/******************************************
		This is part 0. how to write!
	*******************************************/

	/* 대문자로 시작하는 함수는 다른 패키지로부터 export 해오는 함수임을 암시함 */
	fmt.Println("Hello World!")

	/* 변수 설정, type & 상수/변수 명시 */
	/* 상수는 변경 불가능함 name = "Lynn" (x) */
	const name string = "niceguy1575"
	fmt.Println(name)

	/* 변수 설정 */
	var name2 string = "niceguy1575"
	name2 = "lynn"
	fmt.Println(name2)

	/* 변수의 경우 다음과 같이 작성도 가능함! */
	name3 := "niceguy1575"
	fmt.Println(name3)

	/******************************************
		This is part 1. test function!
	*******************************************/
	part1.Hello()

	fmt.Println(part1.Multiply(2, 2))

	totalLength, upperName := part1.LenAndUpper("niceguy1575")
	fmt.Println(totalLength, upperName)

	/* ignore second variable */
	totalLength2, _ := part1.LenAndUpper("niceguy1575")
	fmt.Println(totalLength2)

	part1.RepeatMe("hi", "hi2", "hi3", "hi4")

}