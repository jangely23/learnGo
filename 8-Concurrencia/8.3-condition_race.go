/*
las race condition no son buen codigo ya que da resultados impredecibles.
son dos go rutinas que comparten una variable y ambos tienen acceso simultaneamente
y una vez el que tiene la variable cede el hilo

como funciona, tenemos 2 gorutinas, la gorutina1 lee la variable pausa su ejecucion y cede el hilo ,
al ceder el hilo la gorutina2 lee la varible  pausa su ejecucion y cede el hilo,
posteriormente la gorutina1 continua su ejecucion incrementando la variable en 1 y cede el hilo,
seguido la gorutina2 continua la ejecucion que tenia pausando aumentando en 1.

al no tener control de acceso, ambas vaiables tiene un valor de 1 por que ambas pararon su ejecucion
cuando leyeron la variable en 0, el problema es que la gorutina2 deberia leer cuando la gorutina1
realice el incremento y este en 1, mas no antes.

este ejercicio recrea una race condition.

Ejercicio en el libro Go inaction pagina 141
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

	for i := 0; i < gs; i++ {
		go func() {
			v := contador
			runtime.Gosched() // yield thread (cede el hilo)
			v++
			contador = v
			wg.Done()
		}()
		fmt.Println("Numero de Gorutinas", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Cuenta: ", contador, "\n")
}
