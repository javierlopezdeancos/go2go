package main

import (
	"fmt"
	"../helpers"
)

func main() {

	headers := helpers.GetHeaders()

	title := `*                       Los Arrays                             *`

	fmt.Println(headers.Break)
	fmt.Println(headers.Title)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.TitleSide)
	fmt.Println(title)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.Title)

	fmt.Println(headers.Given)
	fmt.Println("Declaramos el array var x [3]int")
	var x [3]int

	fmt.Println(headers.When)
	fmt.Println("Asignamos un valor al array x[1] = 25")
	x[1] = 25

	fmt.Println(headers.Then)
	fmt.Println("x = ", x)

	fmt.Println(headers.Then)
	fmt.Println("Podemos acceder a un valor de un array x[1] =  ", x[1])

	fmt.Println(headers.Break)

	fmt.Println(headers.Given)
	fmt.Println("Podemos Declarar e inicializar el array y := [2]int{5, 10}")
	y := [2]int{5, 10}

	fmt.Println(headers.Then)
	fmt.Println("y = ", y)

	fmt.Println(headers.Break)

	fmt.Println(headers.Given)
	fmt.Println("Podemos Declarar e inicializar también el array z := [...]int{3, 4, 5}")
	z := [...]int{3, 4, 5}

	fmt.Println(headers.Then)
	fmt.Println("z = ", z)

	fmt.Println(headers.Break)

	fmt.Println(headers.Given)
	fmt.Println("Podemos tener arrays de cualquier otro tipo f := [...]float64{0.3, 1.8}")
	f := [...]float64{0.3, 1.8}

	fmt.Println(headers.Then)
	fmt.Println("f = ", f)


	fmt.Println(headers.Break)

	fmt.Println(headers.Given)
	fmt.Println("Podemos calcular la longitud/tamaño de un array con la funcion len()")

	fmt.Println(headers.Then)
	fmt.Println("len(f) = ", len(f))

	fmt.Println(headers.Break)

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{2, 1}

	fmt.Println(headers.Given)
	fmt.Println("Podemos comparar arrays, son iguales si:")
	fmt.Println("1. El mismo tamaño")
	fmt.Println("2. El mismo tipo")
	fmt.Println("3. El mismo contenido")

	fmt.Println(headers.When)
	fmt.Println("a := [2]int{1, 2}")
	fmt.Println("b := [...]int{1, 2}")
	fmt.Println("c := [2]int{2, 1}")

	fmt.Println(headers.Then)
	fmt.Println("a = b => ", a == b)
	fmt.Println("a = c => ", a == c)

}
