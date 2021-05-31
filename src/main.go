/* go 파일을 컴파일하기 위해서 이름은 무조건 main.go로 지어야함 */

/* 활용 패키지 명시해줘야함 */
package main

import (
	"fmt"
)

/* go는 compile할때 특정 function을 찾는다. == go의 시작점 */
func main() {
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

}
