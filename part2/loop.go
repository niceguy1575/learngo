/* 반복문을 다뤄봅시다!*/
package part2

import (
	"fmt"
)

/* range를 활용해 모든 숫자를 합함 */

func SuperAdd(numbers ...int) int {
	total := 0
	/* range로 for로 할 수 있음 */
	for _, number := range numbers {
		total += number
	}
	return total
}

func Hello() {
	fmt.Println("hi~")
}

func GetValue(numbers ...int) int {
	total := 0
	/* range로 for로 할 수 있음 */
	for _, number := range numbers {
		total += number
	}
	return total
}
