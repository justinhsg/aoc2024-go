package utils

func NDigits(num int) int {
	if num < 0 {
		return NDigits(-num)
	}
	digits := 0
	for num != 0 {
		num /= 10
		digits += 1
	}
	return digits
}

func Abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func Min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
