/*
Mutex bloquea el acceso a una variable cuando una gorutina esta operando en ella.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	contador := 0
	const gs = 100 //numero de go rutinas a iniciars

	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock() //bloquea la variable
			v := contador
			runtime.Gosched() // yield thread (cede el hilo)
			v++
			contador = v
			mu.Unlock() //desbloquea
			wg.Done()
		}()
		fmt.Println("Numero de Gorutinas", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Cuenta: ", contador, "\n")
}
