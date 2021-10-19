package plus

func PlusMinus(num int) string {
	var digits []int
	for num > 0 {
		digit := num % 10
		digits = append(digits, digit)
		num = num / 10
	}

	for l := 0; l < len(digits)/2; l++ {
		r := len(digits) - l - 1
		digits[l], digits[r] = digits[r], digits[l]
	}

	nSigns := len(digits) - 1
	n := 1 << nSigns

	var i int
	for i = 0; i < n; i++ {
		var sum = digits[0]
		mapBits(i, nSigns, func(j int, isOne bool) {
			if isOne {
				sum += digits[j+1]
			} else {
				sum -= digits[j+1]
			}
		})

		if sum == 0 {
			break
		}
	}

	var ans string
	if i == n {
		ans = "not possible"
	} else {
		runes := make([]rune, nSigns)
		mapBits(i, nSigns, func(j int, isOne bool) {
			if isOne {
				runes[j] = '+'
			} else {
				runes[j] = '-'
			}
		})
		ans = string(runes)
	}

	return ans
}

func mapBits(bits, nBits int, f func(i int, isOne bool)) {
	for j := 0; j < nBits; j++ {
		powTwo := 1 << j
		sign := bits & powTwo
		f(j, sign > 0)
	}
}
