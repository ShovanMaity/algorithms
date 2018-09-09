package samplepackage

import "math"

func jumpSearch(searchArray []int64, searchEement int64) int {
	arraySize := len(searchArray)
	step := int(math.Sqrt(float64(arraySize)))
	startIndex := 0
	i := func() int {
		if step < arraySize {
			return step
		}
		return arraySize
	}()
	for searchArray[i] < searchEement {
		i = func() int {
			if step < arraySize {
				return step
			}
			return arraySize
		}()
		startIndex = step
		step = step + int(math.Sqrt(float64(arraySize)))
		if startIndex >= arraySize {
			return -1
		}
	}
	for searchArray[startIndex] < searchEement {
		startIndex++
		if startIndex == arraySize || startIndex == step {
			return -1
		}
	}
	if searchArray[startIndex] == searchEement {
		return startIndex
	}
	return -1
}
