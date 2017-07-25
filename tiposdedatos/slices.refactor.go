package main

import "fmt"

func main() {

	header := `
    *************************************************************
    *                                                           *
    *                      Tipos de datos                       *
    *                          Slices                           *
    *                                                           *
    *************************************************************
    `

	fmt.Println(header)

	// Declaracion
	// slide es un tipo de array dinamico, no sabemos su capacidad
	// la cual, es dinámica
	var s []int
	fmt.Println("Declaramos un slide o un array dinamico")
	fmt.Println("No sabemos/declaramos su longitud var s = []int: ")
	fmt.Println(s)

	// Todos los ejemplos de arrays sirven para los slide, pero nunca
	// declaramos su capacidad
	s1 := []int{0, 1, 2}
	fmt.Println("\nDeclaramos un slide o un array dinamico de 3 elementos")
	fmt.Println("No sabemos/declaramos su longitud var s1 = []int{1,2,3}: ")
	fmt.Println(s1)

	// Declaracion con make(tipo, len, cap)
	// si no le pasamos capacidad, esta será igual a la longitud
	smk := make([]int, 5)
	fmt.Println("\nSlice smk := make([]int, 5) ----> ", smk)
	fmt.Println("Longitud del slide: ", len(smk))
	fmt.Println("Capacidad del slide: ", cap(smk))

	// Declaracion con make(tipo, len, cap)
	// si le pasamos capacidad, tendremos posiciones
	// que podremos usar mas adelante
	smk10 := make([]int, 5, 10)
	fmt.Println("\nSlice smk := make([]int, 5, 10) ----> ", smk10)
	fmt.Println("Longitud del slide: ", len(smk10))
	fmt.Println("Capacidad del slide: ", cap(smk10))

	// Ejercicio

	// creo un array de 10 posiciones
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("\nArray arr: ", arr)

	// creo dos Slices
	var a, b []int
	fmt.Println("\nDefinimos el slice a: ", a)
	fmt.Println("Definimos el slice b: ", b)

	// a apunta al segemto de la posicion 2 a la 5 ( sin incluir esta última ) del array arr
	a = arr[2:5]

	// ahora a tiene los numeros enteros 2,3,4
	fmt.Println("\nSi apuntamos el slice 'a' al segmento del array arr[2:5], entonces a = ", a)

	// b apunta al segemto de la posicion 3 a la 5 (sin incluir esta última) del array arr
	b = arr[3:5]

	// ahora b tiene los numeros enteros 3,4
	fmt.Println("Si apuntamos el slice 'b' al segmento del array arr[3:5], entonces b = ", b)

	// si modificamos el elemento 0 del slide b estamos modificando el elemento del array al que apunta
	b[0] = 10
	// ahora b tiene los numeros enteros 10,4
	fmt.Println("\nSi modificamos la posicion 0 del slice 'b' b[0]=10, entonces b = ", b)
	// y como b es un puntero al array arr entonces el arr = [10]int{0, 1, 2, 10, 4, 5, 6, 7, 8, 9}
	fmt.Println("y como el slice 'b' es un puntero al array 'arr', entonces arr = ", arr)
	// y por lo mismo a tiene los numeros enteros 2,10,4
	fmt.Println("y como el slice 'a' es un puntero al array 'arr', entonces a = ", a)

	// si comprobamos la capacidad del slide a, cuando lo definimos no lo hicimos con make
	// es decir, no le indicamos capacidad
	// capacidad de a = 8
	// Significa que go pone la capacidad desde el primer elemento del slide hasta el final del array al que apunta
	// es decir desde el 2 ........ hasta el ... 9 = 8 posiciones
	fmt.Println("\nLa capacidad de 'a' es ", cap(a))
	fmt.Println("La capacidad de 'b' es si a = arr[2,5] y arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} ")
	fmt.Println("entonces es desde el 2 .... hasta el 9 ( final del array al que apunta ) = 8 posiciones")
	// si comprobamos la longitud del slide a tendremos que es las posiciones ocupadas del slide
	// es decir
	// longitud de a = 3

	// si comprobamos la capacidad del slide b, cuando lo definimos no lo hicimos con make
	// es decir, no le indicamos capacidad
	// capacidad de b = 7
	// Significa que go pone la capacidad desde el primer elemento del slide hasta el final del array al que apunta
	// es decir desde el 3 ........ hasta el ... 9 = 7 posiciones
	fmt.Println("\nLa capacidad de 'b' es ", cap(b))
	fmt.Println("La capacidad de 'b' es si b = arr[3,5] y arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} ")
	fmt.Println("entonces es desde el 3 .... hasta el 9 ( final del array al que apunta ) = 7 posiciones")
	// si comprobamos la longitud del slide b tendremos que es las posiciones ocupadas del slide
	// es decir
	// longitud de b = 2
	fmt.Println("La longitud de 'a' es ", len(b))

	// funcion append
	// agregar elementos a un slice
	appendHeader := `
    *************************************************************
    *                                                           *
    *                      Funcion append                       *
    *                Añadir elementos a un slice                *
    *                                                           *
    *************************************************************
    `

	fmt.Println(appendHeader)

	// declaramos el slice
	x := make([]byte, 4, 10)
	fmt.Println("Al declarar el slice x: ", x)

	// inicializamos x
	// las 'H' con comillas simples son interpretadas por go como el carácter utf8 de la letra entre comillas
	// por eso nuestro slice es de tipo byte
	x = []byte{'H', 'O', 'L', 'A'}
	fmt.Println("Al inicializar x = []byte{'H', 'O', 'L', 'A'},\ntenemos el slice x: ", x)
	// imprimimos x dándole formato UTF-8
	fmt.Printf("Slice x en formato UTF-8: %q \n", x)

	for i := 0; i < len(x); i++ {
		fmt.Printf("Slice x[%d]: %q \n", i, x[i])
	}

	// El Slice x tiene longitud 4 y capacidad 10, si quewmos añadir otro elemento
	// al slice necesitamos hacer con la funcion append
	fmt.Println("El slice x tiene longitud 4 y capaciad 10")
	fmt.Println("\nAntes de un append")
	fmt.Println("capacidad de x: ", cap(x))
	fmt.Println("longitud de x: ", len(x))
	fmt.Println(x)

	// append añade el/los elementos hasta que la longitud del slice al añadirlos supere la capacidad,
	// cuando esto suceda, append, duplicará la capacidad del sice al doble de la longitud que teniamos
	// antes de proceder al append.

	// delcaramos el slice sin capacidad ni longitud
	var y []int

	for i := 1; i < len(y); i++ {

		y = append(y, i)

		// imprimimos el slice
		fmt.Println("Slice y: ", y)

		// imprimimos la longitud y la capacidad del slice según le añadimos con append i
		fmt.Printf("longitud de y %d - capacidad de y %d - elementos %d \n", len(y), cap(y), i)
	}

}
