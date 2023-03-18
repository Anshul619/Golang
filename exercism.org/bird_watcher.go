package main

import "log"

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {

    out := 0

    for _, v := range birdsPerDay {
        out += v
    }

    return out
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {

    out := 0

    for i, v := range birdsPerDay {
        if i/7 == week-1 {
            //log.Println
            out += v
        }
    }

    return out
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, _ := range birdsPerDay {
	    if i%2 == 0 {
	        birdsPerDay[i]++
	    }
	}
	return birdsPerDay
}

func main() {
    log.Println(TotalBirdCount([]int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}))
    log.Println(BirdsInWeek([]int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}, 2))
    log.Println(FixBirdCountLog([]int{2, 5, 0, 7, 4, 1}))
}