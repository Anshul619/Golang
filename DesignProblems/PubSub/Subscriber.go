package main

type Subscriber struct {
    ID              string
    FilterFunc      func(*Message) bool
    LastConsumedIdx int // Index of last consumed message
}