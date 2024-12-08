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
