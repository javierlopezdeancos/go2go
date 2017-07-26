package main

import (
	"fmt"
	"strconv"
	"../helpers"
)

func main() {

	headers := helpers.GetHeaders()

	title := `*                       Los Strings                            *`

	fmt.Println(headers.Break)
	fmt.Println(headers.Title)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.TitleSide)
	fmt.Println(title)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.TitleSide)
	fmt.Println(headers.Title)

	var cadena string
	cadena = "Hola Mundo"

	fmt.Println(headers.Given)
	fmt.Println("\nLos string son una secuencia de bytes, no de caracteres, un slice de bytes")
	fmt.Println("Los strings son indexables y sus indices son inmutables,")
	fmt.Println("es decir, no podemos asignar un valor a un indice, no podemos hacer cadena[3] = 'c'")

	fmt.Println(headers.When)
	fmt.Println("Nuestra cadena de ejemplo es cadena := ", cadena)

	fmt.Println(headers.Then)
	fmt.Println("La longitud de nuestra cadena de ejemplo es = ", len(cadena))
	fmt.Println("El indice 2 en unicode UTF8 de nuestra cadena de ejemplo es cadena[2] = ", cadena[2])
	fmt.Println("Las 4 primeros caracteres de la caneda son cadena[0:4] = : ", cadena[0:4])

	fmt.Println(headers.Break)

	cadena = cadena + " Nuevamente"

	fmt.Println(headers.When)
	fmt.Println("Añadimos Nuevamente a nuestra anterior cadena = cadena + ' Nuevamente'")

	fmt.Println(headers.Then)
	fmt.Println("Ahora cadena = " + cadena)

	fmt.Println(headers.Break)

	cadena = `
		<html>
			<head>
			</head>
			<body>
			</body>
		</html>
		`

	fmt.Println(headers.Given)
	fmt.Println("Definiendo una cadena multiline con comillas ` `")

	fmt.Println(headers.Then)
	fmt.Println("La cadena multiline seria: " + cadena)

	fmt.Println(headers.Break)

	cadena = "Hola mundo \"25\""

	fmt.Println(headers.Given)
	fmt.Println("Definiendo un string cadena = `Hola mundo \"25\"`")

	fmt.Println(headers.Then)
	fmt.Println("Cadenas con caracteres especiales como \"comillas dobles\": " + cadena)

	fmt.Println(headers.Break)

	cadena = "Hola mundo \\"

	fmt.Println(headers.Given)
	fmt.Println("Definiendo un string cadena = `Hola mundo \\`")

	fmt.Println(headers.Then)
	fmt.Println("\n cadenas con caracteres especiales como \"doble slash\": " + cadena)

	fmt.Println(headers.Break)

	edad := 29

	fmt.Println(headers.Given)
	fmt.Println("Definiendo una variable entera edad = 29")

	fmt.Println(headers.Then)
	fmt.Println("Podemos convertir a string con strconv")
	fmt.Println("cadena = 'La edad es ' + strconv.Itoa(edad)")
	cadena = "cadena = La edad es " + strconv.Itoa(edad)
	fmt.Println(cadena)

}
