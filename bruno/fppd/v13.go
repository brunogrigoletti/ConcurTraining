// Diferentes execuções chegam no mesmo resultado

package main

import (
    "fmt"
    "sync"
	"time"
)

var (
    mutex          sync.Mutex
    simulationTime int = 0
    n              int = 0
)

func p() {
	temp := n
    n = temp + 1
	fmt.Printf("p.temp: %d\n", temp)
}

func q() {
	temp := n
    n = temp + 1
	fmt.Printf("q.temp: %d\n", temp)
}

func main() {
    go p()
    go q()
	time.Sleep(3 * time.Second)
	fmt.Printf("n: %d\n", n)
}