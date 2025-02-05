// Golang concurrency demonstration code v61 - Producer consumer example with asynchronous channel
// By Cesar De Rose - 2021

package main

import (
    "fmt"
    "sync"
    "time"
)

var (
    mutex           sync.Mutex // mutex is used to define a critical section of code
    simulation_time int        = 0
)

func Producer(TaskId, iterations int, data chan int) {
    for i := 1; i <= iterations; i++ {

        data <- i + TaskId*10

        mutex.Lock()
        fmt.Printf("%2d:", simulation_time)
        simulation_time++
        for s := 1; s <= TaskId*6; s++ {
            fmt.Printf(" ")
        }
        fmt.Printf("P%d\n", i) // working
        time.Sleep(1 * time.Second)
        mutex.Unlock()
    }
}

func Consumer(TaskId int, data chan int) {
    //for i := range data { // consuming values from the chanel (implicit i <- data)
    for i := 1; i <= 9; i++ {
        mutex.Lock()
        fmt.Printf("%2d:", simulation_time)
        simulation_time++
        for s := 1; s <= TaskId*6; s++ {
            fmt.Printf(" ")
        }
        fmt.Printf("C%d\n", <-data) // working
        time.Sleep(1 * time.Second)
        mutex.Unlock()
    }
}

func main() {
    data := make(chan int, 2)

    fmt.Println("\n   Beginning main goroutine")
    fmt.Println("\n        [T1]  [T2]  [T3]  [T4]\n")

    go Consumer(4, data)
    go Producer(1, 3, data)
    go Producer(2, 3, data)
    go Producer(3, 3, data)

    fmt.Println("   Dispatched work from main goroutine")

    time.Sleep(16 * time.Second) // give the other goroutine time to finish

    fmt.Println("   Ending main goroutine")

    // At this point the program execution stops and all
    // active goroutines are killed.
}