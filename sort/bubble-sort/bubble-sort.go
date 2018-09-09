package samplepackage

func bubbleSort(array []int64) []int64 {
	var swapped bool
	arraySize := len(array)
	for i := 0; i < arraySize-1; i++ {
		swapped = false
		for j := 0; j < arraySize-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return array
}

func recursiveBubbleSort(array []int64, n int) {
	if n == 1 {
		return
	}
	for i := 0; i < n-1; i++ {
		if array[i] > array[i+1] {
			array[i], array[i+1] = array[i+1], array[i]
		}
	}
	recursiveBubbleSort(array, n-1)
}
