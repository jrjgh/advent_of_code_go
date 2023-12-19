package day5

import (
	"github.com/jrjgh/advent_of_code/util"
	"testing"
)

func TestNice1(t *testing.T) {
	testCases := []struct {
		input  string
		expect bool
	}{
		{"aeioo", true},     // contains 3 vowels, a double letter, no naughty substring
		{"aeiouuu", true},   // contains 3 vowels, a double letter, no naughty substring
		{"aab", false},      // contains 3 vowels, a double letter, but has naughty substring
		{"cdxy", false},     // doesn't contain 3 vowels, doesn't have a double letter, has naughty substring
		{"aaeeii", true},    // contains 3 vowels, a double letter, no naughty substring
		{"pqrs", false},     // doesn't contain 3 vowels, doesn't have a double letter, has naughty substring
		{"abcdpqxy", false}, // contains 3 vowels, a double letter, but has naughty substring
	}

	for _, testCase := range testCases {
		result := nice1(testCase.input)
		if result != testCase.expect {
			t.Errorf("Expected nice1(\"%v\") to be %v, got %v", testCase.input, testCase.expect, result)
		}
	}

	lines, _ := util.ReadLinesFromFile("day5.txt")
	c := 0
	for _, line := range lines {
		if nice1(line) {
			c++
		}
	}
	if c != 238 {
		t.Errorf("Expected day1part1() to be %v, got %v", 238, c)
	}
}

func TestNice2(t *testing.T) {
	testCases := []struct {
		input  string
		expect bool
	}{
		{"aabcbaa", true},   // contains spaced double "a" and non-overlapping pair "aa"
		{"abcd", false},     // no spaced double, no non-overlapping pair
		{"aaabcdaaa", true}, // contains spaced double "a" and non-overlapping pair "aa"
		{"", false},         // edge case: empty string
		{"abcbabc", true},
		{"xyxabcabc", true}, // contains spaced double "x" and non-overlapping pair "abc"
		{"xyxzabcabc", true},
	}

	for _, testCase := range testCases {
		result := nice2(testCase.input)
		if result != testCase.expect {
			t.Errorf("Expected nice2(\"%v\") to be %v, got %v", testCase.input, testCase.expect, result)
		}
	}

	lines, _ := util.ReadLinesFromFile("day5.txt")
	c := 0
	for _, line := range lines {
		if nice2(line) {
			c++
		}
	}
	if c != 69 {
		t.Errorf("Expected day1part1() to be %v, got %v", 69, c)
	}
}

func TestContainsDoubleLetter(t *testing.T) {
	testCases := []struct {
		input  string
		expect bool
	}{
		{"abcd", false},
		{"aab", true},
		{"xyzzy", true},
		{"abcda", false},
		{"qrsstu", true},
		{"pqrst", false},
		{"", false},
		{"aa", true},
		{"ab", false},
	}

	for _, testCase := range testCases {
		result := containsDoubleLetter(testCase.input)
		if result != testCase.expect {
			t.Errorf("Expected containsDoubleLetter(\"%v\") to be %v, got %v", testCase.input, testCase.expect, result)
		}
	}
}

func TestNoNaughtySubstring(t *testing.T) {
	testCases := []struct {
		input  string
		expect bool
	}{
		{"abcd", false},
		{"efgh", true},
		{"mnopq", false},
		{"rstuvxy", false},
		{"wxyz", false},
		{"ab", false},
		{"cd", false},
		{"pq", false},
		{"xy", false},
		{"ac", true},
		{"bd", true},
		{"pr", true},
		{"xz", true},
	}

	for _, testCase := range testCases {
		result := noNaughtySubstring(testCase.input)
		if result != testCase.expect {
			t.Errorf("Expected noNaughtySubstring(\"%v\") to be %v, got %v", testCase.input, testCase.expect, result)
		}
	}
}

func TestContains3Vowels(t *testing.T) {
	testCases := []struct {
		input  string
		expect bool
	}{
		{"aei", true},
		{"ae", false},
		{"aeio", true},
		{"pqrst", false},
		{"aeiou", true},
		{"aeae", true},
		{"aaa", true},
		{"aabbcc", false},
	}

	for _, testCase := range testCases {
		result := contains3Vowels(testCase.input)
		if result != testCase.expect {
			t.Errorf("Expected contains3Vowels(\"%v\") to be %v, got %v", testCase.input, testCase.expect, result)
		}
	}
}

func TestContainsNonOverlappingPairPair(t *testing.T) {
	testCases := []struct {
		input  string
		expect bool
	}{
		{"aabb", false},
		{"aaaa", true},  // "aa" appears multiple times, but they overlap
		{"abab", true},  // "ab" appears at two different, non-overlapping indexes
		{"xyzxy", true}, // "xy" appears at two different, non-overlapping indexes
		{"abcd", false}, // no pairs appear more than once
		{"", false},     // edge case: empty string
		{"aa", false},   // edge case: one pair, can't overlap or not
	}

	for _, testCase := range testCases {
		result := containsNonOverlappingPairPair(testCase.input)
		if result != testCase.expect {
			t.Errorf("Expected containsNonOverlappingPairPair(\"%v\") to be %v, got %v", testCase.input, testCase.expect, result)
		}
	}
}

func TestContainsSpacedDouble(t *testing.T) {
	testCases := []struct {
		input  string
		expect bool
	}{
		{"abcda", false}, // "a" appears with a space in between
		{"abcd", false},  // no characters appear with a space in between
		{"aa", false},    // edge case: two characters, but no space in between
		{"aaa", true},    // "a" appears with a space in between
		{"", false},      // edge case: empty string
		{"abcbdc", true}, // "b" appears with a space in between
		{"xyzzy", false}, // no characters appear with a space in between
	}

	for _, testCase := range testCases {
		result := containsSpacedDouble(testCase.input)
		if result != testCase.expect {
			t.Errorf("Expected containsSpacedDouble(\"%v\") to be %v, got %v", testCase.input, testCase.expect, result)
		}
	}
}
