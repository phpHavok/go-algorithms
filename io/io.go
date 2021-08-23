package io

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

// CompareSlices returns true if s1 and s2 contain identical integers, otherwise
// false.
func CompareSlices(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// GenerateSortedIntSlice will generate a sorted slice of numElements integers,
// each ranging from 0 to numElements - 1. If duplicateProbability is greater
// than 0.0, then with that probability an integer may appear more than once in
// the slice.
func GenerateSortedIntSlice(r *rand.Rand, numElements int, duplicateProbability float64) []int {
	slice := make([]int, 0, numElements)
	for i, elem := 0, 0; i < numElements; i++ {
		slice = append(slice, elem)
		if r.Float64() > duplicateProbability {
			elem++
		}
	}
	return slice
}

// FileToIntSlice reads a file with one integer per line and converts it to a
// slice of integers in memory.
func FileToIntSlice(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	numbers := make([]int, 0)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("unable to convert line to integer: %v", err)
		}
		numbers = append(numbers, num)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("encountered error while reading lines: %v", err)
	}
	return numbers, nil
}
