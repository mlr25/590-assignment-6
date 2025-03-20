package main

import (
    "fmt"
    "math/big"
    "time"
)

const (
    n = 127
    p = 12
)

func power_h(n *big.Int, p int, res *big.Int) {
    if p == 1 {
        fmt.Println("Result:", res)
        return
    }
    nextRes := new(big.Int).Mul(res, n)
    go power_h(n, p-1, nextRes)         
}

func power(n int, p int) {
    bigN := big.NewInt(int64(n))
    bigRes := big.NewInt(int64(n))
    go power_h(bigN, p, bigRes)
}

func main() {
    power(n, p)
    time.Sleep(2 * time.Second) 
}
