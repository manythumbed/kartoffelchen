package kartoffelchen

import "testing"

func TestLcm(t *testing.T) {
	if lcm(21, 6) != 42 {
		t.Errorf("For %d and %d expected lcm of %d, received %d", 21, 6, 42, lcm(21, 6))
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
