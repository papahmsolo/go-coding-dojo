package plus

import "testing"

func TestSum(t *testing.T) {
	cases := []struct {
		name   string
		num    int
		result string
	}{
		{
			name:   "single digit: non-zero",
			num:    1,
			result: "not possible",
		},
		{
			name:   "two digit: possible",
			num:    11,
			result: "-",
		},
		{
			name:   "positive result",
			num:    35132,
			result: "-++-",
		},
		{
			name:   "multiple solutions",
			num:    26712,
			result: "-+--",
		},
		{
			name:   "not possible",
			num:    199,
			result: "not possible",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := PlusMinus(tc.num)
			if res != tc.result {
				t.Errorf("expected %q, but got %q", tc.result, res)
			}
		})
	}
}
