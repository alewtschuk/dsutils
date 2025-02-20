package utils

import (
	"fmt"
	"math"
	"testing"
)

func TestRemoveSliceString(t *testing.T) {
	var slice []string
	var numRemoved int
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
