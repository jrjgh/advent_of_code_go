package day5

import "github.com/jrjgh/advent_of_code/util"

func nice1(s string) bool {
	return contains3Vowels(s) &&
		containsDoubleLetter(s) &&
		noNaughtySubstring(s)
}

func nice2(s string) bool {
	return containsSpacedDouble(s) && containsNonOverlappingPairPair(s)
}

func contains3Vowels(s string) bool {
	vowelCount := 0
	for _, c := range s {
		if util.IsVowel(c) {
			vowelCount++
			if vowelCount >= 3 {
				return true
			}
		}
	}
	return false
}

func containsDoubleLetter(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func noNaughtySubstring(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		substr := s[i : i+2]
		switch substr {
		case "ab":
			fallthrough
		case "cd":
			fallthrough
		case "pq":
			fallthrough
		case "xy":
			return false
		default:
			continue
		}
	}
	return true
}

func containsNonOverlappingPairPair(s string) bool {
	seen := map[string]int{} //int is the index
	for i := 0; i < len(s)-1; i++ {
		substr := s[i : i+2]
		if index, ok := seen[substr]; ok {
			if index != i {
				return true
			}
		} else {
			seen[substr] = i + 1
		}
	}
	return false
}

func containsSpacedDouble(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}
