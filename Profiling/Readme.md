# Profiling

|                          | Package                                                            |
|--------------------------|--------------------------------------------------------------------|
| Profiling with benchmark | `go test -bench="." -cpuprofile='cpu.prof' -memprofile='mem.prof'` |
| Runtime Profiling        | runtime/pprof                                                      |
| Active Web Profiling     | net/http/pprof                                                     |

[Read more](https://github.com/google/pprof)

# Predefined Profiles

| Profile     |
|-------------|
| CPU         |
| Heap/Memory |
| Block       |
| Thread      |
| Goroutine   |
| Mutex       |

# pprof Commands

| Command                  | Remarks                                   |
|--------------------------|-------------------------------------------|
| pprof.Lookup("block")    | Profiles the blocking routines            |
| go tool pprof block.prof | Opens pprof shell for given profiled file |

