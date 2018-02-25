//Linear Search
//Note: If there is duplicate entry, the search will return the very first element
package main

import (
	"fmt"
)

//Linear Search Function
func linearSearch(searchArray []int, searchElement int) int {
	for i, item := range searchArray {
		if item == searchElement {
			return i
		}
	}
	return -1
}

func main() {
	var item, size int
	fmt.Println("Enter size of array:")
	fmt.Scan(&size)
	var array = make([]int, size)
	fmt.Println("Enter elements of array:")
	for i := 0; i < size; i++ {
		fmt.Scan(&item)
		array[i] = item
	}
	var key int
	fmt.Println("Enter key which you want to search:")
	fmt.Scan(&key)
	index := linearSearch(array, key)
	if index == -1 {
		fmt.Println("Item not found!")
	} else {
		fmt.Printf("%d found at position %d\n", array[index], index+1)
	}
}
