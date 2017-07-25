package main

import "fmt"
import "strconv"

func main() {
	// Tipo de datos STring
	// Los string son una secuencia de bytes, no de caracteres un slice de bytes
	// Son indexables
	// Son inmutables

	var cadena string
	cadena = "Hola Mundo"

	fmt.Println("\n Nuestra cadena de ejemplo es: ", cadena)
	fmt.Println("\n la longitud de nuestra cadena de ejemplo es: ", len(cadena))
	fmt.Println("\n El indice 2 en unicode UTF8 de nuestra cadena de ejemplo es: ", cadena[2])
	fmt.Println("\n Las 4 primeros caracteres de la caneda son: ", cadena[0:4])

	// los strings son indexables e inmutables
	// por ejemplo no podriamos hacer:
	// cadena[3] = "c"
	// pero si podemos hacer
	cadena = cadena + " Nuevamente"
	// cadena += " Nuevamente"
	fmt.Println("\n Añadimos Nuevamente a nuestra anterior cadena: " + cadena)
	cadena = `
    <html>
        <head>
        </head>
        <body>
        </body>
    </html>
    `
	fmt.Println("\n Una cadena multiline seria: " + cadena)

	cadena = "Hola mundo \"25\""
	fmt.Println("\n cadenas con caracteres especiales como \"comillas dobles\": " + cadena)

	cadena = "Hola mundo \\"
	fmt.Println("\n cadenas con caracteres especiales como \"soble slash\": " + cadena)

	edad := 29
	cadena = "La edad es " + strconv.Itoa(edad)
	fmt.Println(cadena)

}
