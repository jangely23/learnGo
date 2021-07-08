/*
Atomic es u subpaquete de sync, tiene funciones que implementan algoritmos de sincronizacion
y tambien permiten eliminar el race condition
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	var contador int64
	const gs = 100 //numero de go rutinas a iniciars

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)                         //primer parametro memoria de variable a cambiar, segundo parametro el cambio a esa variable
			runtime.Gosched()                                     // yield thread (cede el hilo)
			fmt.Println("Contador:", atomic.LoadInt64(&contador)) //leer esa variable pasando la direccion de memoria
			wg.Done()
		}()
		fmt.Println("Numero de Gorutinas", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Cuenta: ", contador, "\n")
}
