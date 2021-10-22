package plus

const (
	plus = 0

	notPossible = "not possible"
)

func PlusMinus(n int) string {
	if n < 0 {
		return notPossible
	}
	if n == 0 {
		return ""
	}

	// 123
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	// 19945 -> l = 4
	l := len(digits) - 1

	max := 0
	solution := ""
	for i := 0; i < 1<<l; i++ {
		sum := 0
		minuses := 0
		currentSolution := ""

		for j := 0; j < l; j++ {
			sign := (i >> j) & 1
			if sign == plus {
				sum += digits[j]
				currentSolution += "+"
			} else {
				sum -= digits[j]
				minuses++
				currentSolution += "-"
			}
		}

		if sum == 0 && minuses > max {
			max = minuses
			solution = currentSolution
		}
	}

	if solution == "" {
		return notPossible
	}

	return ""
}
