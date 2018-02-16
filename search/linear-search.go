package main

import (
	"fmt"
)

func linearSearch(searchArray []int, searchElement int) int {
	for i, item := range searchArray{
	    if item == searchElement{
		    return i
		}
	}
    return -1
}

func main(){
    var item int
    var size int
    fmt.Println("Enter size of array.")
    fmt.Scan(&size)
    var items = make([]int,size)
    fmt.Println("Enter elements of array")
    for i :=0; i<size; i++ {
        fmt.Scan(&item)
        items[i] = item
    }
    var key int
    fmt.Println("Enter key which you want to search")
    fmt.Scan(&key)
    index := linearSearch(items,key)
    if index == -1 {
        fmt.Println("item not found!")
    }else{
        fmt.Printf("%d found at position %d ",items[index],index+1)
    }
}
