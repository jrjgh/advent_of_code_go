package day1

import "testing"

func TestPartOne(t *testing.T) {
	n := 54597
	res := PartOne()
	if res != n {
		t.Fatalf("expected %d but got %d", n, res)
	}
}

func TestPartTwo(t *testing.T) {
	n := 54504
	res := PartTwo()
	if res != n {
		t.Fatalf("expected %d but got %d", n, res)
	}
}

type testCase struct {
	name     string
	input    string
	expected int
}

func TestDecode(t *testing.T) {
	cases := []testCase{
		{
			name:     "12",
			input:    "12",
			expected: 12,
		},
		{
			name:     "34",
			input:    "34",
			expected: 34,
		},
		{
			name:     "56",
			input:    "56",
			expected: 56,
		},
		{
			name:     "78",
			input:    "78",
			expected: 78,
		},
		{
			name:     "89",
			input:    "89",
			expected: 89,
		},
		{
			name:     "2a2",
			input:    "1a2",
			expected: 12,
		}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := decode(c.input)
			if res != c.expected {
				t.Fatalf("expected %d but got %d", c.expected, res)
			}
		})
	}
}
func TestDecode2(t *testing.T) {
	cases := []testCase{
		{
			name:     "12",
			input:    "12",
			expected: 12,
		},
		{
			name:     "34",
			input:    "34",
			expected: 34,
		},
		{
			name:     "56",
			input:    "56",
			expected: 56,
		},
		{
			name:     "78",
			input:    "78",
			expected: 78,
		},
		{
			name:     "89",
			input:    "89",
			expected: 89,
		},
		{name: "one2a2",
			input:    "one2a2",
			expected: 12,
		},
		{
			name:     "eightwo2",
			input:    "eightwo2",
			expected: 82,
		},
		{
			name:     "onetwo",
			input:    "onetwo",
			expected: 12,
		},
		{
			name:     "threefour",
			input:    "34",
			expected: 34,
		},
		{
			name:     "fivesix",
			input:    "fivesix",
			expected: 56,
		},
		{
			name:     "seveneight",
			input:    "seveneight",
			expected: 78,
		},
		{
			name:     "eightnine",
			input:    "eightnine",
			expected: 89,
		},
		{
			name:     "eightwo",
			input:    "eightwo",
			expected: 82,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := decode2(c.input)
			if res != c.expected {
				t.Fatalf("expected %d but got %d", c.expected, res)
			}
		})
	}
}
