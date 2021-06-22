/* 숫자를 다뤄봅시다!*/
package part1

import (
	"fmt"
	"strings"
)

/* 기본 자료형은 글자, 숫자, bool */

/* 변수 뒤에 자료형. */
/* 함수의 리턴값에 대한 자료형 명시 */
/* 패키지화 후 main에서 찾아오기 위해서는 export할 수 있도록 앞을 대문자로 해줘야함 */
/* 함수가 return하는 것에 대해 명시 해야함. */

func Multiply(a, b int) int {
	return a * b
}

/* multiple return! */
func LenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

/* retrurn을 안한 상태로도 return 할 수 있다. */
func LenAndUpper2(name string) (length int, uppercase string) {
	defer fmt.Println("I am done!")
	length = len(name) /* := 는 변수할당, = 는 업데이트 하는 값임을 의미함 */
	uppercase = strings.ToUpper(name)
	return
}

/* 여러개의 단어를 받는 방법 */
func RepeatMe(words ...string) {
	fmt.Println(words)
}

func Hello() {
	fmt.Println("hi~")
}
