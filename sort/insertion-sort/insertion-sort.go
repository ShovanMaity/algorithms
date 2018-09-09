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

func recursiveInsertionSort(array []int64, n int) {
	if n <= 1 {
		return
	}
	recursiveInsertionSort(array, n-1)
	last := array[n-1]
	j := n - 2
	for j >= 0 {
		if array[j] < last {
			break
		}
		array[j+1] = array[j]
		j--
	}
	array[j+1] = last
}
