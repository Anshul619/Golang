package main

import sync

func main() {
    var mu sync.RWMutex
    
    // Multiple readers can hold RLock at the same time
    mu.RLock()
    // read...
    mu.RUnlock()
    
    // Only one writer can hold Lock, excludes readers & writers
    mu.Lock()
    // write...
    mu.Unlock()
}