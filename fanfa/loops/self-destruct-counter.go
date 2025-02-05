package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 10; i >= 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
		if i == 0 {
			fmt.Println("KABOOM MF!!")
		}
	}
}
