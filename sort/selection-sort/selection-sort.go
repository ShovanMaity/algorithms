package samplepackage

func selectionSort(array []int64) []int64 {
	arraySize := len(array)
	minIndex := 0
	for i := 0; i < arraySize; i++ {
		minIndex = i
		for j := i + 1; j < arraySize; j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			array[minIndex], array[i] = array[i], array[minIndex]
		}
	}
	return array
}
