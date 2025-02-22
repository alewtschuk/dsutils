// Author: Alex Lewtschuk
package dsutils

// Removes any occuences of the element you wish to delete from a slice
// Returns the slice with the element removed
// Utilizes the comparable constraint to ensure that the element type is comparable
// and that the elements are generics so any type can be used
func Remove[Element comparable](elementToDelete Element, slice []Element) []Element {
	var i int = 0
	for index, element := range slice { // Iterate over the slice, get both the index and element
		if element != elementToDelete { // If the element is not the element we want deleted
			slice[i] = slice[index] // Set the element at index i to the element at index index
			i++
		}
	}
	return slice[:i] // Return the slice up to index i, which is the new length of the slice
}

// Ensures that any slice passed to the function contains no duplicate values in the slice.
// Returns a slice that holds all unique values in the original slice.
// Utilizes a pseudo-set by creating a map holding only index values and an empty struct
// This usage of a map and empty struct reduces memory overhead as struct{}{} takes up no memory
// and map utilization ensures O(1) lookup speed
func EnsureUnique[Element comparable](slice []Element) []Element {
	var set map[Element]struct{} = make(map[Element]struct{}) // Initalizing the map with index generic and empty struct
	var uniqueElements []Element

	//For every element in the slice
	for _, element := range slice {
		//Check if the element exists in the set and if it does not exist
		if _, exists := set[element]; !exists {
			//Add element to the set to ensure uniqueness
			set[element] = struct{}{}
			//Append the unique elements to the slice that will be returned
			uniqueElements = append(uniqueElements, element)
		}
	}
	return uniqueElements
}
