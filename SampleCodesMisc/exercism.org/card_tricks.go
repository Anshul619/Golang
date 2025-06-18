package main

import "log"

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2,6,9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index < 0 || index >= len(slice) {
	    return -1
	}
	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
    if index < 0 || index >= len(slice) {
        return append(slice, value)
    }
    slice[index] = value
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
	    return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
    log.Println(FavoriteCards())
    log.Println(GetItem([]int{1, 2, 4, 1}, 2))
    log.Println(GetItem([]int{1, 2, 4, 1}, 10))
    log.Println(SetItem([]int{1, 2, 4, 1}, 2, 6))
    log.Println(SetItem([]int{1, 2, 4, 1}, -1, 6))
    log.Println(PrependItems([]int{3, 2, 6, 4, 8}, 5, 1))
    log.Println(PrependItems([]int{3, 2, 6, 4, 8}))
    log.Println(RemoveItem([]int{3, 2, 6, 4, 8}, 2))
    log.Println(RemoveItem([]int{3, 2, 6, 4, 8}, 11))
}