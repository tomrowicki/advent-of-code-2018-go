package day1

import (
	"AdventOfCode2018/common"
	"testing"
)

type testFixture struct {
	inputLines     []string
	expectedOutput int64
}

func TestDay1Part1(t *testing.T) {
	cases := []testFixture{
		{[]string{"+1", "-1"}, 0},
		{[]string{"+1", "+1", "+1"}, 3},
		{[]string{"+1", "+1", "-2"}, 0},
		{[]string{"-1", "-2", "-3"}, -6},
	}

	for _, c := range cases {
		returnedVal := SumFrequencies(c.inputLines)
		if c.expectedOutput != returnedVal {
			t.Errorf("Should be: %d, got: %d!", c.expectedOutput, returnedVal)
		}
	}
}

// +1, -1 first reaches 0 twice.
// +3, +3, +4, -2, -4 first reaches 10 twice.
// -6, +3, +8, +5, -6 first reaches 5 twice.
// +7, +7, -2, -7, -4 first reaches 14 twice.
func TestDay1Part2(t *testing.T) {
	cases := []testFixture{
		{[]string{"+1", "-1"}, 0},
		{[]string{"+3", "+3", "+4", "-2", "-4"}, 10},
		{[]string{"-6", "+3", "+8", "+5", "-6"}, 5},
		{[]string{"+7", "+7", "-2", "-7", "-4"}, 14},
	}

	for _, c := range cases {
		returnedVal := Rollfrequenciesbeforevaluerepeats(c.inputLines)
		if c.expectedOutput != returnedVal {
			t.Errorf("Should be: %d, got: %d!", c.expectedOutput, returnedVal)
		}
	}
}

func BenchmarkD1P2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := Rollfrequenciesbeforevaluerepeats(common.InputLinesToSlice("day1_input.txt"))
		if result != 390 {
			b.Errorf("Result %d does not equal %d", result, 390)
		}
	}
}
