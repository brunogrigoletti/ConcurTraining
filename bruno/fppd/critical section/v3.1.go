// Example code v31.go - Golang critical section with channel demonstration code
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
    counter         int        = 0
)

func IncGlobalCounter(TaskId, iterations int, add_channel chan int) {
    for i := 1; i <= iterations; i++ {
        add_channel <- 1
    }
}

func AddCounter(add_channel chan int) {
    var temp int
    for i := range add_channel { // consuming operations from the chanel (implicit i <- add_channel)
        temp = counter
        temp = temp + i
        counter = temp
    }
}

func main() {
    add_channel := make(chan int)

    fmt.Println("\n  Beginning main goroutine")

    go AddCounter(add_channel) // ready to Add from channel add_channel
    go IncGlobalCounter(2, 20000, add_channel)
    go IncGlobalCounter(3, 20000, add_channel)
    go IncGlobalCounter(4, 20000, add_channel)

    fmt.Println("\n  Hello from main goroutine")

    time.Sleep(10 * time.Second) // give the other goroutine time to finish

    fmt.Printf("Final value of counter: %d\n", counter) // working
    fmt.Println("\n  Ending main goroutine \n")

    // At this point the program execution stops and all
    // active goroutines are killed.
}