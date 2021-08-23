package io

import (
	"math/rand"
	"testing"
	"time"
)

func getRandomDeterministic() *rand.Rand {
	// Changing this will require a rewrite of some unit tests.
	const deterministicSeed = 31337
	return getRandom(deterministicSeed)
}

func getRandomNondeterministic() *rand.Rand {
	return getRandom(time.Now().UnixNano())
}

func getRandom(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}

// TODO: Write tests for CompareSlices

// TODO: Write tests for FileToIntSlice

func TestGenerateSortedIntSliceEmpty(t *testing.T) {
	slice := GenerateSortedIntSlice(getRandomNondeterministic(), 0, 0.0)
	if len(slice) != 0 {
		t.Fatalf("slice should have no elements: %v", slice)
	}
}

func TestGenerateSortedIntSliceSingleton(t *testing.T) {
	slice := GenerateSortedIntSlice(getRandomNondeterministic(), 1, 0.0)
	if len(slice) != 1 || slice[0] != 0 {
		t.Fatalf("slice should have a single 0: %v", slice)
	}
}

func TestGenerateSortedIntSliceOneHundredAllZero(t *testing.T) {
	const numElements = 100
	const value = 0
	slice := GenerateSortedIntSlice(getRandomNondeterministic(), numElements, 1.0)
	if len(slice) != numElements {
		t.Fatalf("slice should have %d elements: %v", numElements, slice)
	}
	for _, v := range slice {
		if v != value {
			t.Fatalf("all slice elements should be %d: %v", value, slice)
		}
	}
}

func TestGenerateSortedIntSliceOneHundredAllDifferent(t *testing.T) {
	const numElements = 100
	slice := GenerateSortedIntSlice(getRandomNondeterministic(), numElements, 0.0)
	if len(slice) != numElements {
		t.Fatalf("slice should have %d elements: %v", numElements, slice)
	}
	for i := 0; i < numElements; i++ {
		if i != slice[i] {
			t.Fatalf("slice elements should range entire domain: %v", slice)
		}
	}
}

func TestGenerateSortedIntSliceTenWithDuplicates(t *testing.T) {
	const numElements = 10
	deterministicElements := []int{0, 1, 1, 1, 2, 2, 2, 2, 2, 3}
	slice := GenerateSortedIntSlice(getRandomDeterministic(), numElements, 0.5)
	if len(slice) != numElements {
		t.Fatalf("slice should have %d elements: %v", numElements, slice)
	}
	if !CompareSlices(slice, deterministicElements) {
		t.Fatalf("slice %v doesn't match expected %v", slice, deterministicElements)
	}
}
