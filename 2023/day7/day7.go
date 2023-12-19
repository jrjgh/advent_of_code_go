package day7

import (
	u "github.com/jrjgh/advent_of_code/util"
	"sort"
	"strings"
)

type handBid struct {
	hand []int
	bid  int
	typ  int
}

type handBidSort []*handBid

func (s handBidSort) Len() int {
	return len(s)
}

func (s handBidSort) Less(i, j int) bool {
	hand1, hand2 := s[i], s[j]
	if hand1.typ != hand2.typ {
		return hand1.typ < hand2.typ
	}
	for k := 0; k < len(hand1.hand); k++ {
		card1, card2 := hand1.hand[k], hand2.hand[k]
		if card1 != card2 {
			return card1 < card2
		}
	}
	return false
}

func (s handBidSort) Swap(i, j int) {
	s[j], s[i] = s[i], s[j]
}

func PartOne(filename string) int {
	handBids := parseFile(filename)
	sort.Sort(handBidSort(handBids))
	bids := u.Map(func(h *handBid) int {
		return h.bid
	}, handBids)

	return u.SumInt(u.MapI[int, int](rank, bids))
}

func PartTwo(filename string) int {
	handBids := parseFile2(filename)
	sort.Sort(handBidSort(handBids))
	bids := u.Map(func(h *handBid) int {
		return h.bid
	}, handBids)

	return u.SumInt(u.MapI[int, int](rank, bids))
}

func parseFile(filename string) []*handBid {
	lines, _ := u.ReadLinesFromFile(filename)
	return u.Map(parseHandBid, lines)
}

func parseFile2(filename string) []*handBid {
	lines, _ := u.ReadLinesFromFile(filename)
	return u.Map(parseHandBid2, lines)
}

func parseHandBid(line string) *handBid {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		panic("line should have exactly 2 sections")
	}
	handStr, bidStr := parts[0], parts[1]
	hand := parseHand(handStr)
	bid := u.ParseInt(bidStr)
	return &handBid{
		hand: hand,
		bid:  bid,
		typ:  handType(hand),
	}
}

func parseHandBid2(line string) *handBid {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		panic("line should have exactly 2 sections")
	}
	handStr, bidStr := parts[0], parts[1]
	hand := parseHand2(handStr)
	bid := u.ParseInt(bidStr)
	return &handBid{
		hand: hand,
		bid:  bid,
		typ:  handType2(hand),
	}
}

var cardOrder = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

var cardOrder2 = map[string]int{
	"J": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

func string2Card(str string) int {
	return cardOrder[str]
}

func string2Card2(str string) int {
	return cardOrder2[str]
}

func parseHand(handStr string) []int {
	cardStrs := strings.Split(handStr, "")
	cards := u.Map(string2Card, cardStrs)
	return cards
}

func parseHand2(handStr string) []int {
	cardStrs := strings.Split(handStr, "")
	cards := u.Map(string2Card2, cardStrs)
	return cards
}

func handType(hand []int) int {
	tallys := make([]int, 13)
	for _, card := range hand {
		tallys[card-1] += 1
	}
	uniqueCards := u.Reject1[int](u.IsZero[int], tallys)
	biggestMatch, _ := u.Max(uniqueCards)
	switch len(uniqueCards) {
	case 1: //five of a kind
		return 6
	case 2: //full house or four of a kind
		//4 or 5
		if biggestMatch == 4 {
			return 5
		} else {
			return 4
		}
	case 3: //two pair or three of a kind
		if biggestMatch == 3 {
			return 3
		} else {
			return 2
		}
	case 4: //one pair
		return 1
	default: // high card
		return 0
	}
}

func handType2(hand []int) int {
	tallys := make([]int, 13)
	for _, card := range hand {
		tallys[card-1] += 1
	}
	jokers := tallys[0]
	uniqueCards := u.Reject1[int](u.IsZero[int], tallys[1:])
	nuniqueCards := len(uniqueCards)
	biggestMatchX, _ := u.Max(uniqueCards)
	biggestMatch := biggestMatchX + jokers
	switch biggestMatch {
	case 5:
		return 6 //five of a kind
	case 4:
		return 5 //four of a kind
	case 3:
		if nuniqueCards == 2 {
			return 4 //full house
		} else {
			return 3 //three of a kind
		}
	case 2:
		if nuniqueCards == 3 {
			return 2 //two pair
		} else {
			return 1 //one pair
		}
	default:
		return 0
	}
}

func rank(a, b int) int {
	return a * (b + 1)
}
