package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	x := 10
	y := 20
	defer fmt.Println("x value is:", x)
	defer fmt.Println("y value is:", y)
	fmt.Println("logical processor:", runtime.NumCPU())
	fmt.Println("os:", runtime.GOOS)
	fmt.Println("architecture:", runtime.GOARCH)
	fmt.Println("maximum processes:", runtime.GOMAXPROCS(8))
	wg.Add(2)
	go EvenData(&wg)
	go OddData(&wg)
	fmt.Println("invoked main routine")

	wg.Wait()
}
func EvenData(wg *sync.WaitGroup) error {
	fmt.Println("even numbers 1 to 20..")
	numberrange := 20
	for i := 0; i <= numberrange; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
	wg.Done()
	return nil

}
func OddData(wg *sync.WaitGroup) error {
	fmt.Println()
	fmt.Println("Odd numbers 1 to 20..")
	numberrange := 20
	for i := 0; i <= numberrange; i++ {
		if i%2 != 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
	wg.Done()
	return nil
}
