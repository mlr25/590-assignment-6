package main

import (
    "fmt"
    "math/big"
    "sync"
    "time"
)

const (
    n = 127
    p = 12
)

var (
    res = big.NewInt(1)
    mu  sync.Mutex
)

func multiply() {
    mu.Lock()
    res.Mul(res, big.NewInt(n)) 
    mu.Unlock()
}

func power(n int, p int) {
    for i := 0; i < p; i++ {
        go multiply()
    }
}

func main() {
    power(n, p)
    time.Sleep(2 * time.Second) 
    fmt.Println("Result:", res)
}
