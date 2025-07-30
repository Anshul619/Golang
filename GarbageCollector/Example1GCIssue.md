# Example of a GC-related issue in production
- In one production service, we had a Go microservice responsible for ingesting and processing high-frequency telemetry data (tens of thousands of messages per second).
- We noticed spikes in request latency and occasional missed SLAs under heavy load.

# Investigation
- After investigation with [pprof](../Profiling/Readme.md) and [runtime.ReadMemStats](../Profiling/Readme.md), we saw that the service was creating a large number of short-lived small objects (e.g., slices, JSON intermediate structs) per request.
- This led to high allocation rates, causing the GC to trigger frequently.
- Even though Go's GC pause times were short (under a millisecond), the frequent GC cycles increased CPU usage and contributed to latency jitter.

# Solution

|                            | Description                                                                                                                                                           |
|----------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Object pooling             | We introduced **sync.Pool** for temporary buffers and reusable objects (like JSON decoders and pre-allocated slices). This significantly reduced allocation pressure. |
| Zero-copy techniques       | Where possible, we processed incoming data without intermediate allocations — e.g., using json.Decoder directly on streams rather than marshalling/unmarshalling.     |
| Tuned GOGC                 | We experimented with GOGC settings to balance heap growth and GC frequency, ultimately increasing it slightly to reduce GC frequency under high load.                 |
| Allocation profiling in CI | We added allocation benchmarks to ensure future changes didn’t regress.                                                                                               |

# Outcome
- These changes reduced GC-related CPU usage by ~30% and smoothed out latency spikes under load.

# Reference
- [Go Garbage Collector Optimization: Tuning Memory & Reducing GC Pauses](https://medium.com/@jedwaltondev/deep-dive-into-gos-garbage-collector-tuning-memory-reducing-gc-pauses-e00c409f1d39)
