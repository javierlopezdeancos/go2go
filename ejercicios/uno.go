package main

import (
	"fmt"
)

func main() {

	header := `
    *************************************************************
    *                                                           *
    *        Ejercicio uno: Contador de números impares         *
    *                                                           *
    *************************************************************
    `
	footer := `
    *************************************************************
    *                                                           *
    *                 Resultado del ejercicio                   *
    *                                                           *
    *************************************************************
    `

	error := `
    *************************************************************
    *                                                           *
    *                   1 Error encontrado                      *
    *                                                           *
    *************************************************************
    `

	var uno int
	var dos int
	var unpair int

	fmt.Println(header)

	fmt.Println("Introduce el número uno")
	fmt.Scanln(&uno)
	fmt.Println("Introduce el número dos")
	fmt.Scanln(&dos)

	if uno < dos {
		fmt.Println("\n----------start contador-------------")
		for i := uno; i <= dos; i++ {
			if i%2 != 0 {
				unpair++
				fmt.Printf("(%d) -> %d es un número impar\n", unpair, i)
			}
		}
		fmt.Println("-------------end contador---------------")
		fmt.Println(footer)
		fmt.Printf("Entre %d y %d, tenemos %d números impares\n", uno, dos, unpair)

	} else {
		fmt.Println(error)
		fmt.Print(" El primer número introducido es mayor que el segundo número introducido")
		fmt.Println("\nVuelve a correr el programa corrigiendo este error")
	}

}
