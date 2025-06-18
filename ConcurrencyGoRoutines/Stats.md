# Go Routine Stats
- Roughly `100K go routines` can run on a single machine (4GB RAM, 2 CPU Core).
- They are very powerful and a simple machine with `1GB of RAM and 1 CPU` can run thousands of goroutines.
- Each GoRoutine consumes [8KB memory](https://stackoverflow.com/questions/22326765/go-memory-consumption-with-many-goroutines).