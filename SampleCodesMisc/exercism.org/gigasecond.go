package main

// import path for the time package from the standard library
import (
	"log"
	"time"
)

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1000000000 * time.Second)
}

func main() {
	date := time.Date(2011, time.April, 25, 0, 0, 0, 0, time.UTC)
	log.Println(AddGigasecond(date))
}
