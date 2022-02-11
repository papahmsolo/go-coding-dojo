package plus

// findMinMinusCount finds the string with the maximum count of '-'.
func findMinMinusCount(results []string) string {
	prevC := 0
	max := ""
	for _, s := range results {
		c := 0
		for _, r := range s {
			if r == '-' {
				c++
			}
		}
		if c > prevC {
			prevC = c
			max = s
		}
	}

	return max
}

func PlusMinusBTree(n int) string {
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

	root := Node{value: digits[0], controlSum: digits[0], level: 1}

	for i := 1; i < len(digits); i++ {
		root.AddLevel(digits[i])
	}


	if !root.Check() {
		return "not possible"
	}

	resCh := make(chan string)
	var results []string

	go func() {
		root.FindPath(resCh)
		close(resCh)
	}()

	for s := range resCh {
		s = reverse(s)
		results = append(results, s)
	}

	// Check nodes characteristics
	//root.Print()
	//fmt.Printf("\npossible results: %+v\n", results)
	res := findMinMinusCount(results)
	//fmt.Printf("correct result: %s\n\n", res)

	return res
}