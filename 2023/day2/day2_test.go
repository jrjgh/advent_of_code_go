package day2

import "testing"

func TestPartOne(t *testing.T) {
	expected := 2541
	actual := PartOne()
	if actual != expected {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 66016
	actual := PartTwo()
	if actual != expected {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}

func TestParseGameSubset(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected *gamesubset
	}{
		{
			name:  "empty",
			input: "",
			expected: &gamesubset{
				red:   0,
				green: 0,
				blue:  0,
			},
		},
		{
			name:  "all full",
			input: "1 red, 2 green, 3 blue",
			expected: &gamesubset{
				red:   1,
				green: 2,
				blue:  3,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := parseGamesubset(c.input)
			if actual.red != c.expected.red ||
				actual.green != c.expected.green ||
				actual.blue != c.expected.blue {
				t.Fatalf("expected %+v\n but got %+v", c.expected, actual)
			}
		})
	}
}
