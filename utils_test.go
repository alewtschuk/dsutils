package dsutils

import (
	"fmt"
	"math"
	"testing"
)

func TestRemoveSliceString(t *testing.T) {
	var slice []string
	var numRemoved int

	// For each time a number is divisble by 10 put "WOW" at i position in slice
	for i := 1; i <= 100; i++ {
		if i%10 == 0 {
			slice = append(slice, "WOW")
			numRemoved++
		} else {
			slice = append(slice, fmt.Sprintf("%d", i))
		}
	}
	fmt.Printf("Slice before remove: %v\n\n", slice)
	slice = Remove("WOW", slice)
	fmt.Printf("Slice after remove: %v\n", slice)
	if numRemoved != 10 {
		t.Errorf("ERROR: %d out of 100 elements should have been removed. Length of slice post removal is %d", numRemoved, len(slice))
	}
}

func TestRemoveSliceInt(t *testing.T) {
	var slice []int
	var numRemoved int
	for i := 1; i <= 100; i++ {
		if i%10 == 0 {
			slice = append(slice, 500)
			numRemoved++
		} else {
			slice = append(slice, i)
		}
	}
	fmt.Printf("Slice before remove: %v\n\n", slice)
	slice = Remove(500, slice)
	fmt.Printf("Slice after remove: %v\n", slice)
	if numRemoved != 10 {
		t.Errorf("ERROR: %d out of 100 elements should have been removed. Length of slice post removal is %d", numRemoved, len(slice))
	}
}

func TestRemoveSliceFloat64(t *testing.T) {
	var slice []float64 = []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0}
	var numRemoved int
	for i := 0; i < len(slice); i++ {
		var md float64 = math.Mod(slice[i]+0.00001, 0.2) //Using the math.Mod and the small number 0.00001 to push the resulting mod into range
		if -0.001 < md && md < 0.001 {                   //md must be in between the interval (-0.001, 0.001) of small float number as it wont be exactly 0
			slice[i] = 10.0
			numRemoved++
		}
	}
	fmt.Printf("Slice before remove: %v\n\n", slice)
	slice = Remove(10.0, slice)
	fmt.Printf("Slice after remove: %v\n", slice)
	if numRemoved != 5 {
		t.Errorf("ERROR: %d out of 10 elements should have been removed. Length of slice post removal is %d", numRemoved, len(slice))
	}
}

func TestEnsureUniqueStrings(t *testing.T) {
	var slice []string = []string{"wow", "hello", "super", "uber", "apple", "byte", "apple", "banana", "uber", "pickle", "brass", "choice", "super", "melon"}
	var numRepeats int

	fmt.Printf("Slice before uniqueness is ensured: %v\n\n", slice)
	slice = EnsureUnique(slice)

	// Create map to store if a word exists
	var seen map[string]int = make(map[string]int)

	// For each word in the slice add the word to the map and iterate the associated int value by one for each repeat
	// This ensures that the words will only be counted when they appear and utilizng a map ensures that there are no int itterations
	// by self comparison
	for _, word := range slice {
		seen[word]++
		if seen[word] > 1 {
			numRepeats++
		}
	}
	fmt.Printf("Slice after uniqueness is ensured: %v\n\n", slice)

	if numRepeats != 0 {
		t.Errorf("ERROR: uniqeness not ensured. %d repeats in slice", numRepeats)
	}
}

func TestEnsureUniqueInt(t *testing.T) {
	var slice []int = []int{1, 2, 3, 2, 5, 6, 7, 8, 7, 4, 10, 100, 110, 5, 100, 5000}
	var numRepeats int

	fmt.Printf("Slice before uniqueness is ensured: %v\n\n", slice)
	slice = EnsureUnique(slice)

	// Create map to store if a word exists
	var seen map[int]int = make(map[int]int)

	for _, num := range slice {
		seen[num]++
		if seen[num] > 1 {
			numRepeats++
		}
	}
	fmt.Printf("Slice after uniqueness is ensured: %v\n\n", slice)

	if numRepeats != 0 {
		t.Errorf("ERROR: uniqeness not ensured. %d repeats in slice", numRepeats)
	}
}

func TestEnsureUniqueFloat32(t *testing.T) {
	var slice []float32 = []float32{1, .1, .2, 1, .3, .3, .3, 5.3, 5.55, 10.5, 100.0, 45.7}
	var numRepeats int

	fmt.Printf("Slice before uniqueness is ensured: %v\n\n", slice)
	slice = EnsureUnique(slice)

	// Create map to store if a word exists
	var seen map[float32]int = make(map[float32]int)

	for _, num := range slice {
		seen[num]++
		if seen[num] > 1 {
			numRepeats++
		}
	}
	fmt.Printf("Slice after uniqueness is ensured: %v\n\n", slice)

	if numRepeats != 0 {
		t.Errorf("ERROR: uniqeness not ensured. %d repeats in slice", numRepeats)
	}
}
