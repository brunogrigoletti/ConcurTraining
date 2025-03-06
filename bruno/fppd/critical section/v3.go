// Example code v3.go - Golang critical section demonstration code
// By Cesar De Rose - 2021

package main

import (
"fmt"
"sync"
"time"
)

var (
mutex sync.Mutex // mutex is used to define a critical section of code
simulation_time int = 0
counter int = 0
)

func IncGlobalCounter(TaskId, iterations int) {
var temp int
for i := 1; i <= iterations; i++ {
//mutex.Lock()
temp = counter
temp = temp + 1
counter = temp
//mutex.Unlock()
}
}

func main() {

fmt.Println("\n Beginning main goroutine")

go IncGlobalCounter(1, 10000)
go IncGlobalCounter(2, 10000)
go IncGlobalCounter(3, 10000)
go IncGlobalCounter(4, 10000)

fmt.Println("\n Hello from main goroutine")

time.Sleep(10 * time.Second) // give the other goroutine time to finish

fmt.Println("\n Ending main goroutine \n")
fmt.Printf("Final value of counter: %d\n", counter) // working
// At this point the program execution stops and all
// active goroutines are killed.
}