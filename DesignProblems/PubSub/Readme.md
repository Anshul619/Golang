# In-memory Pub/Sub queue
- Implement a persistent (in memory) pub-sub queue mechanism with guaranteed delivery of every published message for all the subscribed consumers in that subscribed topic in the same order.

# Features:
- Expose topics.
- Publisher to push messages against a topic.
- Subscribe and unsubscribe from a topic.
- Subscriber to consume from a topic. A topic can have multiple subscribers.
- Maintain the state of consumption of each message in each topic for each consumer.
- Maintain order of message consumption for each consumer.
- Filtered message consumption in a topic by the consumer.
- Handle cases of consumers not being available. What to do with those messages?

# Additional Improvements
- Use goroutines + channels for concurrency.
- Add retry logic or dead-letter queue.
- Add REST API layer using gin or net/http.
- Persist state to file periodically (simulate real persistence).
- Track metrics: messages delivered, pending, failed, etc.

# References
- [Razorpay Lead SDE Machine Coding interview](https://leetcode.com/discuss/post/6890687/razorpay-lead-sde-machine-coding-intervi-sy8a/)