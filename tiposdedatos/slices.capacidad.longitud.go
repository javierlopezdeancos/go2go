package main

import (
	"fmt"
	"../helpers"
)

func main() {

	headers := helpers.GetHeaders()

	title := `*            Los Slices - Capacidad y Longitud                 *`

	fmt.Println(headers.Title)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.TitleSide)
	fmt.Println(title)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.Title)

	// Vamos a crear  un array de 10 posiciones
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(headers.Given)
	fmt.Println("\nArray arr: ", arr)

	// Si comprobamos la capacidad del slide a, cuando lo definimos no lo hicimos con make
	// es decir, no le indicamos capacidad
	// capacidad de a = 8
	// Significa que go pone la capacidad desde el primer elemento del slide hasta el final del array al que apunta
	// es decir desde el 2 ........ hasta el ... 9 = 8 posiciones
	fmt.Println(headers.When)
	fmt.Println("\nSi declaramos dos slices a := arr[3,5] y b := arr[2,5] sin la función make")

	a := arr[3:5]
	b := arr[2:5]

	fmt.Println(headers.Then)
	fmt.Println("La capacidad de 'a' si a = arr[3,5] es del 3 hasta el 9 ( el final del array que apunta ), capacidad de a = ", cap(a))

	// Si comprobamos la longitud del slide b tendremos que es las posiciones ocupadas del slide
	// es decir
	// longitud de a = 2
	fmt.Println("La longitud de 'a' es ", len(a))

	fmt.Println(headers.Split)

	fmt.Println("La capacidad de 'b' es si b = arr[2,5] y arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} ")
	fmt.Println("entonces es desde del 2 .... hasta el 9 ( final del array al que apunta ), capacidad de b = ", cap(b))

	// Si comprobamos la capacidad del slide b, cuando lo definimos no lo hicimos con make
	// es decir, no le indicamos capacidad
	// capacidad de b = 7

	// Significa que go pone la capacidad desde el primer elemento del slide hasta el final del array al que apunta
	// es decir desde el 3 ........ hasta el ... 9 = 7 posiciones

	// Si comprobamos la longitud del slide a tendremos que es las posiciones ocupadas del slide
	// es decir
	// longitud de b = 3
	fmt.Println("La Longitud de 'b' es ", len(b))

	fmt.Println(headers.Break)

	// Declaracion con make(tipo, len, cap)
	// si le pasamos capacidad, tendremos posiciones
	// que podremos usar mas adelante
	fmt.Println(headers.Given)
	fmt.Println("\nArray arr: ", arr)
	fmt.Println("\nSi declaramos un slice c := make([]int, 5, 10) con la función make")

	c := make([]int, 5, 10)

	fmt.Println(headers.When)
	fmt.Println("\nSlice c := make([]int, 5, 10) ----> ", c)
	fmt.Println("\nc[0] = arr[0]")
	fmt.Println("c[1] = arr[1]")
	fmt.Println("c[2] = arr[2]")
	fmt.Println("c[3] = arr[3]")
	fmt.Println("c[4] = arr[4]")

	c[0] = arr[0]
	c[1] = arr[1]
	c[2] = arr[0]
	c[3] = arr[3]
	c[4] = arr[4]

	fmt.Println(headers.Then)
	fmt.Println("Ahora c = ", c)
	fmt.Println("\nLongitud del slide c: ", len(c))
	fmt.Println("Capacidad del slide c: ", cap(c))

}
