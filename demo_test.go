package main

import "testing"

func TestAddTwoToTwo(t *testing.T) {
	returnedVal := addTwoToTwo()
	if returnedVal != 4 {
		t.Errorf("Should be 2!")
	}
}

func TestAddTwiInts(t *testing.T) {
	cases := []struct {
		inA, inB, out int
	}{
		{1, 2, 3},
		{2, 4, 6},
		{5, 10, 15},
	}

	for _, c := range cases {
		returnedVal := addTwoInts(c.inA, c.inB)
		if c.out != returnedVal {
			t.Errorf("%d + %d does not equal %d!", c.inA, c.inB, returnedVal)
		}
	}
}

func BenchmarkAddTwiInts(b *testing.B) {
	cases := []struct {
		inA, inB, out int
	}{
		{1, 2, 3},
		{2, 4, 6},
		{5, 10, 15},
	}
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			returnedVal := addTwoInts(c.inA, c.inB)
			if c.out != returnedVal {
				b.Errorf("%d + %d does not equal %d!", c.inA, c.inB, returnedVal)
			}
		}
	}
}
