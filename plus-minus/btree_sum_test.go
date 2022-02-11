package plus

import "testing"

func TestSumBTree(t *testing.T) {
	cases := []struct {
		name   string
		num    int
		result string
	}{
		{
			name:   "-1",
			num:    -1,
			result: "not possible",
		},
		{
			name:   "0",
			num:    0,
			result: "",
		},
		{
			name:   "1",
			num:    1,
			result: "not possible",
		},
		{
			name:   "11",
			num:    11,
			result: "-",
		},
		{
			name:   "35132",
			num:    35132,
			result: "--+-",
		},
		{
			name:   "35155",
			num:    35133,
			result: "not possible",
		},
		{
			name:   "26712",
			num:    26712,
			result: "-+--",
		},
		{
			name:   "199",
			num:    199,
			result: "not possible",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := PlusMinusBTree(tc.num)
			if res != tc.result {
				t.Errorf("expected %q, but got %q", tc.result, res)
			}
		})
	}
}