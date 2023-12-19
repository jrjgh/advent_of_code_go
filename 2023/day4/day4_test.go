package day4

import "testing"

func TestPartOne(t *testing.T) {
	expected := 19135
	actual := PartOne()

	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 5704953
	actual := PartTwo()

	if expected != actual {
		t.Fatalf("expected %d but got %d", expected, actual)
	}
}
func TestCard_Score(t *testing.T) {

}

func TestParseCard(t *testing.T) {
	cases := []struct {
		input          string
		expected       *card
		winningNumbers int
		playerNumbers  int
	}{
		{
			input:          "Card 1: 1 2 | 1 3",
			winningNumbers: 2,
			playerNumbers:  2,
			expected: &card{
				id: 1,
				winningNumbers: map[int]struct{}{
					1: struct{}{},
					2: struct{}{},
				},
				playerNumbers: []int{1, 3},
			},
		},
		{
			input:          "Card 1: 1 2 3 | 1 3",
			winningNumbers: 3,
			playerNumbers:  2,
			expected: &card{
				id: 1,
				winningNumbers: map[int]struct{}{
					1: struct{}{},
					2: struct{}{},
					3: struct{}{},
				},
				playerNumbers: []int{1, 3},
			},
		},
	}

	for _, testcase := range cases {
		actual := parseCard(testcase.input, testcase.winningNumbers, testcase.playerNumbers)
		if actual.id != testcase.expected.id {
			goto Fail
		}
		if len(testcase.expected.winningNumbers) != len(actual.winningNumbers) {
			goto Fail
		}
		for k, _ := range testcase.expected.winningNumbers {
			if _, ok := actual.winningNumbers[k]; !ok {
				goto Fail
			}
		}
		if len(testcase.expected.playerNumbers) != len(actual.playerNumbers) {
			goto Fail
		}
		for i, _ := range testcase.expected.playerNumbers {
			if testcase.expected.playerNumbers[i] != actual.playerNumbers[i] {
				goto Fail
			}
		}
		continue
	Fail:
		t.Fatalf("expected %+v but got %+v", testcase.expected, actual)
	}
}
