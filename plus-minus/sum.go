package plus

func PlusMinus(n int) string {
	if n < 0 {
		return "not possible"
	}
	if n == 0 {
		return ""
	}

	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	l := len(digits) - 1

	for i := 0; i < 1<<l; i++ {
		sum := 0
	}

	return ""
}
