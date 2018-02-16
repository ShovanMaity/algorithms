package samplepackage

//Linear Search Function Note: If there is duplicate entry
// the search will return the very first element
func linearSearch(searchArray []int64, searchElement int64) int {
	for i, item := range searchArray {
		if item == searchElement {
			return i
		}
	}
	return -1
}
