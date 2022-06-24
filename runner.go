package main

import "fmt"

func main() {
	result := d1p2()
	fmt.Println("Result: ", result)
}

func d1p1() int64 {
	return day1SumFrequencies(InputLinesToSlice("day1_input.txt"))
}

func d1p2() int64 {
	return day1Rollfrequenciesbeforevaluerepeats(InputLinesToSlice("day1_input.txt"))
}
