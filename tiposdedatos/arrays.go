package main

import "fmt"

func main() {

	header := `
    *************************************************************
    *                                                           *
    *                      Tipos de datos                       *
    *                          Arrays                           *
    *                                                           *
    *************************************************************
    `

	// declarar Arrays
	var x [3]int

	fmt.Println(header)
	fmt.Println("var x [3]int es igual a: ")
	fmt.Println(x)

	// asignar valor a un Arrays
	x[1] = 25
	fmt.Println("\nhe asignado el valor 25 al indice 1 del array, x[1]=25: ")
	fmt.Println(x)

	// acceder a un indice concreto
	fmt.Println("\naccedo directamente al indice 1 del array x[1]: ")
	fmt.Println(x[1])

	// declarar e inicializar Arrays
	y := [2]int{5, 10}
	fmt.Println("\nvar y=[2]int{5,10} es igual a: ")
	fmt.Println(y)

	// declarar e inicializar Arrays 2
	z := [...]int{3, 4, 5}
	fmt.Println("\nvar z=[...]int{3,4,5} asignación dinámica del tamaño del array en la declaración: ")
	fmt.Println(z)

	// array de calquier Tipos
	f := [...]float64{0.3, 1.8}
	fmt.Println("\narrays tipo float: ")
	fmt.Println(f)

	// funcion len() indica el tamaño de un array
	fmt.Println("\nel tamaño del array f := [...]float64{0.3, 1.8} es: ")
	fmt.Println(len(f))

	// comparar Arrays
	// el mismo tamaño
	// el mismo Tipo
	// el mismo contenido
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{2, 1}

	fmt.Println("\nsi a := [2]int{1, 2}")
	fmt.Println("si b := [...]int{1, 2}")
	fmt.Println("y si c := [3]int{1, 2, 4}")
	fmt.Println("a = b: ")
	fmt.Println(a == b)
	fmt.Println("a = c: ")
	fmt.Println(a == c)

}
