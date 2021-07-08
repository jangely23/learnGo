package main

import (
	"fmt"
	"runtime" //obtener informacion del sistema como # de go rutinas, numero de procesadores, arquitectura.
	"sync"
)

var wg sync.WaitGroup // se crea una variable de tipo waitgroup

func main() {

	fmt.Println("OS\t\t", runtime.GOOS)                 //sistema operativo
	fmt.Println("ARCH\t\t", runtime.GOARCH)             //Arquitectura
	fmt.Println("CPUS\t\t", runtime.NumCPU())           //CPU
	fmt.Println("Goroutines\t", runtime.NumGoroutine()) //Go rutinas

	wg.Add(1) //contador que moitorea cuantas go rutinas se van a lanzar (antes de las go rutinas) se usa con variable.Done()

	go foo() // se realiza una gorutina que finalizara cuando el main termine de ejecutarse si no se usa la varibales.Add(1)
	bar()

	fmt.Println("CPUS\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Wait() // espera que terminen las gorutinsa para finalizar
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}

	wg.Done() //resta una go routina al contador cuando esta finaliza
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
