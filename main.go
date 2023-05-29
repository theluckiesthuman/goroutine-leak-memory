package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for {
		go func() {
			s := "This is a memory leak example" // Allocating memory in a goroutine
			fmt.Println(s)
			time.Sleep(time.Second * 10)
		}()
		time.Sleep(time.Second)

		// Analyzing goroutine leaks
		var stats runtime.MemStats
		runtime.ReadMemStats(&stats)
		fmt.Printf("Number of Goroutines: %d\n", runtime.NumGoroutine())
	}
}
