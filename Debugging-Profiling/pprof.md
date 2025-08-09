# PPROF
- [pprof](https://github.com/google/pprof) is a tool for visualization and analysis of profiling data.
- pprof reads a collection of profiling samples in profile.proto format and generates reports to visualize and help analyze the data. 
- It can generate both text and graphical reports (through the use of the dot visualization package).

# Profiling types

|                                                                | Package                                                            |
|----------------------------------------------------------------|--------------------------------------------------------------------|
| Runtime Profiling                                              | runtime/pprof                                                      |
| Active Web Profiling                                           | net/http/pprof                                                     |

# Predefined Parameters

| Parameters  |
|-------------|
| CPU         |
| Heap/Memory |
| Block       |
| Thread      |
| Goroutine   |
| Mutex       |

# Commands

| Command                  | Remarks                                   |
|--------------------------|-------------------------------------------|
| pprof.Lookup("block")    | Profiles the blocking routines            |
| go tool pprof block.prof | Opens pprof shell for given profiled file |