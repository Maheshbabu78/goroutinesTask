package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func foo(c chan int, value int) {
	defer wg.Done()
	c <- value
}
func main() {
	foovalue := make(chan int, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(foovalue, i)

	}
	wg.Wait()
	close(foovalue)
	for data := range foovalue {
		fmt.Println(data)
		time.Sleep(time.Millisecond * 10)
	}
}
func worker(i int) { fmt.Println("doing work on", i) }
