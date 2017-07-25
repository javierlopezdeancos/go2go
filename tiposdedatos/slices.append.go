package main

import (
	"fmt"
	"../helpers"
)

func main() {

	headers := helpers.GetHeaders()

	title := `*     Los Slices - Append, Añadir elementos a un slice         *`

	fmt.Println(headers.Title)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.TitleSide)
	fmt.Println(title)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.Title)

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