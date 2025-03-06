// Example code v8.go - Golang pipeline example with channels
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
    chans                      = []chan int{ // three int channels chan[0], chan[1] and chan[2]
        make(chan int),
        make(chan int),
        make(chan int),
    }
)

func PipeStage(TaskId, iterations int, in chan int, out chan int) {
    for i := 1; i <= iterations; i++ {
        // receive
        temp := <-in
        //process
        temp++
        //mutex.Lock()
        fmt.Printf("%2d:", simulation_time)
        simulation_time++
        for s := 1; s <= TaskId*6; s++ {
            fmt.Printf(" ")
        }
        fmt.Printf("f%d\n", i) // working
        time.Sleep(1 * time.Second)
        //mutex.Unlock()
        //send
        out <- temp
    }
}

func main() {
    // Set up the pipeline.

    go PipeStage(1, 3, chans[2], chans[0])
    go PipeStage(2, 3, chans[0], chans[1])
    go PipeStage(3, 3, chans[1], chans[2])

    // technicaly we have deadlock now
    // Lets'break it and start the action

    chans[2] <- 1

    time.Sleep(12 * time.Second) // give the other goroutine time to finish
}