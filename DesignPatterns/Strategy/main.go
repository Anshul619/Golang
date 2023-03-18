package main

func main() {
	lfu := new(Lfu)
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")

	cache.add("c", "3")

	lru := new(Lru)
	cache.setEvictionAlgo(lru)

	cache.add("d", "4")

	fifo := new(FIFO)
	cache.setEvictionAlgo(fifo)

	cache.add("e", "5")

}
