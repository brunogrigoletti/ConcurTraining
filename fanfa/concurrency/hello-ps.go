package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func ConcurrencyFunc() {
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func() {
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	go func() {
		defer waitGrp.Done()
		fmt.Println("World!")
	}()
	waitGrp.Wait()
}
