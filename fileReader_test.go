package main

import (
	"fmt"
	"testing"
)

func TestReadingLinesAsStrings(t *testing.T) {
	lines := InputLinesToSlice("day1_input.txt")
	fmt.Printf("Lines loaded: %d\n", len(lines))
	if lines == nil {
		t.Errorf("Input is not loaded correctly!")
	}
}
