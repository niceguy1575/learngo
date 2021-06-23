package main

import (
	"fmt"
	"learngo/part1"
)

type person struct {
	name    string
	age     int
	favFood []string
}

/* go는 compile할때 특정 function을 찾는다. == go의 시작점 */
func main() {

	/******************************************
		This is part 0. how to write!
	*******************************************/
	fmt.Println("------- This is part 0 -------")

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
	fmt.Println("------- This is part 1 -------")

	part1.Hello()

	fmt.Println(part1.Multiply(2, 2))

	totalLength, upperName := part1.LenAndUpper("niceguy1575")
	fmt.Println(totalLength, upperName)

	/* ignore second variable */
	totalLength2, _ := part1.LenAndUpper("niceguy1575")
	fmt.Println(totalLength2)

	part1.RepeatMe("hi", "hi2", "hi3", "hi4")

	/* naked return => defer */
	/* defer: 함수가 끝난 이후 무언가를 실행시킴*/
	totalLength2, upperName2 := part1.LenAndUpper2("niceguy1575")
	fmt.Println(totalLength2, upperName2)

	/******************************************
		This is part 1-2. for / range / ... args
	*******************************************/

	/* go에는 foreach 등이 없음 #part2.superAdd(1, 2, 3, 4, 5, 6) */
	result := part1.SuperAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)

	/* if else */
	fmt.Println(part1.CanIDrink(16))
	fmt.Println(part1.CanIDrink2(16))

	/******************************************
		This is part 2. pointers
	*******************************************/

	fmt.Println("------- This is part 2 -------")

	/* & 는 변수의 포인터 위치를 나타내도록 함. */
	a := 2
	b := &a
	fmt.Println(&a, b)
	/* '*' 는 해당 포인터가 가리키는 값을 의미함. */

	a = 5 /* a가 값이 바뀌어도, 해당 위치를 바라보는 *b는 값이 안바뀜 */
	fmt.Println(*b)

	*b = 20 /* 해당 값에 직접 접근해 값을 바꾸는것 또한 가능하다 */
	fmt.Println(a)

	/******************************************
		This is part 2-1. array
	*******************************************/

	/* array 선언 */
	names := [5]string{"niceguy1575", "dam", "bin"}
	names[3] = "blah"
	names[4] = "blah"

	fmt.Println(names)

	/* slice를 활용한 array 선언 => array 개수에 제한이 없다. */
	names_slice := []string{"niceguy1575", "dam", "bin"}
	names_slice = append(names_slice, "value_append") /* append names_slice의 값을 바꾸지 않기 때문에 업데이트 선언 (=) 해줘야한다. */

	fmt.Println(names_slice)

	/******************************************
		This is part 2-2. Map
	*******************************************/

	/* python의 dictionary와 유사함. */
	/* key와 value로 구성되어 있으며 map[key]value 의 형태로 이루어짐. */
	niceguy1575 := map[string]string{"name": "niceguy1575", "age": "29"}
	fmt.Println(niceguy1575)

	/* for 를 활용한 반복문 */
	for key, _ := range niceguy1575 {
		fmt.Println(key)
	}

	/******************************************
		This is part 2-3. Struct
	*******************************************/
	favFood := []string{"mara", "ramen"}
	fmt.Println(favFood)

	/* struct 기준으로 명시적인 표현을 해주는것이 좋다. */
	someone := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(someone.name)
}
