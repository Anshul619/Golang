package main

import "log"

type Lfu struct{}

func (l *Lfu) evict(c *Cache) {
	log.Println("Evicting by LRU")
}
