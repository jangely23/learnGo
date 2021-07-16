/*crear tres go rutinas y ejecutarlas todas usando waitgroup*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("Gorutinas ", runtime.NumGoroutine(), "\n")

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		fmt.Println("esta es la gorutina 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("esta es la gorutina 2")
		wg.Done()
	}()

	go func() {
		fmt.Println("esta es la gorutina 3")
		wg.Done()
	}()

	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("Gorutinas ", runtime.NumGoroutine(), "\n")
	fmt.Println("casi finaliza main")

	wg.Wait()

	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("Gorutinas ", runtime.NumGoroutine(), "\n")
}
