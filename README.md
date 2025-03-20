# 590-assignment-6

Author: Madison Roberts
PID: 730460151

# Description 
This assignment implements exponentiation (N^P) using concurrency in Go. Two different approaches were taken:
- Prog One (prog1.go): Uses a tail-recursive approach, where each multiplication step spawns a new goroutine until p == 1.
- Prog Two (prog2.go): Spawns p goroutines simultaneously, each multiplying a global accumulator variable.

# Race Condition Analysis 
Did you observe race conditions?
Yes, in Prog Two, running the program with a larger p (e.g., p = 20) sometimes produced incorrect results. This happens because multiple goroutines concurrently modify the shared global res variable without synchronization. Since Go does not guarantee execution order of concurrent processes, some updates to res can be lost or overwritten incorrectly.

How did you resolve the race condition?
To prevent race conditions, we introduced a mutex (sync.Mutex) in prog2.go. The mutex ensures that only one goroutine modifies res at a time, preventing simultaneous writes and data corruption.

Why does prog1.go not have a race condition?
Since Prog One uses goroutines recursively and each goroutine passes its own copy of the accumulator (instead of modifying a shared variable), there are no concurrent writes to the same memory location. This eliminates race conditions by design.

# How to Run: 
go run prog1.go
go run prog2.go