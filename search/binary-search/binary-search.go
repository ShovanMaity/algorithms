package samplepackage

// binarySearch returns location of x in given array if is present, otherwise -1
func binarySearch(searchArray []int64, startIndex int, endIndex int, searchElement int64) int {
	if startIndex > endIndex {
		return -1
	}
	mid := startIndex + (endIndex-startIndex)/2
	if searchArray[mid] == searchElement {
		return mid
	}
	if searchArray[mid] > searchElement {
		return binarySearch(searchArray, startIndex, mid-1, searchElement)
	}
	if searchArray[mid] < searchElement {
		return binarySearch(searchArray, mid+1, endIndex, searchElement)
	}
	return -1
}
