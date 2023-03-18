package main

import (
	"log"
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func mergeMaps(in FreqMap, out FreqMap) {
	for k, v := range in {
		out[k] += v
	}
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	out := FreqMap{}
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}

	for _, r := range l {
		wg.Add(1)
		go func(r string) {
			defer wg.Done()
			m := Frequency(r)

			mutex.Lock()
			mergeMaps(m, out)
			mutex.Unlock()

		}(r)
	}
	wg.Wait()
	return out
}

func main() {
	str := []string{"I am a sick man....", "Yes, but here I come to a stop!"}
	log.Println(ConcurrentFrequency(str))
}
