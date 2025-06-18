# Integration Testing
- Test how components work together — includes external systems (DB, API, file system, etc.)

# Characteristics

|                                                     |
|-----------------------------------------------------|
| Slower                                              |
| May be non-deterministic (e.g., network latency)    |
| Real dependencies (Postgres, Redis, Kafka, etc.)    |
| Often lives in /test/ or _integration_test.go files |
