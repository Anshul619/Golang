package main

import "log"

type FIFO struct{}

func (l *FIFO) evict(c *Cache) {
	log.Println("Evicting")
}
