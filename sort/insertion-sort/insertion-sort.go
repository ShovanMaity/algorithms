package samplepackage

func insertionSort(array []int64) []int64 {
	arraySize := len(array)
	for i := 1; i < arraySize; i++ {
		key := array[i]
		j := i - 1
		for j >= 0 {
			if array[j] < key {
				break
			}
			array[j+1] = array[j]
			j = j - 1
		}
		array[j+1] = key
	}
	return array
}
