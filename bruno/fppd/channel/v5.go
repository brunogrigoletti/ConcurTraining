// Golang concurrency demonstration code v5 - Use a channel to wait for all processes
// By Cesar De Rose - 2021

package main

import (
"fmt"
"sync"
"time"
)

var (
mutex1 sync.Mutex // mutex is used to define a critical section of code
simulation_time1 int = 0
)

func DoWork(TaskId, iterations int, done chan bool) {
for i := 1; i <= iterations; i++ {
mutex1.Lock()
fmt.Printf("%2d:", simulation_time1)
simulation_time1++
for s := 1; s <= TaskId*6; s++ {
fmt.Printf(" ")
}
fmt.Printf("f%d\n", i) // working
time.Sleep(1 * time.Second)
mutex1.Unlock()
}
done <- true
}

func main1() {
done := make(chan bool)
fmt.Println("\n Beginning main goroutine")
fmt.Println("\n        [T1] [T2] [T3] [T4]\n")

go DoWork(1, 4, done)
go DoWork(2, 4, done)
go DoWork(3, 4, done)
go DoWork(4, 4, done)

fmt.Println(" Dispatched work from main goroutine")

workers := 4
for i := 1; i <= workers; i++ { // Wait for the goroutines to finish
<-done
}

fmt.Println(" Ending main goroutine")

// At this point the program execution stops and all
// active goroutines are killed.
}