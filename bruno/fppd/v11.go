// Diferentes execuções chegam em diferentes resultados

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
    p := 1
    n = p
}

func q() {
    q := 2
    n = q
}

func main() {
    go p()
    go q()
	time.Sleep(3 * time.Second)
	fmt.Printf("%d\n", n)
}