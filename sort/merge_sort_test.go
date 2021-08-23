package sort

import (
	"math/rand"
	"testing"
	"time"

	"github.com/phpHavok/go-algorithms/io"
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

func TestMergeSortEmpty(t *testing.T) {
	output := MergeSort([]int{})
	if len(output) != 0 {
		t.Fatalf("output should be empty: %v", output)
	}
}

func TestMergeSortSingleton(t *testing.T) {
	const value = 1337
	output := MergeSort([]int{value})
	if len(output) != 1 || output[0] != value {
		t.Fatalf("output should be a singleton containing %d: %v", value, output)
	}
}

func TestMergeSortOneMillionNoDuplicates(t *testing.T) {
	const numElements = 1000000
	original := io.GenerateSortedIntSlice(getRandomDeterministic(), numElements, 0.0)
	shuffled := append(make([]int, 0, len(original)), original...)
	getRandomNondeterministic().Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	sorted := MergeSort(shuffled)
	if !io.CompareSlices(sorted, original) {
		t.Fatal("slice should be sorted but isn't")
	}
}
func TestMergeSortOneMillionWithDuplicates(t *testing.T) {
	const numElements = 1000000
	original := io.GenerateSortedIntSlice(getRandomDeterministic(), numElements, 0.5)
	shuffled := append(make([]int, 0, len(original)), original...)
	getRandomNondeterministic().Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	sorted := MergeSort(shuffled)
	if !io.CompareSlices(sorted, original) {
		t.Fatal("slice should be sorted but isn't")
	}
}
