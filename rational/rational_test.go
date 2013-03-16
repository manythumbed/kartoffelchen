package rational

import (
	"sort"
	"testing"
)

func TestAddition(t *testing.T) {
	a := [][]Rational{
		[]Rational{Zero, {1, 1}, {1, 1}},
		[]Rational{{1, 1}, Zero, {1, 1}},
		[]Rational{{1, 1}, {1, 1}, {2, 1}},
		[]Rational{{1, 2}, {1, 2}, {1, 1}},
		[]Rational{{1, 2}, {1, 3}, {5, 6}},
		[]Rational{{2, 21}, {1, 6}, {11, 42}},
	}

	for _, data := range a {
		if x := Add(data[0], data[1]); x != data[2] {
			t.Errorf("When adding %d and %d expected %d and received %d", data[0], data[1], data[2], x)
		}
	}
}

func TestScale(t *testing.T) {
	a := [][]Rational{
		[]Rational{{1, 1}, {1, 1}, {1, 1}},
		[]Rational{{1, 2}, {1, 2}, {1, 4}},
		[]Rational{{1, 2}, {7, 3}, {7, 6}},
		[]Rational{{2, 21}, {1, 6}, {1, 63}},
	}

	for _, data := range a {
		if x := Scale(data[0], data[1]); x != data[2] {
			t.Errorf("When scaling %d by %d expected %d and received %d", data[0], data[1], data[2], x)
		}
	}
}

func TestSort(t *testing.T) {
	s := Rationals{
		{1, 1},
		{1, 2},
		{1, 3},
		{3, 2},
		{4, 2},
		{1, 4},
		{2, 3},
	}

	s1 := Rationals{
		{1, 4},
		{1, 3},
		{1, 2},
		{2, 3},
		{1, 1},
		{3, 2},
		{4, 2},
	}

	sort.Sort(s)
	for i := range s1 {
		if s1[i] != s[i] {
			t.Errorf("Value of sorted array, %v, is not equal to expected value, %v", s[i], s1[i])
		}
	}
}

func TestGcd(t *testing.T) {
	a := [][]int{
		[]int{54, 24, 6},
		[]int{42, 56, 14},
		[]int{56, 42, 14},
		[]int{18, 84, 6},
	}

	for _, data := range a {
		if gcd(data[0], data[1]) != data[2] {
			t.Errorf("Expected gcd of %d and %d to be %d, received %d", data[0], data[1], data[2], gcd(data[0], data[1]))
		}
	}
}

func TestReduce(t *testing.T) {
	a := [][]int{
		[]int{2, 3, 2, 3},
		[]int{42, 56, 3, 4},
	}

	for _, data := range a {
		if c, d := reduce(data[0], data[1]); c != data[2] && d != data[3] {
			t.Errorf("Expected reduction of %d and %d to be %d and %d, received %d and %d", data[0], data[1], data[2], data[3], c, d)
		}
	}
}

func TestLcm(t *testing.T) {
	if lcm(21, 6) != 42 {
		t.Errorf("For %d and %d expected lcm of %d, received %d", 21, 6, 42, lcm(21, 6))
	}
}
