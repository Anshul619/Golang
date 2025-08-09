# Profiling tools

|                                                                | Package                                                            |
|----------------------------------------------------------------|--------------------------------------------------------------------|
| [pprof](pprof.md)                                              | pprof is extensively used to profile go applications.              |
| [Goroutine leaks detection](https://github.com/uber-go/goleak) |                                                                    |
| Benchmark                                                      | `go test -bench="." -cpuprofile='cpu.prof' -memprofile='mem.prof'` |

# Various Debugging Blogs

| Blog                                                                                                                                                                                                               | Remarks |
|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------|
| [Grab - Debugging High Latency Due to Context Leaks](https://engineering.grab.com/debugging-high-latency-market-store)                                                                                             |         |
| [RazorPay - Go Consuming All Your Resources?](https://engineering.razorpay.com/golang-consuming-all-your-resources-5730cac4b61)                                                                                    |         |
| [Debug Go Like a Pro](https://medium.com/better-programming/debug-go-like-a-pro-213d4d74e940)                                                                                                                      |         |
| [How I investigated memory leaks in Go using pprof on a large codebase](https://www.freecodecamp.org/news/how-i-investigated-memory-leaks-in-go-using-pprof-on-a-large-codebase-4bec4325e192/)                     |         |
| [Infoq - Debugging Go Code: Using pprof and trace to Diagnose and Fix Performance Issues](https://www.infoq.com/articles/debugging-go-programs-pprof-trace/)                                                       |         |
| [Detectify - How we tracked down (what seemed like) a memory leak in one of our Go microservices?](https://blog.detectify.com/industry-insights/how-we-tracked-down-a-memory-leak-in-one-of-our-go-microservices/) |         |
| [OVO Energy - Debugging memory leaks using pprof](https://blog.dsb.dev/posts/debugging-memory-leaks-using-pprof/)                                                                                                  |         |

