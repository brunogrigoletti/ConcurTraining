// Golang concurrency demonstration code v10 - Sync with waitgroup to wait for all processes

package main

import (
    "fmt"
    "sync"
    "time"
)

var (
    mutex           sync.Mutex     // mutex is used to define a critical section of code
    simulation_time int            = 0
    wg              sync.WaitGroup // wg is used to wait for the program to finish
)

func DoWork(TaskId, iterations int) {
    defer wg.Done() // Schedule the call to Done to tell main we are done
    for i := 1; i <= iterations; i++ {
        mutex.Lock()
        fmt.Printf("%2d:", simulation_time)
        simulation_time++
        for s := 1; s <= TaskId*6; s++ {
            fmt.Printf(" ")
        }
        fmt.Printf("f%d\n", i) // working
        time.Sleep(1 * time.Second)
        mutex.Unlock()
    }
}

func main() {

    wg.Add(4) // Add a count of four, one for each goroutine
    // Esperando quatro eventos acontecerem - isto é, 4 Dones serem enviados pelo sinal da função DoWork
    fmt.Println("\n   Beginning main goroutine")
    fmt.Println("\n        [T1]  [T2]  [T3]  [T4]\n")

    go DoWork(1, 4)
    go DoWork(2, 4)
    go DoWork(3, 4)
    go DoWork(4, 4)

    fmt.Println("   Dispatched work from main goroutine")

    wg.Wait() // Wait for the goroutines to finish
    fmt.Println("   Ending main goroutine")

    // At this point the program execution stops and all
    // active goroutines are killed.
}