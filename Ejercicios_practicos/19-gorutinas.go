/* generar un race condition
usando go  rutinas crear un programa incremento
lanzar varias go rutinas que lean el valor del incremento,
lo almacene en una nueva variable
ceda el procesador
incrementa la nueva variable
escribir el valor de la nueva variable de vuelta a la variable incremento
usar waitgorups
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	incremento := 0
	conGs := 100
	fmt.Println("las CPU : ", runtime.NumCPU())
	fmt.Println("Gorutinas: ", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < conGs; i++ {
		go func() {
			nuevaVar := incremento
			runtime.Gosched()
			nuevaVar++
			incremento = nuevaVar
			wg.Done()
		}()
		fmt.Println("valor de i: ", i)
		fmt.Println("Numero de gorutina:", runtime.NumGoroutine())
	}

	fmt.Println("valor de incremento es: ", incremento)

	wg.Wait()
}
