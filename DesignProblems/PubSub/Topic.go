package main

type Topic struct {
    Name        string
    Messages    []*Message
    Subscribers map[string]*Subscriber
    mu          sync.Mutex
}