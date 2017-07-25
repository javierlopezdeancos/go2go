package main

import (
	"fmt"
	"../helpers"
)

func main() {

	headers := helpers.GetHeaders()

	title := `*     Los Slices - Append, Añadir elementos a un slice         *`

	fmt.Println(headers.Break)
	fmt.Println(headers.Title)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.TitleSide)
	fmt.Println(title)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.Title)

	fmt.Println(headers.Given)

	x := make([]byte, 4, 10)

	fmt.Println("Al declarar el slice x := make([]byte, 4, 10) -> x= ", x)

	fmt.Println(headers.When)

	x = []byte{'H', 'O', 'L', 'A'}

	fmt.Println("Las 'H' con comillas simples son interpretadas por go como el carácter utf8 de la letra entre comillas")
	fmt.Println("por eso nuestro slice es de tipo byte")

	fmt.Println("\nAl inicializar x = []byte{'H', 'O', 'L', 'A'}, tenemos el slice x = ", x)

	fmt.Println("\nImprimimos x dándole formato UTF-8:")
	fmt.Println("fmt.Printf('el slice x en formato UTF-8 = %q', x)")
	fmt.Printf("el slice x en formato UTF-8 = %q \n", x)

	for i := 0; i < len(x); i++ {
		fmt.Printf("Slice x[%d]: %q \n", i, x[i])
	}

	fmt.Println(headers.Then)
	fmt.Println("El slice x tiene longitud 4 y capaciad 10")
	fmt.Println("si quewmos añadir otro elemento al slice necesitaremos hacerlo con la funcion append")

	fmt.Println("\nAntes de un append")
	fmt.Println("capacidad de x: ", cap(x))
	fmt.Println("longitud de x: ", len(x))
	fmt.Println("Siendo el slice x = ", x)

	fmt.Println(headers.Given)
	fmt.Println("Declarando un slice 'var y []int' sin capacidad ni longitud")

	var y []int

	fmt.Println(headers.When)
	fmt.Println("El slice y = ", y)
	fmt.Println("la longitud del slice =", len(y))
	fmt.Println("la capacidad del slice =", cap(y))

	fmt.Println("\nSi usamos append para añadir elementos...")
	fmt.Println("Append añade el/los elementos hasta que la longitud del slice al añadirlos supere la capacidad")
	fmt.Println("que tiene el slice al momento de aplicar el append")

	fmt.Println("\nAppend, duplicará la capacidad del sice al doble de la longitud que teniamos")
	fmt.Println("antes de proceder al append.")

	for i := 1; i < 6; i++ {
		y = append(y, i)
		fmt.Println("\nAñado", i)
		fmt.Println("Slice y: ", y)
		fmt.Printf("longitud de y = %d, capacidad de y = %d,  elementos de y = %d\n", len(y), cap(y), i)
	}

}