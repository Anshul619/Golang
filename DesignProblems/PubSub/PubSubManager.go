package main

type PubSubManager struct {
    Topics map[string]*Topic
    mu     sync.Mutex
}

func NewPubSubManager() *PubSubManager {
    return &PubSubManager{
        Topics: make(map[string]*Topic),
    }
}

func (ps *PubSubManager) CreateTopic(name string) {
    ps.mu.Lock()
    defer ps.mu.Unlock()
    if _, exists := ps.Topics[name]; !exists {
        ps.Topics[name] = &Topic{
            Name:        name,
            Subscribers: make(map[string]*Subscriber),
        }
        fmt.Println("Created topic:", name)
    }
}

func (ps *PubSubManager) Publish(topicName, content string) {
    ps.mu.Lock()
    topic, exists := ps.Topics[topicName]
    ps.mu.Unlock()
    if !exists {
        fmt.Println("Topic does not exist")
        return
    }

    topic.mu.Lock()
    defer topic.mu.Unlock()
    msg := &Message{
        ID:        len(topic.Messages),
        Content:   content,
        Timestamp: time.Now(),
    }
    topic.Messages = append(topic.Messages, msg)
    fmt.Printf("Published to %s: %s\n", topicName, content)
}

func (ps *PubSubManager) Subscribe(topicName, subscriberID string, filter func(*Message) bool) {
    ps.mu.Lock()
    topic, exists := ps.Topics[topicName]
    ps.mu.Unlock()
    if !exists {
        fmt.Println("Topic does not exist")
        return
    }

    topic.mu.Lock()
    defer topic.mu.Unlock()
    topic.Subscribers[subscriberID] = &Subscriber{
        ID:              subscriberID,
        FilterFunc:      filter,
        LastConsumedIdx: -1,
    }
    fmt.Printf("Subscriber %s subscribed to topic %s\n", subscriberID, topicName)
}

func (ps *PubSubManager) Unsubscribe(topicName, subscriberID string) {
    ps.mu.Lock()
    topic, exists := ps.Topics[topicName]
    ps.mu.Unlock()
    if !exists {
        fmt.Println("Topic does not exist")
        return
    }

    topic.mu.Lock()
    defer topic.mu.Unlock()
    delete(topic.Subscribers, subscriberID)
    fmt.Printf("Subscriber %s unsubscribed from topic %s\n", subscriberID, topicName)
}

func (ps *PubSubManager) Consume(topicName, subscriberID string) {
    ps.mu.Lock()
    topic, exists := ps.Topics[topicName]
    ps.mu.Unlock()
    if !exists {
        fmt.Println("Topic does not exist")
        return
    }

    topic.mu.Lock()
    defer topic.mu.Unlock()
    sub, exists := topic.Subscribers[subscriberID]
    if !exists {
        fmt.Println("Subscriber not found")
        return
    }

    for i := sub.LastConsumedIdx + 1; i < len(topic.Messages); i++ {
        msg := topic.Messages[i]
        if sub.FilterFunc == nil || sub.FilterFunc(msg) {
            fmt.Printf("Subscriber %s consumed: %s\n", subscriberID, msg.Content)
            sub.LastConsumedIdx = i
            break // consume one message at a time
        }
    }
}

func main() {
    ps := NewPubSubManager()
    ps.CreateTopic("news")

    // Subscribe with a filter
    ps.Subscribe("news", "alice", func(msg *Message) bool {
        return msg.Content != "spam"
    })

    ps.Publish("news", "hello world")
    ps.Publish("news", "spam")
    ps.Publish("news", "breaking news")

    ps.Consume("news", "alice") // should see "hello world"
    ps.Consume("news", "alice") // skips "spam", shows "breaking news"
}
