package helpers

import "sort"

// Search the given string in the slice of strings and return boolean
func IsStringInSlice(slice []string, value string) bool {
	i := sort.SearchStrings(slice, value)
	return i < len(slice) && slice[i] == value
}

func IsIntInSlice(slice []int, value int) bool {
	i := sort.SearchInts(slice, value)
	return i < len(slice) && slice[i] == value
}
