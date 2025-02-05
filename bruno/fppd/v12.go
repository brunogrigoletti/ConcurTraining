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
    n = n + 1
}

func q() {
	n = n + 1
}

func main() {
    go p()
    go q()
	time.Sleep(3 * time.Second)
	fmt.Printf("%d\n", n)
}