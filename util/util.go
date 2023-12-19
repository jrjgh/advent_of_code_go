package util

import (
	"bufio"
	"cmp"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadLinesFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func Copy[T any](ts []T) []T {
	n := make([]T, len(ts))
	copy(ts, n)
	return n
}

func ParseInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("could not parse %s as int", str))
	}
	return n
}

func PartitionInt(n int) [][]int {
	ps := make([]int, n+1)
	return MapI(func(_, i int) []int {
		return []int{n - i, i}
	}, ps)
}

func IsVowel(r rune) bool {
	switch r {
	case 'a':
		fallthrough
	case 'e':
		fallthrough
	case 'i':
		fallthrough
	case 'o':
		fallthrough
	case 'u':
		return true
	default:
		return false
	}
}

func IsZero[T int | int32 | int64 | float32 | float64](n T) bool {
	return n == 0
}

func Sign[T int | int32 | int64 | float32 | float64](n T) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 0
	}
	return 1
}

func SumInt(ns []int) int {
	sum := 0
	for _, n := range ns {
		sum += n
	}
	return sum
}

func ProdInt(ns []int) int {
	prod := 1
	for _, n := range ns {
		prod *= n
	}
	return prod
}

func Min[T cmp.Ordered](nums []T) (T, error) {
	if len(nums) == 0 {
		return *new(T), errors.New("list must not be empty")
	}
	m := nums[0]
	for _, t := range nums[1:] {
		m = min(m, t)
	}
	return m, nil
}

func Max[T cmp.Ordered](nums []T) (T, error) {
	if len(nums) == 0 {
		return *new(T), errors.New("list must not be empty")
	}
	m := nums[0]
	for _, t := range nums[1:] {
		m = max(m, t)
	}
	return m, nil
}

func IsNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func Float64(n int) float64 {
	return float64(n)
}

func Int[T float64 | float32](f T) int {
	return int(f)
}

func ForEach[T any](f func(T), ts []T) {
	for _, t := range ts {
		f(t)
	}
}

func MapI[T, U any](f func(T, int) U, ts []T) []U {
	out := make([]U, len(ts))
	for i, t := range ts {
		out[i] = f(t, i)
	}
	return out
}

func Map[T, U any](f func(T) U, ts []T) []U {
	out := make([]U, len(ts))
	for i, t := range ts {
		out[i] = f(t)
	}
	return out
}

func And(b, c bool) bool {
	return b && c
}

func Filter[T any](f func(T) bool, ts []T) []T {
	out := make([]T, 0)
	for _, t := range ts {
		if f(t) {
			out = append(out, t)
		}
	}
	return out
}

func Reject1[T any](f func(T) bool, ts []T) []T {
	return Filter(Neg1(f), ts)
}

func Reduce[T, U any](f func(t T, u U) U, ts []T, acc U) U {
	for _, t := range ts {
		acc = f(t, acc)
	}
	return acc
}

func All[T any](f func(T) bool, ts []T) bool {
	bs := Map[T, bool](f, ts)
	return Reduce(And, bs, true)
}

func Concat[T any](ts1 []T, ts2 []T) []T {
	return append(ts2, ts1...)
}

func Flatten[T any](tts [][]T) []T {
	ts := make([]T, 0)
	return Reduce[[]T, []T](Concat[T], tts, ts)
}

func Compose[T, U, V any](f1 func(t T) U, f2 func(u U) V) func(T) V {
	return func(t T) V {
		return f2(f1(t))
	}
}

func Zip2[T any](ts1, ts2 []T) [][]T {
	return MapI(func(_ T, i int) []T {
		return []T{ts1[i], ts2[i]}
	}, ts1)
}

//func Partition[T any](ts []T, n int) [][]T {
//
//}

func Neg1[T any](f func(t T) bool) func(t T) bool {
	return func(t T) bool {
		return !f(t)
	}
}
