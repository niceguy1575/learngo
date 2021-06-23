/* if else switch 등에 대해서 배워봅시다. */

package part1

func CanIDrink(age int) bool {

	/* if 문 생성 이전에 변수 생성 가능. */
	/* 아래 문장에서는 koreanAge는 if 문에서만 활용되는것을 유추할 수 있다. */
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func CanIDrink2(age int) bool {

	/* switch */
	switch koreanAge := age + 2; koreanAge { /* 변수를 여기에 생성할 수 있다. */
	case 10:
		return false
	case 18:
		return true
	}

	return false
}
