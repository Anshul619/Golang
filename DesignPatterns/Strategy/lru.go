package main

import "log"

type Lru struct{}

func (l *Lru) evict(c *Cache) {
	log.Println("Evicting by LRU")
}
