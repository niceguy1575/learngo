/* 숫자를 다뤄봅시다!*/
package part1

import (
	"fmt"
	"strings"
)

/* 변수 뒤에 자료형. */
/* 함수의 리턴값에 대한 자료형 명시 */
/* 패키지화 후 main에서 찾아오기 위해서는 export할 수 있도록 앞을 대문자로 해줘야함 */
func Multiply(a, b int) int {
	return a * b
}

/* multiple return! */
func LenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

/* 여러개의 단어를 받는 방법 */
func RepeatMe(words ...string) {
	fmt.Println(words)
}

func Hello() {
	fmt.Println("hi~")
}
