package utils

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
