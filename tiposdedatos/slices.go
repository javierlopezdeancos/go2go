package main

import (
	"fmt"
	"../helpers"
)

func main() {

	headers := helpers.GetHeaders()

	title := `*                       Los Slices                             *`

	fmt.Println(headers.Break)
	fmt.Println(headers.Title)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.TitleSide)
	fmt.Println(title)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.Title)

	var s []int

	fmt.Println(headers.Given)
	fmt.Println("Declaramos un slice o array dinamico, si lo hacemos sin la función make, no sabemos su capacidad ni longitud")
	fmt.Println("no sabemos/declaramos su longitud var s = []int, s = ", s)

	s1 := []int{0, 1, 2}

	fmt.Println("\nDeclaramos un slide o array dinamico de 3 elementos")
	fmt.Println("no sabemos/declaramos su longitud ni capacidad, var s1 = []int{1,2,3}, s1 = ", s1)

	fmt.Println(headers.When)
	fmt.Println("Declaramos un slide o array dinamico de 5 elementos de longitud con make")
	fmt.Println("pero no indicamos su capacidad, está será igual a su longitud")

	smk := make([]int, 5)

	fmt.Println("Slice smk := make([]int, 5) ----> smk = ", smk)

	fmt.Println(headers.Then)
	fmt.Println("Longitud del slide: ", len(smk))
	fmt.Println("Capacidad del slide: ", cap(smk))

	fmt.Println(headers.When)
	fmt.Println("Declaramos un slide o array dinamico con make(tipo, len, cap)")
	fmt.Println("si le pasamos capacidad, tendremos posiciones que podremos usar mas adelante")

	smk10 := make([]int, 5, 10)

	fmt.Println("Slice smk := make([]int, 5, 10) ----> ", smk10)

	fmt.Println(headers.Then)
	fmt.Println("Longitud del slide: ", len(smk10))
	fmt.Println("Capacidad del slide: ", cap(smk10))

	fmt.Println(headers.Break)

	fmt.Println(headers.Given)

	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("Declaramos un slice o array dinamico de 10 posiciones sin make arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, arr = ", arr)

	var a, b []int

	fmt.Println(headers.When)
	fmt.Println("\nDefinimos el slice a = ", a)
	fmt.Println("Definimos el slice b = ", b)

	a = arr[2:5]
	b = arr[3:5]

	fmt.Println("\nApuntamos 'a' al segemto de la posicion 2 a la 5 ( sin incluir esta última ) del array arr, a = ", a)
	fmt.Println("y apuntamos 'b' al segemto de la posicion 3 a la 5 ( sin incluir esta última ) del array arr, b = ", b)

	fmt.Println("\nSi modificamos el elemento 0 del slide 'b', b[0] = 10,  estamos modificando el elemento del array al que apunta")

	b[0] = 10

	fmt.Println(headers.Then)
	fmt.Println("\nSi modificamos la posicion 0 del slice 'b' b[0]=10, entonces b = ", b)
	fmt.Println("y como el slice 'b' es un puntero al array 'arr', entonces arr = ", arr)
	fmt.Println("y como el slice 'a' es un puntero al array 'arr', entonces a = ", a)
	fmt.Println("\nSi comprobamos la capacidad del slide 'a', cuando lo definimos no lo hicimos con make,")
	fmt.Println("es decir, no le indicamos capacidad")
	fmt.Println("\nLa capacidad de 'a' = ", cap(a))
	fmt.Println("La longitud de 'a' = ", len(a))
	fmt.Println("\nLa capacidad de 'a' es si a = arr[2,5] y arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} ")
	fmt.Println("es desde el 2 .... hasta el 9 ( final del array al que apunta ) = 8 posiciones")
	fmt.Println("Si comprobamos la longitud del slice 'a', no lo definimos con make, tendremos que es las posiciones ocupadas del slice")
	fmt.Println("\nSi comprobamos la capacidad del slide 'b', cuando lo definimos no lo hicimos con make,")
	fmt.Println("es decir, no le indicamos capacidad")
	fmt.Println("\nLa capacidad de 'b' = ", cap(b))
	fmt.Println("La longitud de 'b' = ", len(b))
	fmt.Println("\nSignifica que go pone la capacidad desde el primer elemento del slide hasta el final del array al que apunta, ")
	fmt.Println("la capacidad de 'b' es si a = arr[3,4] y arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} ")
	fmt.Println("entonces, en el  caso del slice 'b', es desde el 3 .... hasta el 9 ( final del array al que apunta ) = 7 posiciones")
}
