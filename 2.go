package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	tick := time.Tick(1 * time.Second)
	go func() {
		for {
			select {
			case <- done:
				fmt.Println("Горутина остановлена")
				return
			case <- tick:
				fmt.Println("Программа работает")
			}
		}
	}()
	time.Sleep(5 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
}