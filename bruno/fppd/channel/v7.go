// Golang concurrency demonstration code v7 - Calling SO for I/O with channels
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

func Service(com chan int) {
    key := 2
    service := <-com

    if service == 1 { // get a key from keyboard
        fmt.Printf("SO: process want a key from keyboard, waiting for it ...\n") // working
        time.Sleep(5 * time.Second)
        fmt.Printf("SO: got %d from Keyboard\n", key) // working
    }

    // return
    fmt.Printf("SO: return key to calling process\n")
    com <- key
}

func main() {
    so := make(chan int)

    go Service(so)

    // do something

    fmt.Println("P1: Working")

    // ask for I/O - get a number from keyboard

    fmt.Println("P1: Ask SO for keybord imput (function 1) and wait ...")
    so <- 1

    // do something with key

    fmt.Printf("P1: Received %2d from SO\n", <-so)

    // finish

    fmt.Println("P1: Finish")

    // At this point the program execution stops and all
    // active goroutines are killed.
}